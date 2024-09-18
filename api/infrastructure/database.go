package infrastructure

import (
	"context"
	"fmt"
	"log"
	"os"

	"entgo.io/ent/dialect"
	"github.com/ei-sugimoto/tatekae/api/infrastructure/ent"
	"github.com/ei-sugimoto/tatekae/api/infrastructure/ent/migrate"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*ent.Client
}

func NewDB() *DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	client, err := ent.Open(dialect.MySQL, dsn)
	if err != nil {
		panic(fmt.Sprintf("failed opening connection to mysql: %v", err))
	}

	env := os.Getenv("ENV")

	if env != "production" {
		client = client.Debug()
	}

	return &DB{client}
}

func (db *DB) Migrate() {
	err := db.Schema.Create(context.Background(), migrate.WithDropIndex(true), migrate.WithDropColumn(true))

	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
