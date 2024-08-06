package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rocketseat-education/semana-tech-go-react-server/internal/store/pgstore"
)

type apihandler struct {
	q *pgstore.Queries
	r *chi.Mux
}

func (h apihandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	a := apihandler{
		q: q,
	}

	r := chi.NewRouter()

	a.r = r
	return a
}
