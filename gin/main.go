package main

import (
	"github.com/amoniacou/meditation6/gin/server"
)

func main() {
	session := server.NewSession("persons")
	server := server.NewServer(session)
	server.Run(":8080")
}
