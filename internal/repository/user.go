package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	m "main/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type IUserRepository interface {
	DropCache(ctx context.Context, device string) error
	CreateCache(ctx context.Context, device string, users []m.UserCard) error
	GetUserCache(ctx context.Context, device string, index string) (*m.UserCard, error)

	GetUsersByNodes(ctx context.Context, nodeIds []string) ([]m.UserCard, error)
	GetUserRelations(ctx context.Context, userGid string) (UserRelations, error)
	GetWorkerId(ctx context.Context, cardId int, projectId int) (int, error)
	GetByGid(ctx context.Context, gID string) (int, m.UserCard, error)
}

type UserRepository struct {
	pgPool  *pgxpool.Pool
	rClient *redis.Client
}

func NewUserRepository(pgPool *pgxpool.Pool, rClient *redis.Client) *UserRepository {
	return &UserRepository{pgPool, rClient}
}

func (ur *UserRepository) DropCache(ctx context.Context, device string) error {

	var cursor uint64
	var keys []string

	pattern := fmt.Sprintf("checkbox:%s:*", device)

	for {
		var err error
		var scannedKeys []string

		// Сканируем ключи по шаблону
		scannedKeys, cursor, err = ur.rClient.Scan(ctx, cursor, pattern, 100).Result()
		if err != nil {
			return err
		}

		keys = append(keys, scannedKeys...)

		// Если дошли до конца
		if cursor == 0 {
			break
		}
	}

	// Удаляем все найденные ключи
	if len(keys) > 0 {
		_, err := ur.rClient.Unlink(ctx, keys...).Result()
		if err != nil {
			return err
		}
	}
	return nil
}

func (ur *UserRepository) CreateCache(ctx context.Context, device string, users []m.UserCard) error {

	pipe := ur.rClient.Pipeline()

	for ind, elem := range users {
		key := fmt.Sprintf("checkbox:%s:%d", device, ind+1)
		data, err := json.Marshal(elem)
		if err != nil {
			return err
		}
		pipe.Set(ctx, key, string(data), 24*time.Hour)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetUserCache(ctx context.Context, device string, index string) (*m.UserCard, error) {

	key := fmt.Sprintf("checkbox:%s:%s", device, index)

	data, err := ur.rClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var user m.UserCard
	err = json.Unmarshal([]byte(data), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetUsersByNodes(ctx context.Context, nodeIds []string) ([]m.UserCard, error) {
	nodes := make([]m.UserCard, 0)
	queryString := `SELECT distinct
                                (hc.doc->>'gID') as gID,
                                (hc.doc->>'gZBH') as gZBH,
                                (hc.doc->>'name') as name,
                                (hc.doc->>'deptName') as deptName,
                                (hc.doc->'img'->>'name') as img_name,
								(hc.doc->'img'->>'path') as img_path,
                                (hc.doc->>'fingerFeature') as fingerFeature
                            FROM structure.tree_nodes tn 
                            INNER JOIN tabel.tree_node_resource tnr ON tn.id = (tnr.doc->>'tree_node_id')::int
                            INNER JOIN checkbox.human_card hc ON (hc.doc->>'human_id')::int = (tnr.doc->>'resource_id')::int
                            WHERE string_to_array(tn.doc->>'path', '-') && $1 
                            AND tn.doc->>'status' = 'active' 
                            AND tnr.doc->>'status' = 'active';`

	rows, err := ur.pgPool.Query(ctx, queryString, nodeIds)
	if err != nil {
		return nodes, err
	}
	defer rows.Close()

	for rows.Next() {
		var user m.UserCard
		err := rows.Scan(&user.GID, &user.GZBH, &user.Name, &user.DeptName, &user.Img.Name, &user.Img.Path, &user.FaceFeature)
		if err != nil {
			return nodes, err
		}
		nodes = append(nodes, user)
	}
	return nodes, nil
}

func (ur *UserRepository) GetUserRelations(ctx context.Context, userGid string) (UserRelations, error) {

	var userRelatuin UserRelations

	queryString := `select 
						hc.id as card_id,
						COALESCE(array_agg(distinct node::int) FILTER (WHERE node IS NOT NULL), '{}') as nodes
					from checkbox.human_card hc
					left join tabel.tree_node_resource tnr on (tnr.doc->>'resource_id') = (hc.doc->>'human_id') and (tnr.doc->>'status') = 'active'
						left join "structure".tree_nodes tn on tn.id = (tnr.doc->>'tree_node_id')::int and (tn.doc->>'status') = 'active'
						left join lateral unnest(string_to_array(tn.doc->>'path', '-')) as node on true
					where (hc.doc->>'gID') = $1
					group by hc.id`

	err := ur.pgPool.QueryRow(ctx, queryString, userGid).Scan(&userRelatuin.UserCardId, &userRelatuin.NodeIds)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = nil
		}
	}

	return userRelatuin, err
}

func (ur *UserRepository) GetWorkerId(ctx context.Context, cardId int, projectId int) (int, error) {

	var workerId int

	queryString := `select id from checkbox.workers
					where (doc->>'human_card_id')::int = $1 and (doc->>'project_id')::int = $2`

	err := ur.pgPool.QueryRow(ctx, queryString, cardId, projectId).Scan(&workerId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = nil
		}
	}

	return workerId, err
}

func (ur *UserRepository) GetByGid(ctx context.Context, gID string) (int, m.UserCard, error) {

	var cardId int
	var user m.UserCard

	queryString := `select id, doc->>'gID', doc->'img'->>'name',doc->'img'->>'path',
						doc->>'gZBH', doc->>'name', doc->>'deptName',doc->'human_id',(doc->>'fromDevice')::bool,doc->>'fingerFeature'
					from checkbox.human_card hc where doc->>'gID' = $1`

	rows, err := ur.pgPool.Query(ctx, queryString, gID)
	if err != nil {
		return cardId, user, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&cardId, &user.GID, &user.Img.Name, &user.Img.Path, &user.GZBH, &user.Name, &user.DeptName, &user.HumanID, &user.FromDevice, &user.FingerFeature)
		if err != nil {
			return cardId, user, err
		}
	}

	return cardId, user, nil
}
