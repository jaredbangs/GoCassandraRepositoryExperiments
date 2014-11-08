package cassandrarepository

import (
	"github.com/gocql/gocql"
)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (repository *Repository) ListKeyspaces() *gocql.Iter {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "system"
	cluster.Consistency = gocql.Quorum

	session, _ := cluster.CreateSession()
	defer session.Close()

	return session.Query(`SELECT keyspace_name FROM schema_keyspaces`).Iter()
}
