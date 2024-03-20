package handlers

import (
	"PostService/internal/domain/models"
	"context"
	"github.com/CapyDevelop/protos/gen/go/PostService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PostServ interface {
	GetPosts(ctx context.Context, request models.GetPostsRequest) (models.GetPostsResponse, error)
	GetPost(ctx context.Context, request models.GetPostResponse) (models.GetPostResponse, error)
	//SetRating
	//GetComments
	//SetComment
	//SetCommentRating
	CreatePost(ctx context.Context, request models.CreatePostRequest) error
}

type serverAPI struct {
	PostService.PostServiceServer
	postServ PostServ
}

func (s *serverAPI) GetPost(ctx context.Context, req *PostService.GetPostRequest) (*PostService.GetPostResponse, error) {
	if req.GetPostId() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "postId is invalid")
	}

	post := PostService.Post{
		Title:     "qwerty",
		Body:      "qazxswedcvfrtgbnhyujmkiol",
		CreatedAt: "",
		Rating:    7.4,
		Author:    nil,
		Tags:      []string{"a", "b"},
		PostId:    12,
	}
	return &PostService.GetPostResponse{
		Post:    &post,
		IsOwner: true,
	}, nil
}

func (s *serverAPI) SetRating(ctx context.Context, request *PostService.SetRatingRequest) (*PostService.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serverAPI) GetComments(ctx context.Context, request *PostService.GetCommentsRequest) (*PostService.GetCommentsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serverAPI) SetComment(ctx context.Context, request *PostService.SetCommentRequest) (*PostService.SetCommentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serverAPI) SetCommentRating(ctx context.Context, request *PostService.SetCommentRatingRequest) (*PostService.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serverAPI) CreatePost(ctx context.Context, request *PostService.CreatePostRequest) (*PostService.Empty, error) {
	//TODO implement me
	//panic("implement me")
	tags := request.GetTags() // предполагается, что это возвращает []*Tag
	var result []int32
	for _, tag := range tags {
		result = append(result, (*tag).Id)
	}
	req := models.CreatePostRequest{
		Post: models.Post{
			Author: request.GetUser().Login,
			Title:  request.GetPost().Title,
			Body:   request.GetPost().Body,
		},
		Tags: result,
	}
	err := s.postServ.CreatePost(ctx, req)
	if err != nil {

		return nil, err
	}
	return &PostService.Empty{}, nil
}

func (s *serverAPI) GetPosts(ctx context.Context, req *PostService.GetPostsRequest) (*PostService.GetPostsResponse, error) {
	panic("implement me")
}

func Register(gRPC *grpc.Server, postServ PostServ) {
	PostService.RegisterPostServiceServer(gRPC, &serverAPI{postServ: postServ})
}
