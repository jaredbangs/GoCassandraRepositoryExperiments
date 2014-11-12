package cassandrarepository

import (
	"github.com/gocql/gocql"
	"log"
)

type Repository struct {
	Hosts []string
}

func NewRepository(hosts ...string) *Repository {
	repository := &Repository{
		Hosts: hosts,
	}
	return repository
}

func (repository *Repository) ListKeyspaces() []string {

	session := repository.createSession("system", gocql.Quorum)
	defer session.Close()

	iter := session.Query(`SELECT keyspace_name FROM schema_keyspaces`).Iter()

	keyspaceNames := []string{}

	var name string
	for iter.Scan(&name) {
		keyspaceNames = append(keyspaceNames, name)
	}

	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	return keyspaceNames
}

func (repository *Repository) createSession(keyspaceName string, consistency gocql.Consistency) *gocql.Session {
	cluster := gocql.NewCluster(repository.Hosts...)
	cluster.Keyspace = keyspaceName
	cluster.Consistency = consistency

	session, _ := cluster.CreateSession()
	return session
}
