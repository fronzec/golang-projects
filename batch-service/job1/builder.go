package job1

import (
	"batch-service/job1/step1"
	"github.com/chararch/gobatch"
	"github.com/jmoiron/sqlx"
)

func BuildJob1(partitions uint, connection *sqlx.DB) gobatch.Job {
	step1Instance := gobatch.NewStep("job1step1").ReadFile(step1.CsvFile).Processor(&step1.Job1step1Processor{}).Writer(step1.NewJob1Step1Writer(connection)).Partitions(partitions).Build()
	return gobatch.NewJob("job1").Step(step1Instance).Build()
}
