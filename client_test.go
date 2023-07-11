package haproxy_test

import (
	"fmt"
	"github.com/vponomarev/go-haproxy"
)

const (
	HAPROXY_SOCKET_ADDR = "unix:///var/run/haproxy.sock"
)

func ExampleHAProxyClient_Stats() {
	client := &haproxy.HAProxyClient{
		Addr:      HAPROXY_SOCKET_ADDR,
		Timeout:   30,
		TimeoutOp: 5,
	}
	stats, err := client.Stats()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, s := range stats {
		fmt.Printf("%s: %s\n", s.SvName, s.Status)
	}
	// Output:
	//static: DOWN
	//app1: UP
	//app2: UP
	//...
}

func ExampleHAProxyClient_Info() {
	client := &haproxy.HAProxyClient{
		Addr: HAPROXY_SOCKET_ADDR,
	}
	info, err := client.Info()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s version %s\n", info.Name, info.Version)
	// Output:
	//HAProxy version 2.8.1
}

func ExampleHAProxyClient_RunCommand() {
	client := &haproxy.HAProxyClient{
		Addr: HAPROXY_SOCKET_ADDR,
	}
	result, err := client.RunCommand("show info")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result.String())
}
