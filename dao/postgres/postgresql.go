package postgres

import (
	"bluebell_sly/settings"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(cfg *settings.PostgresqlConfig) (err error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		fmt.Println("connect db failed err: ", err)
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

func Close() {
	_ = db.Close()
}
