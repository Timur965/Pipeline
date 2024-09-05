package main

import (
	"Pipeline/Consumer"
	"Pipeline/Input"
	"Pipeline/MyPipeline"
	"Pipeline/Stage"
)

func main() {
	source, done := Input.InputData()

	pipeline := MyPipeline.NewPipeline(done, Stage.FilterNegativeNumber, Stage.FilterNotMultiples3, Stage.Buffering)

	Consumer.Output(done, pipeline.Run(source))

	select{}
}
