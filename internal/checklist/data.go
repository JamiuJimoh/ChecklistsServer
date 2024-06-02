package checklist

import (
	"errors"
	"fmt"

	"github.com/JamiuJimoh/checklist/internal/db"
	"github.com/JamiuJimoh/checklist/internal/item"
	"github.com/jackc/pgx/v5"
)

func retrieveById(id int) (checklist Checklist, err error) {

	err = db.Conn.QueryRow(ctx, "select id, title, created_at from checklists where id = $1", id).Scan(&checklist.Id, &checklist.Title, &checklist.CreatedAt)

	items, err := item.RetrieveByChecklistID(id)
	checklist.Items = append(checklist.Items, items...)

	return checklist, err
}

func retrieveAll(limit, offset int) (checklists []Checklist, err error) {
	rows, err := db.Conn.Query(ctx, "select id, title, created_at from checklists order by id limit $1 offset $2", limit, offset)
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
		items, err := item.RetrieveByChecklistID(checklists[i].Id)
		if err != nil {
			return nil, err
		}
		checklists[i].Items = append(checklists[i].Items, items...)
	}

	return checklists, err
}

func insert(checklist *Checklist) (err error) {
	err = db.Conn.QueryRow(ctx, "insert into checklists (title, created_at) values ($1, $2) returning id", checklist.Title, checklist.CreatedAt).Scan(&checklist.Id)
	if len(checklist.Items) == 0 {
		return err
	}

	for i := range checklist.Items {
		err = item.AddChecklistItem(checklist.Id, &checklist.Items[i])
	}
	return err
}

func update(ch *Checklist) (err error) {
	tag, err := db.Conn.Exec(ctx, "update checklists set title = $2 where id = $1", ch.Id, ch.Title)
	if tag.RowsAffected() != 1 {
		err = errors.New(fmt.Sprintf("could not update checklist with id: %d", ch.Id))
	}
	return err
}

func deleteChecklist(id int) (err error) {
	tag, err := db.Conn.Exec(ctx, "delete from checklists where id = $1", id)
	if tag.RowsAffected() != 1 {
		err = errors.New(fmt.Sprintf("could not delete checklist with id: %d", id))
	}
	return err
}
