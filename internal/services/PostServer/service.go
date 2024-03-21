package PostServer

import (
	"PostService/internal/domain/models"
	"context"
	"database/sql"
	"log/slog"
)

type Storage interface {
	Exec(ctx context.Context, fn func(*sql.Tx) error) error
	InsertPost(ctx context.Context, tx *sql.Tx, author, title, body string) (int64, error)
	InsertTagsByPost(ctx context.Context, tx *sql.Tx, postId int64, tagIds []int32) error
}

type Post struct {
	log     *slog.Logger
	storage Storage
}

func (p *Post) GetPosts(ctx context.Context, request models.GetPostsRequest) (models.GetPostsResponse, error) {
	return models.GetPostsResponse{}, nil
}

func (p *Post) GetPost(ctx context.Context, request models.GetPostRequest) (models.GetPostResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Post) SetRating(ctx context.Context, request models.SetRatingRequest) error {
	//TODO implement me
	panic("implement me")
}

func (p *Post) GetComments(ctx context.Context, request models.GetCommentsRequest) (models.GetCommentsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Post) SetComment(ctx context.Context, request models.SetCommentRequest) error {
	//TODO implement me
	panic("implement me")
}

func (p *Post) SetCommentRating(ctx context.Context, request models.SetCommentRatingRequest) error {
	//TODO implement me
	panic("implement me")
}

func New(log *slog.Logger, storage Storage) *Post {
	return &Post{
		log:     log,
		storage: storage,
	}
}

func (p *Post) CreatePost(ctx context.Context, request models.CreatePostRequest) error {

	err := p.storage.Exec(ctx, func(tx *sql.Tx) error {
		postId, err := p.storage.InsertPost(ctx, tx, request.Post.Author, request.Post.Title, request.Post.Body)
		if err != nil {
			return err
		}

		err = p.storage.InsertTagsByPost(ctx, tx, postId, request.Tags)
		if err != nil {
			return err
		}

		return nil
	})
	return err
}
