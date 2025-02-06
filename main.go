package main

import (
	"log"
)

func main() {

	server := NewServer(":4333")
	server.Route("/user", func(w httpWriter) {
		w.WriteHeader(200)
		w.Write([]byte("Hello user "))

	})
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
