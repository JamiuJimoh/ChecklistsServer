package checklist

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/JamiuJimoh/checklist/internal/item"
	"github.com/jackc/pgx/v5"
)

var Db *pgx.Conn

func init() {
	var err error
	connStr := "postgres://jamiu:jamiu@localhost:5432/checklist?sslmode=disable"
	Db, err = pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	// defer Db.Close(context.Background())
}

func retrieveById(id int) (checklist Checklist, err error) {
	err = Db.QueryRow(ctx, "select id, title, created_at from checklists where id = $1", id).Scan(&checklist.Id, &checklist.Title, &checklist.CreatedAt)

	items, err := itemsById(id)
	checklist.Items = append(checklist.Items, items...)

	return checklist, err
}

func retrieveAll(limit, offset int) (checklists []Checklist, err error) {
	rows, err := Db.Query(context.Background(), "select id, title, created_at from checklists limit $1 offset $2", limit, offset)
	if err != nil {
		return nil, err
	}

	chs, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (Checklist, error) {
		var checklist Checklist
		err := row.Scan(&checklist.Id, &checklist.Title, &checklist.CreatedAt)
		return checklist, err
	})
	checklists = append(checklists, chs...)

	for i := range checklists {
		items, err := itemsById(checklists[i].Id)
		if err != nil {
			return nil, err
		}
		checklists[i].Items = append(checklists[i].Items, items...)
	}

	return checklists, err
}

func itemsById(checklistId int) (items []item.ChecklistItem, err error) {
	rows, err := Db.Query(ctx, "select id, description, created_at, expire_at from items where items.checklist_id = $1", checklistId)

	items, err = pgx.CollectRows(rows, func(row pgx.CollectableRow) (item.ChecklistItem, error) {
		var item item.ChecklistItem
		err := row.Scan(&item.Id, &item.Description, &item.CreatedAt, &item.ExpireAt)
		return item, err
	})
	return items, err
}

func insert(checklist *Checklist) (err error) {
	err = Db.QueryRow(ctx, "insert into checklists (title, created_at) values ($1, $2) returning id", checklist.Title, checklist.CreatedAt).Scan(&checklist.Id)
	for i := range checklist.Items {
		err = addItem(checklist.Id, &checklist.Items[i])
	}
	return err
}

func addItem(checklistID int, item *item.ChecklistItem) (err error) {
	err = Db.QueryRow(ctx,
		"insert into items (description, checklist_id, created_at, expire_at) values ($1, $2, $3, $4) returning id",
		&item.Description, &checklistID, &item.CreatedAt, &item.ExpireAt,
	).Scan(&item.Id)
	return err
}

func update(ch *Checklist) (err error) {
	tag, err := Db.Exec(ctx, "update checklists set title = $2 where id = $1", ch.Id, ch.Title)
	if tag.RowsAffected() != 1 {
		err = errors.New(fmt.Sprintf("could not update checklist with id: %d", ch.Id))
	}
	return err
}

func delete(id int) (err error) {
	tag, err := Db.Exec(ctx, "delete from checklists where id = $1", id)
	if tag.RowsAffected() != 1 {
		err = errors.New(fmt.Sprintf("could not delete checklist with id: %d", id))
	}
	return err
}
