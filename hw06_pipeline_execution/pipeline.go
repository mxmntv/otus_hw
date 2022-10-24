package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in

	worker := func(d, in In) Out {
		temp := make(Bi)
		go func() {
			defer close(temp)
			for {
				select {
				case data, ok := <-in:
					if !ok {
						return
					}
					temp <- data
				case <-d:
					return
				}
			}
		}()
		return temp
	}

	for _, s := range stages {
		out = worker(done, s(out))
	}

	return out
}
