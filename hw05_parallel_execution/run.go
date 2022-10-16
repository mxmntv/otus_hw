package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type state struct {
	mu      sync.Mutex
	counter int
	run     bool
}

func (s *state) incr() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.counter++
	return s.counter
}

func (s *state) setflag() {
	s.mu.Lock()
	s.run = !s.run
	s.mu.Unlock()
}

func (s *state) getflag() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.run
}

func worker(t <-chan Task, n, m int) <-chan int {
	s := &state{
		mu:      sync.Mutex{},
		counter: 0,
		run:     true,
	}
	wg := &sync.WaitGroup{}

	errors := make(chan int, 1)
	defer close(errors)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(s *state) {
			defer wg.Done()
			for v := range t {
				if s.getflag() {
					if err := v(); err != nil {
						c := s.incr()
						if c == m {
							errors <- c
							s.setflag()
							break
						}
					}
				} else {
					return
				}
			}
		}(s)
	}

	wg.Wait()
	return errors
}

func Run(tasks []Task, n, m int) error {
	task := func() <-chan Task {
		t := make(chan Task, len(tasks))
		defer close(t)
		for _, v := range tasks {
			t <- v
		}
		return t
	}

	t := task()
	w := worker(t, n, m)

	for range w {
		return ErrErrorsLimitExceeded
	}

	return nil
}
