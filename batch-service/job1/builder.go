package job1

import (
	"batch-service/job1/step1"
	"github.com/chararch/gobatch"
)

func BuildJob1(partitions uint) gobatch.Job {
	step1Instance := gobatch.NewStep("job1step1").ReadFile(step1.CsvFile).Processor(&step1.Job1step1Processor{}).Writer(&step1.Job1Step1Writer{}).Partitions(partitions).Build()
	return gobatch.NewJob("job1").Step(step1Instance).Build()
}
