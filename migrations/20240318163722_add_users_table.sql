-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

create table posts
(
    id         serial primary key,
    author     varchar(255) not null,
    title      varchar(255) not null,
    body       text         not null,
    created_at timestamp default current_timestamp
);

create table tags
(
    id    serial primary key,
    title varchar(255) not null
);

create table post_tags
(
    id      serial primary key,
    post_id integer not null,
    tag_id  integer not null,
    constraint FK_post foreign key (post_id) references Posts (id),
    constraint FK_tags foreign key (tag_id) references Tags (id)
);

create table rating
(
    id         serial primary key,
    author     varchar(255) not null,
    post_id    integer      not null,
    rating     integer      not null,
    created_at timestamp default current_timestamp,
    constraint FK_post foreign key (post_id) references Posts (id)
);

create table comments
(
    id         serial primary key,
    post_id    integer      not null,
    author     varchar(255) not null,
    body       text         not null,
    parent_id  integer   default null,
    created_at timestamp default current_timestamp,
    constraint FK_parent_id foreign key (parent_id) references Comments (id),
    constraint FK_post foreign key (post_id) references Posts (id)
);

create table likes
(
    id          serial primary key,
    comment_id  integer      not null,
    author      varchar(255) not null,
    is_positive bool         not null,
    constraint FK_comment foreign key (comment_id) references Comments (id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

drop table likes;
drop table comments;
drop table rating;
drop table post_tags;
drop table tags;
drop Table posts;

-- +goose StatementEnd
