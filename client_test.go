package haproxy

import (
	"fmt"
)

func ExampleHAProxyClient_Stats() {
	client := &HAProxyClient{
		Addr: "unix:///var/run/haproxy.sock",
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
	client := &HAProxyClient{
		Addr: "unix:///var/run/haproxy.sock",
	}
	info, err := client.Info()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s version %s\n", info.Name, info.Version)
	// Output:
	//HAProxy version 1.6.3
}

func ExampleHAProxyClient_RunCommand() {
	client := &HAProxyClient{
		Addr: "unix:///var/run/haproxy.sock",
	}
	result, err := client.RunCommand("show info")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result.String())
}
