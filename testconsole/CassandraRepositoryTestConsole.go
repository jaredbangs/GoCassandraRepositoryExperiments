package main

import (
	"fmt"
	"github.com/jaredbangs/GoCassandraRepositoryExperiments/cassandrarepository"
)

func main() {

	repo := cassandrarepository.NewRepository("127.0.0.1")

	keyspaceNames := repo.ListKeyspaces()

	for _, name := range keyspaceNames {
		fmt.Println(name)
	}
}
