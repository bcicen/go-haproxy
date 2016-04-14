package haproxystats

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"time"

	"github.com/gocarina/gocsv"
)

type StatsClient struct {
	uri    string
	client *http.Client
}

func (h *StatsClient) Fetch() (Services, error) {
	var services Services

	resp, err := h.client.Get(h.uri)
	if err != nil {
		return services, fmt.Errorf("fetch errror: %s", err)
	}

	allStats := []*Stat{}
	reader := csv.NewReader(resp.Body)
	reader.TrailingComma = true
	err = gocsv.UnmarshalCSV(reader, &allStats)
	if err != nil {
		return services, fmt.Errorf("error reading csv: %s", err)
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
