package ports

type UserRepository interface {
	Get(int) (map[string]interface{}, error)
	GetAll() ([]map[string]interface{}, error)
	Insert(int, string) error
	Delete(int) error
	UpdateName(int, string) error
}
