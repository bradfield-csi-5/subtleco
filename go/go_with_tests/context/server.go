package context

import (
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}

type Store interface {
	Fetch() string
	Cancel()
}

type StubStore struct {
	response  string
	cancelled bool
}

func (s *StubStore) Fetch() string {
	return s.response
}

func (s *StubStore) Cancel() {
	s.cancelled = true
}
