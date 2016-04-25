[![GoDoc](https://godoc.org/github.com/bcicen/go-haproxy?status.svg)](https://godoc.org/github.com/bcicen/go-haproxy)

# go-haproxy
Go library for managing and communicating with HAProxy 

## Usage

Initialize a client object. Supported address schemas are `tcp://` and `unix:///`
```go
client := &haproxy.HAProxyClient{
  Addr: "tcp://localhost:9999",
}
```

Fetch results for a built in command(currently only supports `stats`):
```go
stats, err := client.Stats()
for _, i := range stats {
	fmt.Printf("%s: %s\n", i.SvName, i.Status)
}
```

Or retrieve the result body from an arbitrary command string:
```go
result, err := h.RunCommand("show info")
```
