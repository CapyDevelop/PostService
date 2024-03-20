package models

import (
	"time"
)

// Post table
type Post struct {
	ID        uint
	Author    string
	Title     string
	Body      string
	CreatedAt time.Time
}

// Tag table
type Tag struct {
	ID    uint   `gorm:"primary_key"`
	Title string `gorm:"type:varchar(255);not null"`
}

// PostTag table
type PostTag struct {
	ID     uint `gorm:"primary_key"`
	PostID uint `gorm:"not null"`
	TagID  uint `gorm:"not null"`
}

// Rating table
type Rating struct {
	ID        uint      `gorm:"primary_key"`
	Author    string    `gorm:"type:varchar(255);not null"`
	PostID    uint      `gorm:"not null"`
	Rating    int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
}

// Comment table
type Comment struct {
	ID        uint      `gorm:"primary_key"`
	PostID    uint      `gorm:"not null"`
	Author    string    `gorm:"type:varchar(255);not null"`
	Body      string    `gorm:"type:text;not null"`
	ParentID  *uint     `gorm:"default:null"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
}

// Like table
type Like struct {
	ID         uint   `gorm:"primary_key"`
	CommentID  uint   `gorm:"not null"`
	Author     string `gorm:"type:varchar(255);not null"`
	IsPositive bool   `gorm:"not null"`
}

type GetPostRequest struct {
	PostId int
	UserId string
}

type GetPostsRequest struct {
}

type SetRatingRequest struct {
}
type GetCommentsRequest struct {
}
type SetCommentRequest struct {
}
type SetCommentRatingRequest struct {
}
type CreatePostRequest struct {
	Post Post
	Tags []int32
}

type GetPostResponse struct {
	Post   Post
	Tags   []Tag
	Rating float32
}

type GetPostsResponse struct {
}

type SetRatingResponse struct {
}
type GetCommentsResponse struct {
}
type SetCommentResponse struct {
}
type SetCommentRatingResponse struct {
}
type CreatePostResponse struct {
}
