package main

import (
	"fmt"
	"github.com/jaredbangs/GoCassandraRepositoryExperiments/cassandrarepository"
	"log"
)

func main() {

	repo := cassandrarepository.NewRepository()

	iter := repo.ListKeyspaces()

	var name string

	for iter.Scan(&name) {
		fmt.Println(name)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}
