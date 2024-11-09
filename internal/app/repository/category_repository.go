package repository

import (
	"github.com/project-app-inventaris/internal/model"
	"github.com/project-app-inventaris/internal/model/dto"
	"github.com/project-app-inventaris/utils/common"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryRepository interface {
	BaseRepository[model.Category]
	Paging(requestPaging dto.PaginationParam, queries ...string) ([]*model.Category, *dto.Paging, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(payload *model.Category) (*model.Category, error) {
	category := model.Category{
		ID:          payload.ID,
		Name:        payload.Name,
		Description: payload.Description,
	}

	if err := r.db.Create(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) Get(id string) (*model.Category, error) {
	category := model.Category{}
	query := `id = ?`

	if err := r.db.Where(query, id).First(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepository) Paging(requestPaging dto.PaginationParam, queries ...string) ([]*model.Category, *dto.Paging, error) {

	categorys := []*model.Category{}

	paginationQuery := common.GetPaginationParams(requestPaging)

	var totalRows int64

	if err := r.db.Limit(paginationQuery.Take).Offset(paginationQuery.Skip).Find(&categorys).Error; err != nil {
		return nil, nil, err
	}

	var count int = int(totalRows)

	return categorys, common.Paginate(paginationQuery.Take, paginationQuery.Page, count), nil
}

func (r *categoryRepository) Update(id string, payload *model.Category) (*model.Category, error) {
	category := model.Category{}
	query := `id = ?`

	result := r.db.Model(&category).Where(query, id).Clauses(clause.Returning{}).Updates(model.Category{
		ID:          id,
		Name:        payload.Name,
		Description: payload.Description,
	})

	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}

func (r *categoryRepository) Delete(id string) (*model.Category, error) {
	category := model.Category{}
	query := `id = ?`

	if err := r.db.Clauses(clause.Returning{}).Where(query, id).Delete(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}
