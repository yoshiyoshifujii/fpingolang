package parallelism

type (
	GoExecutor struct{}
)

func (g GoExecutor) Submit(task func()) {
	go task()
}
