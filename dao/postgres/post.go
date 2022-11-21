package postgres

import (
	"bluebell_sly/models"
	"database/sql"

	"go.uber.org/zap"
)

func CreatePost(post *models.Post) (err error) {
	sqlStr := `insert into post(post_id, title, content, author_id, community_id,create_time,update_time)
	values ($1, $2, $3, $4, $5, $6, $7)`
	_, err = db.Exec(sqlStr, post.PostID, post.Title, post.Content, post.AuthorId, post.CommunityID,
		post.CreateTime, post.UpdateTime)
	if err != nil {
		zap.L().Error("insert  post err", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

func GetPostById(postId string) (post *models.ApiPostDetail, err error) {
	post = new(models.ApiPostDetail)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time,status
	from post
	where post_id = $1`
	err = db.Get(post, sqlStr, postId)
	if err == sql.ErrNoRows {
		err = ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query post failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	return
}

func GetPgPostList() (posts []*models.ApiPostDetail, err error) {
	//sqlStr := `select post_id, title, content, author_id, p.community_id, p.create_time,status, u.username,c.community_name
	//from post p
	//left join users u on p.author_id=u.user_id
	//left join community c on p.community_id=c.community_id
	//limit 2`
	sqlStr := `select post_id, title, content, author_id, community_id, create_time,status
	from post
	limit 2
	`
	//m := new(models.Post)
	posts = make([]*models.ApiPostDetail, 0, 2)
	err = db.Select(&posts, sqlStr)
	return
}
