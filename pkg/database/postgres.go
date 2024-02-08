package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/saufiroja/cqrs/config"
	"github.com/saufiroja/cqrs/pkg/logger"
)

type Postgres struct {
	*sql.DB
}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (p *Postgres) StartDatabase(conf *config.AppConfig, log *logger.Logger) *Postgres {
	host := conf.Postgres.Host
	port := conf.Postgres.Port
	user := conf.Postgres.User
	pass := conf.Postgres.Pass
	dbname := conf.Postgres.Name
	ssl := conf.Postgres.SSL

	str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, dbname, ssl)

	db, err := sql.Open("postgres", str)
	if err != nil {
		errMsg := fmt.Sprintf("error connecting to postgres: %v", err)
		log.StartLogger("postgres.go", "NewPostgres").Error(errMsg)
		panic(err)
	}

	// check connection
	err = db.Ping()
	if err != nil {
		errMsg := fmt.Sprintf("error connecting to postgres: %v", err)
		log.StartLogger("postgres.go", "NewPostgres").Error(errMsg)
		panic(err)
	}

	// set max connection
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5)
	db.SetConnMaxIdleTime(5)

	log.StartLogger("postgres.go", "NewPostgres").Info("connected to postgres")

	p.DB = db

	return p
}

func (p *Postgres) StartTransaction() (*sql.Tx, error) {
	return p.Begin()
}

func (p *Postgres) CommitTransaction(tx *sql.Tx) {
	err := tx.Commit()
	if err != nil {
		panic(err)
	}
}

func (p *Postgres) RollbackTransaction(tx *sql.Tx) {
	err := tx.Rollback()
	if err != nil {
		panic(err)
	}
}

func (p *Postgres) Close(ctx context.Context) {
	err := p.DB.Close()
	if err != nil {
		panic(err)
	}
}
