package haproxystats

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
)

type StatsClient struct {
	uri    string
	client *http.Client
	Up     bool
}

func (h *StatsClient) Fetch() (Services, error) {
	var services Services

	resp, err := h.client.Get(h.uri)
	if err != nil {
		h.Up = false
		return services, err
	}

	reader := csv.NewReader(resp.Body)
	reader.TrailingComma = true

	allStats := []*Stat{}
	err = gocsv.UnmarshalCSV(reader, &allStats)
	if err != nil {
		panic(err)
	}

	for _, s := range allStats {
		switch s.SvName {
		case "FRONTEND":
			services.Frontends = append(services.Frontends, s)
		case "BACKEND":
			services.Backends = append(services.Backends, s)
		default:
			services.Listeners = append(services.Listeners, s)
		}
	}

	return services, nil
}

func New(hostAddr string, timeout time.Duration) *StatsClient {
	return &StatsClient{
		uri: hostAddr + "/;csv;norefresh",
		client: &http.Client{
			Timeout: timeout,
		},
	}
}
