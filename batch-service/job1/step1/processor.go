package step1

import (
	"fmt"
	"github.com/chararch/gobatch"
)

type Job1step1Processor struct {
}

func (j *Job1step1Processor) Process(item interface{}, chunkCtx *gobatch.ChunkContext) (interface{}, gobatch.BatchError) {
	//TODO implement me, transform the item into other item
	return fmt.Sprintf("processed-%v", item), nil
}
