package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/sagikazarmark/bingo"
)

type Stat struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
	Count int       `json:"count"`
}

func NewStat() Stat {
	t := time.Now().UTC()

	return Stat{
		Start: t.Add(-1 * 24 * time.Hour),
		End:   t,
		Count: 10,
	}
}

func Stats(bin *bingo.Bin) {
	group, _ := bingo.NewGroup("Stats", "")

	endpoint, _ := bingo.NewEndpoint(
		"GET",
		"/v1/stats/:username",
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			if p.ByName("username") == "john.doe" {
				handleStat(w, r)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		},
	)
	endpoint.Description = "Returns statistics for a specific user"
	endpoint.Parameters.Set("username", "john.doe")
	group.AddEndpoint(endpoint)

	endpoint, _ = bingo.NewEndpoint(
		"GET",
		"/v1/stats/",
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			handleStat(w, r)
		},
	)
	endpoint.Description = "Returns statistics"
	group.AddEndpoint(endpoint)

	bin.AddGroup(group)
}

func handleStat(w http.ResponseWriter, r *http.Request) {
	stat := NewStat()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(stat); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
