package step1

import (
	"github.com/chararch/gobatch"
	"github.com/jmoiron/sqlx"
)

const insertPersonQuery = `
	INSERT INTO persons (first_name, last_name, email, profession, processed, created_at, updated_at)
	VALUES (:first_name, :last_name, :email, :profession, :processed, :created_at, :updated_at)`

type Job1Step1Writer struct {
	db *sqlx.DB
}

func NewJob1Step1Writer(db *sqlx.DB) *Job1Step1Writer {
	return &Job1Step1Writer{db: db}
}

func (j *Job1Step1Writer) Write(items []interface{}, chunkCtx *gobatch.ChunkContext) gobatch.BatchError {
	if len(items) == 0 {
		return nil
	}

	persons := make([]PersonV1Entity, 0, len(items))
	for _, item := range items {
		person, ok := item.(PersonV1Entity)
		if !ok {
			return gobatch.NewBatchError(gobatch.ErrCodeGeneral, "unexpected item type %T in writer", item)
		}
		persons = append(persons, person)
	}

	// sqlx expands a slice of structs into a single multi-row INSERT,
	// binding each :name placeholder from the struct's db tags.
	if _, err := j.db.NamedExec(insertPersonQuery, persons); err != nil {
		return gobatch.NewBatchError(gobatch.ErrCodeDbFail, "failed to insert persons batch", err)
	}

	return nil
}
