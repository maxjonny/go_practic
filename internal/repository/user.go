package repository

import (
	"context"
	"encoding/json"
	"fmt"
	m "main/internal/models"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type IUserRepository interface {
	DropCache(ctx context.Context, device string)
	CreateCache(ctx context.Context, device string, users []m.UserCard)
	GetUser(ctx context.Context, device string, index string) (*m.UserCard, error)
	GetUserByNodes(ctx context.Context, nodeIds []string) ([]m.UserCard, error)
}

type UserRepository struct {
	pgPool  *pgxpool.Pool
	rClient *redis.Client
}

func NewUserRepository(pgPool *pgxpool.Pool, rClient *redis.Client) *UserRepository {
	return &UserRepository{pgPool, rClient}
}

func (ur *UserRepository) DropCache(ctx context.Context, device string) {

	var cursor uint64
	var keys []string

	pattern := fmt.Sprintf("checkbox:%s:*", device)

	for {
		var err error
		var scannedKeys []string

		// Сканируем ключи по шаблону
		scannedKeys, cursor, err = ur.rClient.Scan(ctx, cursor, pattern, 100).Result()
		if err != nil {
			fmt.Printf("ошибка сканирования: %s", err)
		}

		keys = append(keys, scannedKeys...)

		// Если дошли до конца
		if cursor == 0 {
			break
		}
	}

	// Удаляем все найденные ключи
	if len(keys) > 0 {
		deleted, err := ur.rClient.Unlink(ctx, keys...).Result()
		if err != nil {
			fmt.Printf("ошибка удаления: %s", err)
		}
		fmt.Printf("Удалено %d ключей по шаблону '%s'\n", deleted, pattern)
	}
}

func (ur *UserRepository) CreateCache(ctx context.Context, device string, users []m.UserCard) {

	pipe := ur.rClient.Pipeline()

	for ind, elem := range users {
		key := fmt.Sprintf("checkbox:%s:%d", device, ind+1)
		data, _ := json.Marshal(elem)
		pipe.Set(ctx, key, string(data), 24*time.Hour)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		fmt.Println(err)
	}
}

func (ur *UserRepository) GetUser(ctx context.Context, device string, index string) (*m.UserCard, error) {

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

func (ur *UserRepository) GetUserByNodes(ctx context.Context, nodeIds []string) ([]m.UserCard, error) {
	nodes := make([]m.UserCard, 0)
	queryString := fmt.Sprintf(`SELECT distinct
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
                            WHERE string_to_array(tn.doc->>'path', '-') && array['%s'] 
                            AND tn.doc->>'status' = 'active' 
                            AND tnr.doc->>'status' = 'active';
								`, strings.Join(nodeIds, "','"))

	rows, err := ur.pgPool.Query(ctx, queryString)
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
