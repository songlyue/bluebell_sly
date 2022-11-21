package logic

import (
	"bluebell_sly/dao/postgres"
	"bluebell_sly/dao/redis"
	"bluebell_sly/models"
	"bluebell_sly/pkg/snowflake"
	"fmt"

	"go.uber.org/zap"
)

func CreatePost(post *models.Post) (err error) {
	postId, err := snowflake.GetId()
	if err != nil {
		zap.L().Error("snowflak err ", zap.Error(err))
	}
	post.PostID = postId
	// 创建帖子
	if err := postgres.CreatePost(post); err != nil {
		zap.L().Error("postgresql create post err", zap.Error(err))
		return err
	}
	community, err := postgres.GetCommunityById(fmt.Sprint(post.CommunityID))
	if err != nil {
		zap.L().Error("postgresql GetCommunityById err", zap.Error(err))
		return err
	}
	if err := redis.CreatePost(
		fmt.Sprint(post.PostID),
		fmt.Sprint(post.AuthorId),
		post.Title,
		TruncateByWords(post.Content, 120),
		community.CommunityName); err != nil {
		zap.L().Error("redis.CreatePost failed", zap.Error(err))
		return err
	}
	return
}

func GetPost(postId string) (post *models.ApiPostDetail, err error) {
	post, err = postgres.GetPostById(postId)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(postID) failed", zap.String("post_id", postId), zap.Error(err))
		return nil, err
	}
	user, err := postgres.GetUserById(fmt.Sprint(post.AuthorId))
	if err != nil {
		zap.L().Error("postgres.GetUserByID() failed", zap.String("author_id", fmt.Sprint(post.AuthorId)), zap.Error(err))
		return
	}
	post.AuthorId = user.UserID
	post.AuthorName = user.UserName
	community, err := postgres.GetCommunityById(fmt.Sprint(post.CommunityID))
	if err != nil {
		zap.L().Error("postgres.GetCommunityById() failed", zap.String("community_id", fmt.Sprint(post.CommunityID)), zap.Error(err))
		return
	}
	post.CommunityName = community.CommunityName
	return post, nil
}

func GetPgPostList() (data []*models.ApiPostDetail) {
	posts, err := postgres.GetPgPostList()
	fmt.Println(posts, err)
	return
}
