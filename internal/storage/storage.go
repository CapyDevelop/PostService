package storage

import (
	"PostService/internal/domain/models"
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
)

type PostStorage struct {
	log *slog.Logger
	db  *sql.DB
}

func (p *PostStorage) CreatePost(ctx context.Context, request models.CreatePostRequest) error {
	op := "PostStorage.CreatePost"

	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		p.log.Error(fmt.Sprintf("%s:%v", op, err))
		return fmt.Errorf("%s:%v", op, err)
	}
	defer tx.Rollback()

	w, _ := tx.Query("SELECT CURRENT_CATALOG")
	for w.Next() {
		var tableName string
		if err := w.Scan(&tableName); err != nil {
			log.Fatal(err)
		}
		fmt.Println(tableName)
	}

	query := `INSERT INTO posts ("author", "title", "body") VALUES ($1, $2, $3) RETURNING id`
	var postId int64
	err = tx.QueryRowContext(ctx, query, request.Post.Author, request.Post.Title, request.Post.Body).Scan(&postId)
	if err != nil {
		p.log.Error(fmt.Sprintf("%s: %v", op, err))
		return fmt.Errorf("%s: %v", op, err)
	}

	if err != nil {
		fmt.Println(2)

		p.log.Error(fmt.Sprintf("%s: %v", op, err))
		return fmt.Errorf("%s: %v", op, err)
	}

	stmt, err := tx.Prepare("INSERT INTO post_tags (post_id, tag_id) VALUES ($1, $2)")
	if err != nil {
		fmt.Println(3)

		p.log.Error(fmt.Sprintf("%s:%v", op, err))
		return fmt.Errorf("%s:%v", op, err)
	}

	defer stmt.Close()

	for _, tagId := range request.Tags {
		_, err = stmt.ExecContext(ctx, postId, tagId)
		if err != nil {
			fmt.Println(4)

			p.log.Error(fmt.Sprintf("%s:%v", op, err))
			return fmt.Errorf("%s:%v", op, err)
		}
	}

	err = tx.Commit()
	if err != nil {
		p.log.Error(fmt.Sprintf("%s:%v", op, err))
		return fmt.Errorf("%s:%v", op, err)
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
