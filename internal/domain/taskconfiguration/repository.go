package taskconfiguration

type Repository interface {
	Create(t *TaskConfiguration) error
	GetAll() ([]TaskConfiguration, error)
	GetById(id string) (*TaskConfiguration, error)
}
