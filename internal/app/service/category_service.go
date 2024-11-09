package service

import (
	"github.com/project-app-inventaris/internal/app/repository"
	"github.com/project-app-inventaris/internal/model"
	"github.com/project-app-inventaris/internal/model/dto"
	"github.com/project-app-inventaris/utils/exception"
	"gorm.io/gorm"
)

type CategoryService interface {
	AddNewCategory(payload *model.Category) (*dto.CategoryResponse, error)
	FindCategoryByID(id string) (*dto.CategoryResponse, error)
	FindAllCategory(requestPaging dto.PaginationParam, queries ...string) ([]*dto.CategoryResponse, *dto.Paging, error)
	UpdateCategoryByID(id string, payload *model.Category) (*dto.CategoryResponse, error)
	RemoveCategory(id string) (*dto.CategoryResponse, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) AddNewCategory(payload *model.Category) (*dto.CategoryResponse, error) {

	category, err := s.repo.Create(payload)

	categoryResponse := dto.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &categoryResponse, err
}

func (s *categoryService) FindCategoryByID(id string) (*dto.CategoryResponse, error) {

	category, err := s.repo.Get(id)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	categoryResponse := dto.CategoryResponse{
		ID:          id,
		Name:        category.Name,
		Description: category.Description,
	}

	return &categoryResponse, err
}

func (s *categoryService) FindAllCategory(requestPaging dto.PaginationParam, queries ...string) ([]*dto.CategoryResponse, *dto.Paging, error) {

	categorys, paging, err := s.repo.Paging(requestPaging, queries...)

	if err != nil {
		return nil, nil, gorm.ErrRecordNotFound
	}

	var categoryResponses []*dto.CategoryResponse

	for _, category := range categorys {

		categoryResponse := dto.CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}

		categoryResponses = append(categoryResponses, &categoryResponse)
	}

	return categoryResponses, paging, err
}

func (s *categoryService) UpdateCategoryByID(id string, payload *model.Category) (*dto.CategoryResponse, error) {

	category, err := s.repo.Get(id)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	category, err = s.repo.Update(category.ID, payload)

	if err != nil {
		return nil, exception.ErrFailedUpdate
	}

	categoryResponse := dto.CategoryResponse{
		ID:          id,
		Name:        category.Name,
		Description: category.Description,
	}

	return &categoryResponse, err
}

func (s *categoryService) RemoveCategory(id string) (*dto.CategoryResponse, error) {

	category, err := s.repo.Get(id)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	category, err = s.repo.Delete(category.ID)

	if err != nil {
		return nil, exception.ErrFailedDelete
	}

	categryResponse := dto.CategoryResponse{
		ID:          id,
		Name:        category.Name,
		Description: category.Description,
	}

	return &categryResponse, err
}
