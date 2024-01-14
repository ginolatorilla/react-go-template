package main

var version string
var app string

func main() {
	server := NewServer(":8081")
	server.Serve()
}
