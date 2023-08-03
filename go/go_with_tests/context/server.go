package context

import (
	"context"
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			return // ideally with an error log
		}
		fmt.Fprint(w, data)
	}
}

type Store interface {
	Fetch(ctx context.Context) (string, error)
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
