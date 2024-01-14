package main

var version string
var app string

func main() {
	server := NewServer(ReadConfigFromEnv())
	server.Serve()
}
