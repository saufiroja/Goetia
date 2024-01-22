package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/saufiroja/cqrs/config"
	"log"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(conf *config.AppConfig) *Postgres {
	host := conf.Postgres.Host
	port := conf.Postgres.Port
	user := conf.Postgres.User
	pass := conf.Postgres.Pass
	dbname := conf.Postgres.Name
	ssl := conf.Postgres.SSL

	str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, dbname, ssl)

	db, err := sql.Open("postgres", str)
	if err != nil {
		return nil
	}

	// check connection
	err = db.Ping()
	if err != nil {
		return nil
	}

	// set max connection
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5)
	db.SetConnMaxIdleTime(5)

	log.Println("connected to postgres")

	return &Postgres{
		db: db,
	}
}

func (p *Postgres) Open() *sql.DB {
	return p.db
}

func (p *Postgres) StartTransaction() (*sql.Tx, error) {
	return p.db.Begin()
}

func (p *Postgres) CommitTransaction(tx *sql.Tx) {
	_ = tx.Commit()
}

func (p *Postgres) RollbackTransaction(tx *sql.Tx) {
	_ = tx.Rollback()
}

func (p *Postgres) Close(ctx context.Context) {
	err := p.db.Close()
	if err != nil {
		panic(err)
	}
}
