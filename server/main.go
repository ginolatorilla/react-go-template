package main

var version string

func main() {
	server := NewServer(":8081")
	server.Serve()
}
