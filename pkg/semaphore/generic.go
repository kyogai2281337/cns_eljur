package semaphore

type Semaphore chan struct{}

func NewSemaphore(n int) Semaphore {
	return make(Semaphore, n)
}
func (s Semaphore) Acquire(n int) {
	e := struct{}{}
	for i := 0; i < n; i++ {
		s <- e
	}
}
func (s Semaphore) Release(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}
