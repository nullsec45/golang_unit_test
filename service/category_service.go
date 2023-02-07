package service

import (
	"errors"
	"github.com/nullsec45/golang_unit_test.git/entity"
	"github.com/nullsec45/golang_unit_test.git/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)

	if category == nil {
		return category, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}
