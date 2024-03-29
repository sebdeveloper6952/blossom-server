package main

import "log"

func main() {
	storage, err := NewFsStorage("media")
	if err != nil {
		log.Fatal(err)
	}

	hashing, err := NewSha256()
	if err != nil {
		log.Fatal(err)
	}

	server, err := NewServer(
		storage,
		hashing,
	)
	if err != nil {
		log.Fatal(err)
	}

	api := SetupApi("127.0.0.1:8000", server)
	api.Run()
}
