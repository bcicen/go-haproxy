package haproxystats

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/gocarina/gocsv"
)

const (
	socketSchema = "unix:///"
	tcpSchema    = "tcp://"
)

type HAProxyClient struct {
	conn net.Conn
}

func (h *HAProxyClient) RunCommand(cmd string) *bytes.Buffer {
	done := make(chan bool)
	result := bytes.NewBuffer(nil)

	go func() {
		io.Copy(result, h.conn)
		defer func() { done <- true }()
	}()

	go func() {
		h.conn.Write([]byte(cmd + "\n"))
		defer func() { done <- true }()
	}()

	// Wait for both io streams to close
	for i := 0; i < 2; i++ {
		select {
		case <-done:
		}
	}

	return result
}

func (h *HAProxyClient) Stats() (services Services, err error) {
	res := h.RunCommand("show stat")

	allStats := []*Stat{}
	reader := csv.NewReader(res)
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

func New(addr string) (*HAProxyClient, error) {
	var err error
	client := &HAProxyClient{}

	if strings.HasPrefix(addr, socketSchema) {
		client.conn, err = net.Dial("unix", strings.Replace(addr, socketSchema, "", 1))
		if err != nil {
			panic(err)
		}
	}

	if strings.HasPrefix(addr, tcpSchema) {
		client.conn, err = net.Dial("tcp", strings.Replace(addr, tcpSchema, "", 1))
		if err != nil {
			panic(err)
		}
	}

	if client.conn == nil {
		return nil, fmt.Errorf("unknown schema")
	}

	return client, nil
}
