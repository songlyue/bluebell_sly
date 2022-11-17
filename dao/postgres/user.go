package postgres

import (
	"bluebell_sly/models"
	"bluebell_sly/pkg/snowflake"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
)

const secret = "songlongyue.com"

func encryptPassword(data []byte) (result string) {
	md5 := md5.New()
	md5.Write([]byte(secret))
	return hex.EncodeToString(md5.Sum(data))
}

func Register(user *models.User) (err error) {
	sqlstr := "select count(1) from users where username = $1"
	var count int64
	err = db.Get(&count, sqlstr, user.UserName)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return err
	}
	if count > 0 {
		fmt.Println(err)
		return ErrorUserExit
	}
	userId, err := snowflake.GetId()
	if err != nil {
		return ErrorGenIDFailed
	}
	password := encryptPassword([]byte(user.Password))
	sqlstr = "insert into users(user_id,username,password) values ($1,$2,$3)"
	_, err = db.Exec(sqlstr, userId, user.UserName, password)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func Login(user *models.User) (err error) {
	originPassword := user.Password
	sqlstr := "select user_id,username,password from users where username = $1"
	if err = db.Get(user, sqlstr, user.UserName); err != nil {
		return
	}
	if err == sql.ErrNoRows {
		return ErrorUserExit
	}
	password := encryptPassword([]byte(originPassword))
	if user.Password != password {
		return ErrorPasswordWrong
	}
	return
}
