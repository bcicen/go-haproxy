package haproxy

import (
	"bytes"
	"testing"
)

type InfoTestHAProxyClient struct{}

// RunCommand stubs the HAProxyClient returning our expected bytes.Buffer containing the response from a 'show info' command.
func (ha *InfoTestHAProxyClient) RunCommand(cmd string) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	buf.WriteString("Name: HAProxy\n")
	buf.WriteString("Version: 1.5.4\n")
	buf.WriteString("node: SYS64738\n")
	buf.WriteString("description: go-haproxy stub tests.\n")
	return &buf, nil
}

// TestCommandInfo validates the structure of the "show info" command is handled appropriately.
func TestCommandInfo(t *testing.T) {
	ha := new(InfoTestHAProxyClient)
	info, err := Info(ha)

	if err != nil {
		t.Fatalf("Unable to execute 'show info' Info()")
	}

	expect := "HAProxy"
	if info.Name != expect {
		t.Errorf("Expected Name of '%s', but received '%s' instead", expect, info.Name)
	}

	expect = "1.5.4"
	if info.Version != expect {
		t.Errorf("Expected Version of '%s', but received '%s' instead", expect, info.Version)
	}

	expect = "SYS64738"
	if info.Node != expect {
		t.Errorf("Expected Node of '%s', but received '%s' instead", expect, info.Node)
	}

	expect = "go-haproxy stub tests."
	if info.Description != expect {
		t.Errorf("Expected Description of '%s', but received '%s' instead", expect, info.Description)
	}
}
