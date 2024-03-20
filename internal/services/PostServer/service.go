package PostServer

import (
	"PostService/internal/domain/models"
	"context"
	"log/slog"
)

type Storage interface {
	GetPosts(ctx context.Context, request models.GetPostsRequest) (models.GetPostsResponse, error)
	GetPost(ctx context.Context, request models.GetPostResponse) (models.GetPostResponse, error)
	CreatePost(ctx context.Context, request models.CreatePostRequest) error
}

type Post struct {
	log     *slog.Logger
	storage Storage
}

func New(log *slog.Logger, storage Storage) *Post {
	return &Post{
		log:     log,
		storage: storage,
	}
}

func (p *Post) GetPosts(ctx context.Context, request models.GetPostsRequest) (models.GetPostsResponse, error) {
	return models.GetPostsResponse{}, nil
}

func (p *Post) GetPost(ctx context.Context, request models.GetPostResponse) (models.GetPostResponse, error) {
	return models.GetPostResponse{}, nil
}

func (p *Post) CreatePost(ctx context.Context, request models.CreatePostRequest) error {
	return nil
}
