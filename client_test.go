package haproxy

func ExampleHAProxyClient_Stats() {
	client := &HAProxyClient{
		Addr: "unix:///var/run/haproxy.sock",
	}
	stats, err := client.Stats()
	for _, s := range stats {
		fmt.Printf("%s: %s\n", s.SvName, s.Status)
	}
	// Output:
	//	static: DOWN
	//	app1: UP
	//	app2: UP
	//	...
}

func ExampleHAProxyClient_Info() {
	client := &HAProxyClient{
		Addr: "unix:///var/run/haproxy.sock",
	}
	info, err := client.Info()
	fmt.Printf("%s version %s\n", info.Name, info.Version)
	// Output:
	//	HAProxy version 1.6.3
}

func ExampleHAProxyClient_RunCommand() {
	client := &HAProxyClient{
		Addr: "unix:///var/run/haproxy.sock",
	}
	result, err := h.RunCommand("show info")
	fmt.Println(result.String())
	// Output:
	//	Name: HAProxy
	//	Version: 1.6.3
	//	Release_date: 2015/12/25
	//	Nbproc: 1
	//	...
}
