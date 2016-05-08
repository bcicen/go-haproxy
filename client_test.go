package haproxy

import (
	"testing"
)

// TestSchemaValidation ensures that the schema() method correctly parses Addr strings.
func TestSchemaValidation(t *testing.T) {
	ha := &HAProxyClient{Addr: "tcp://sys49152/"}

	if ha.schema() != "tcp" {
		t.Errorf("Expected 'tcp', received '%s'", ha)
	}

	ha = &HAProxyClient{Addr: "unix://sys2064/"}
	if ha.schema() != "socket" {
		t.Errorf("Expected 'socket', received '%s'", ha)
	}

	ha = &HAProxyClient{Addr: "unknown://RUN/"}
	if ha.schema() != "" {
		t.Errorf("Expected '', received '%s'", ha)
	}

}
