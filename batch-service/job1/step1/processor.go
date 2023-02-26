package step1

import (
	"github.com/chararch/gobatch"
	"time"
)

type PersonV1Entity struct {
	ID int `db:"id"`
	Person
	Processed     bool      `db:"processed"`
	CreatedAt     time.Time `db:"created_at"`
	LastUpdatedAt time.Time `db:"last_updated_at"`
}

type Job1step1Processor struct {
}

func (j *Job1step1Processor) Process(item interface{}, chunkCtx *gobatch.ChunkContext) (interface{}, gobatch.BatchError) {
	person := item.(*Person)
	newPerson := PersonV1Entity{
		Person:        *person,
		CreatedAt:     time.Now().UTC(),
		LastUpdatedAt: time.Now().UTC(),
	}
	return newPerson, nil
}
