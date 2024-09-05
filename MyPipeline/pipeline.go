package MyPipeline

type StageProcessing func(<-chan bool, <-chan int) <-chan int
 
type PipeLine struct {
    stages []StageProcessing
    done   <-chan bool
}
 
func NewPipeline(done <-chan bool, stages ...StageProcessing) *PipeLine {
    return &PipeLine{done: done, stages: stages}
}

func (p *PipeLine) Run(source <-chan int) <-chan int {
    var c <-chan int = source
    for index := range p.stages {
        c = p.runStage(p.stages[index], c)
    }
    return c
}

func (p *PipeLine) runStage(stage StageProcessing, sourceChan <-chan int) <-chan int {
    return stage(p.done, sourceChan)
}