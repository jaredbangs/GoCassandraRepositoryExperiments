package cassandrarepository

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (repository *Repository) ListKeySpaces() string {
	return "Hello"
}
