package storage

import (
	"PostService/internal/domain/models"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type PostStorage struct {
	log *slog.Logger
	db  *sql.DB
}

func (p *PostStorage) Exec(ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	return fn(tx)
}

func (p *PostStorage) InsertPost(ctx context.Context, tx *sql.Tx, author, title, body string) (int64, error) {
	op := "storage.InsertPost"

	query := `INSERT INTO posts ("author", "title", "body") VALUES ($1, $2, $3) RETURNING id`
	var postId int64
	err := tx.QueryRowContext(ctx, query, author, title, body).Scan(&postId)
	if err != nil {
		p.log.Error(fmt.Sprintf("%s: %v", op, err))
		return 0, fmt.Errorf("%s: %v", op, err)
	}
	return postId, nil
}

func (p *PostStorage) InsertTagsByPost(ctx context.Context, tx *sql.Tx, postId int64, tagIds []int32) error {
	op := "storage.InsertTagsByPost"
	stmt, err := tx.Prepare("INSERT INTO post_tags (post_id, tag_id) VALUES ($1, $2)")
	if err != nil {
		p.log.Error(fmt.Sprintf("%s: %v", op, err))
		return fmt.Errorf("%s: %v", op, err)
	}

	defer stmt.Close()

	for _, tagId := range tagIds {
		_, err = stmt.ExecContext(ctx, postId, tagId)
		if err != nil {
			p.log.Error(fmt.Sprintf("%s: %v", op, err))
			return fmt.Errorf("%s: %v", op, err)
		}
	}
	return nil
}

func (p *PostStorage) GetPosts(ctx context.Context, request models.GetPostsRequest) (models.GetPostsResponse, error) {
	query := "select * from posts"

	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return models.GetPostsResponse{}, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan()
		if err != nil {
			return models.GetPostsResponse{}, err
		}
	}

	err = rows.Err()
	if err != nil {
		return models.GetPostsResponse{}, err
	}

	return models.GetPostsResponse{}, nil
}

func (p *PostStorage) GetPost(ctx context.Context, request models.GetPostResponse) (models.GetPostResponse, error) {
	return models.GetPostResponse{}, nil
}

func New(log *slog.Logger, db *sql.DB) *PostStorage {
	return &PostStorage{
		log: log,
		db:  db,
	}
}
