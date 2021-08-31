package todo

import "time"

type TodoDTO struct {
	Description string `json:"description" db:"description"`
	Completed   bool   `json:"completed" db:"completed"`
}

type TodoEntity struct {
	id int `db:"id"`
	TodoDTO
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
