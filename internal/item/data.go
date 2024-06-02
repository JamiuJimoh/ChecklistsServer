package item

import (
	"errors"
	"fmt"

	"github.com/JamiuJimoh/checklist/internal/db"
	"github.com/jackc/pgx/v5"
)

func retrieveById(id int) (item ChecklistItem, err error) {
	err = db.Conn.QueryRow(ctx, "select id, description, created_at, expire_at from items where id = $1", id).Scan(&item.Id, &item.Description, &item.CreatedAt, &item.ExpireAt)
	return item, err
}

func RetrieveByChecklistID(checklistId int) (items []ChecklistItem, err error) {
	rows, err := db.Conn.Query(ctx, "select id, description, created_at, expire_at from items where items.checklist_id = $1 order by id", checklistId)

	items, err = pgx.CollectRows(rows, func(row pgx.CollectableRow) (ChecklistItem, error) {
		var item ChecklistItem
		err := row.Scan(&item.Id, &item.Description, &item.CreatedAt, &item.ExpireAt)
		return item, err
	})

	return items, err
}

func AddChecklistItem(checklistID int, item *ChecklistItem) (err error) {
	err = db.Conn.QueryRow(ctx,
		"insert into items (description, checklist_id, created_at, expire_at) values ($1, $2, $3, $4) returning id",
		&item.Description, checklistID, &item.CreatedAt, &item.ExpireAt,
	).Scan(&item.Id)
	return err
}

func update(ci *ChecklistItem) (err error) {
	tag, err := db.Conn.Exec(ctx, "update items set description = $2 where id = $1", ci.Id, ci.Description)
	if err != nil {
		err = errors.New("error occured while updating checklist")
	}
	if tag.RowsAffected() != 1 {
		err = errors.New(fmt.Sprintf("could not update item with id: %d", ci.Id))
	}
	return err
}

func deleteItem(id int) (err error) {
	tag, err := db.Conn.Exec(ctx, "delete from items where id = $1", id)
	if err != nil {
		err = errors.New("error occured while deleting checklist")
	}
	if tag.RowsAffected() != 1 {
		err = errors.New(fmt.Sprintf("could not delete item with id: %d", id))
	}
	return err
}
