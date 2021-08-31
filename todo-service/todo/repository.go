package todo

type Repository interface {
	SaveOrUpdate(todo TodoEntity) (TodoEntity, error)
	GetByID(id int) (TodoEntity, error)
	DeleteByID(id int) (int, error)
}
