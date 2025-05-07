package storage

import "github.com/atindraraut/crudgo/internal/types"
type Storage interface {
	CreateStudent(name string, age int,email string) (int64 , error)
	GetStudentById(id int64) (types.Student, error)
}

