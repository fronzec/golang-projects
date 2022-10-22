package main

import (
	"log"

	"fronzec/memkv/metrics"
	"fronzec/memkv/server"
	"fronzec/memkv/store"
)

func main() {
	go metrics.Expose()

	// Creating data store
	ds, err := store.New()
	if err != nil {
		panic(err.Error())
	}
	// Creating http server
	s, err := server.New(ds)
	if err != nil {
		panic(err.Error())
	}
	log.Print("Starting fronzec/memkv")
	go s.Start()
	select {}
}
