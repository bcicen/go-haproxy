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
	Addr string
	conn net.Conn
}

func (h *HAProxyClient) RunCommand(cmd string) (*bytes.Buffer, error) {
	err := h.dial()
	if err != nil {
		return nil, err
	}

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
	h.conn.Close()

	if strings.HasPrefix(result.String(), "Unknown command") {
		return nil, fmt.Errorf("Unknown command: %s", cmd)
	}

	return result, nil
}

func (h *HAProxyClient) Stats() (services Services, err error) {
	res, err := h.RunCommand("show stat")
	if err != nil {
		return services, err
	}

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

func (h *HAProxyClient) dial() (err error) {
	switch h.schema() {
	case "unix":
		h.conn, err = net.Dial("unix", strings.Replace(h.Addr, socketSchema, "", 1))
	case "tcp":
		h.conn, err = net.Dial("tcp", strings.Replace(h.Addr, tcpSchema, "", 1))
	default:
		return fmt.Errorf("unknown schema")
	}
	return err
}

func (h *HAProxyClient) schema() string {
	if strings.HasPrefix(h.Addr, socketSchema) {
		return "socket"
	}
	if strings.HasPrefix(h.Addr, tcpSchema) {
		return "tcp"
	}
	return ""
}
