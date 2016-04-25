// Package haproxy provides a minimal client for communicating with, and issuing commands to, HAproxy over a network or file socket.
package haproxy

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"strings"
)

const (
	socketSchema = "unix:///"
	tcpSchema    = "tcp://"
)

// HAProxyClient is the main structure of the library.
type HAProxyClient struct {
	Addr string
	conn net.Conn
}

// Entrypoint to the client. Sends an arbitray command string to HAProxy.
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
