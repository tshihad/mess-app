package data

import (
	"mess-app/internal/core"
	"mess-app/internal/models"
	"time"
)

var (
	groupInsertQuery = `INSERT INTO cluster (title,author_id,created_at) VALUES($1,$2,$3)`
	insertGroupQuery = `INSERT INTO cluster_users ()`
)

func (r *repo) InsertGroup(group models.GroupPayload, authorID int) error {
	res, err := r.DB.Exec(groupInsertQuery, *group.GroupName, authorID, time.Now())
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if rowsAffected != 1 {
		return core.ErrConflict
	}
	return nil
}

func (r *repo) InsertUserToGroup(group models.GroupPayload, users []models.UserPayload) error {
	// insert user to existing group

	return nil
}
