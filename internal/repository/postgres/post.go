package postgres

import (
	"boobook/internal/repository"
	"boobook/internal/repository/model"
	"database/sql"
	"fmt"
)

type postRepository struct {
	DBReadConnection  *sql.DB
	DBWriteConnection *sql.DB
}

func NewPostRepository(DBReadConnection *sql.DB, DBWriteConnection *sql.DB) repository.PostRepository {
	return &postRepository{
		DBReadConnection:  DBReadConnection,
		DBWriteConnection: DBWriteConnection,
	}
}

func (r *postRepository) GetList() ([]*model.Post, error) {

	const fnErr = "repository.postgres.postRepository.GetList"

	var posts []*model.Post
	var rows *sql.Rows
	var err error

	rows, err = r.DBReadConnection.Query(`
		SELECT id, user_id, title, body, created_at, updated_at
		FROM posts 
		ORDER BY created_at DESC
		LIMIT 1000
	`)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", fnErr, err)
	}

	defer func() {
		if err = rows.Close(); err != nil {
			err = fmt.Errorf("%s: %w", fnErr, err)
		}
	}()

	for rows.Next() {
		var post model.Post
		if err = rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Title,
			&post.Body,
			&post.CreatedAt,
			&post.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("%s: %w", fnErr, err)
		}
		posts = append(posts, &post)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", fnErr, err)
	}

	return posts, nil
}
