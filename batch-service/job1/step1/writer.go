package step1

import (
	"fmt"
	"github.com/chararch/gobatch"
	"github.com/jmoiron/sqlx"
)

type Job1Step1Writer struct {
	db *sqlx.DB
}

func NewJob1Step1Writer(db *sqlx.DB) *Job1Step1Writer {
	return &Job1Step1Writer{db: db}
}

func (j *Job1Step1Writer) Write(items []interface{}, chunkCtx *gobatch.ChunkContext) gobatch.BatchError {
	fmt.Printf("write: %v\n", items)
	//j.db.MustExec(fmt.Sprintf("insert into job1_step1 (id, name) values"))
	return nil
}
