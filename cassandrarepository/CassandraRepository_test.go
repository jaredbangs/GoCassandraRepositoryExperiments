package cassandrarepository

import (
	"github.com/gocql/gocql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSession(t *testing.T) {

	repository := NewRepository("127.0.0.1")

	session := repository.createSession("system", gocql.Quorum)

	assert.NotNil(t, session, "session shouldn't be nil")
}

func TestListKeyspaces(t *testing.T) {

	repository := NewRepository("127.0.0.1")

	keyspaceNames := repository.ListKeyspaces()

	assert.NotNil(t, keyspaceNames, "slice shouldn't be nil")

	assert.True(t, len(keyspaceNames) > 0, "slice shouldn't be empty")
}
