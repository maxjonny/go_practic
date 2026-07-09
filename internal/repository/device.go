package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IDeviceRepository interface {
	GetActiveNode(ctx context.Context, device string) ([]string, error)
	GetDeviceRelations(ctx context.Context, device string) ([]DeviceRelations, error)
}

type DeviceRepository struct {
	pgPool *pgxpool.Pool
}

func NewDeviceRepository(pgPool *pgxpool.Pool) *DeviceRepository {
	return &DeviceRepository{pgPool}
}

func (dr *DeviceRepository) GetActiveNode(ctx context.Context, device string) ([]string, error) {
	nodes := make([]string, 0, 1)
	queryString := fmt.Sprintf(`select
                                (link.doc->>'node_id') as node_id
                            FROM checkbox.boxes box
                              inner join checkbox.box_relation br on (br.doc->> 'box_id')::int = box.id
                              inner join structure.link_nodes link on (link.doc->>'link_object_id') = (br.doc->>'project_id')
                            WHERE (box.doc->>'equipmentModel') = '%s'
                                and (link.doc->>'link_object') = 'project'
								`, device)
	rows, err := dr.pgPool.Query(ctx, queryString)
	if err != nil {
		return nodes, err
	}
	defer rows.Close()

	for rows.Next() {
		var node_id string
		err := rows.Scan(&node_id)
		if err != nil {
			return nodes, err
		}
		nodes = append(nodes, node_id)
	}
	return nodes, nil
}

func (dr *DeviceRepository) GetDeviceRelations(ctx context.Context, device string) ([]DeviceRelations, error) {

	allRelations := make([]DeviceRelations, 0)
	queryString := `select 
					(link.doc->>'node_id')::int as node_id, (br.doc->>'project_id')::int as project_id
				from checkbox.boxes box
				inner join checkbox.box_relation br on (br.doc->>'box_id')::int = box.id
				inner join structure.link_nodes link on (link.doc->>'link_object_id') = (br.doc->>'project_id') and (link.doc->>'link_object') = 'project'
				where (box.doc->>'equipmentModel') = $1`

	rows, err := dr.pgPool.Query(ctx, queryString, device)

	if err != nil {
		return allRelations, err
	}
	defer rows.Close()

	for rows.Next() {
		var DeviceRelations DeviceRelations
		err = rows.Scan(&DeviceRelations.NodeId, &DeviceRelations.ProjectId)
		if err != nil {
			return nil, err
		}
		allRelations = append(allRelations, DeviceRelations)
	}

	return allRelations, nil
}
