package cassandrarepository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListKeyspaces(t *testing.T) {

	repository := NewRepository()

	repository.ListKeySpaces()

	assert.NotNil(t, repository, "obj shouldn't be nil")
}
