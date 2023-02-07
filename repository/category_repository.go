package repository

import "github.com/nullsec45/golang_unit_test.git/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
