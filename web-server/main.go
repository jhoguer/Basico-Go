package main

func main() {
	server := NewServer(":3001")
	server.Handle("/", HandlerRoot)
	server.Handle("/api", HandlerHome)
	server.Listen()
}
