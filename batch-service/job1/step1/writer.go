package step1

import (
	"fmt"
	"github.com/chararch/gobatch"
)

type Job1Step1Writer struct{}

func (j *Job1Step1Writer) Write(items []interface{}, chunkCtx *gobatch.ChunkContext) gobatch.BatchError {
	fmt.Printf("write: %v\n", items)
	return nil
}
