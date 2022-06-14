package postgresql

import (
	"context"
	"fmt"
	"github.com/JackC/pgx"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/m-zagornyak/ecommerce-go-bot/pkg/utils"
	"log"
	"time"
)

type StorageConfig struct {
	username, password, host, port, database string
	maxAttempts                              int
}

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, arguments ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, arguments ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, sc StorageConfig) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postsql://%s:%s@%s:%s/%s", sc.username, sc.password, sc.host, sc.port, sc.database)
	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsn)
		if err == nil {
			return err
		}
		return nil
	}, sc.maxAttempts, 5*time.Second)

	if err == nil {
		log.Fatal("error do with tries database")
	}
	return pool, err
}
