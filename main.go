package haproxystats

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
)

type HAProxyStats struct {
	URI      string
	interval time.Duration
	client   *http.Client
	Up       bool
	Fields   []string
}

func (h *HAProxyStats) Poll() (Services, error) {
	var services Services
	resp, err := h.client.Get(h.URI)
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

		fmt.Println(s.SvName)
	}
	fmt.Println(len(services.Frontends))
	fmt.Println(len(services.Backends))
	fmt.Println(len(services.Listeners))

	return services, nil

	//	for {
	//		row, err := reader.Read()
	//		switch err {
	//		case nil:
	//		case io.EOF:
	//			return
	//		case err.(*csv.ParseError):
	//			fmt.Printf("csv parse error: %v", err)
	//			return
	//		default:
	//			fmt.Printf("unexpected error: %v", err)
	//			return
	//		}

	// read metric names from header row
	//		if len(h.Fields) < 1 {
	//			h.readHeader(row)
	//			continue
	//		}
	//
	//		// zip remaining columns with fieldnames
	//		item := make(map[string]string)
	//		for idx, col := range row {
	//			item[h.Fields[idx]] = col
	//		}
	//
	//		j, err := json.Marshal(item)
	//		j, err := json.Unmarshal(item)
	//		if err != nil {
	//			panic(err)
	//		}
	//		fmt.Println(string(j))
	//	}
}

func (h *HAProxyStats) Run() {
	go func() {
		h.Poll()
		time.Sleep(h.interval * time.Second)
	}()
}

func (h *HAProxyStats) readHeader(row []string) {
	r := strings.NewReplacer("#", "", " ", "")
	for _, col := range row {
		h.Fields = append(h.Fields, r.Replace(col))
	}
}

func NewHAProxyStats(hostAddr string, interval, timeout time.Duration) *HAProxyStats {
	return &HAProxyStats{
		URI:      hostAddr + "/;csv;norefresh",
		interval: interval,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}
