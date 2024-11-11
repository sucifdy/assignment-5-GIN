package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
)

type CategoryService interface {
	Store(category *model.Category) error
	Update(id int, category model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]model.Category, error)
}

type categoryService struct {
	categoryRepository repo.CategoryRepository
}

func NewCategoryService(categoryRepository repo.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository}
}

func (c *categoryService) Store(category *model.Category) error {
	return c.categoryRepository.Store(category)
}

func (c *categoryService) Update(id int, category model.Category) error {
	// Memperbarui kategori berdasarkan ID
	return c.categoryRepository.Update(id, category)
}

func (c *categoryService) Delete(id int) error {
	// Menghapus kategori berdasarkan ID
	return c.categoryRepository.Delete(id)
}

func (c *categoryService) GetByID(id int) (*model.Category, error) {
	// Mengambil kategori berdasarkan ID
	return c.categoryRepository.GetByID(id)
}

func (c *categoryService) GetList() ([]model.Category, error) {
	// Mengambil semua kategori
	return c.categoryRepository.GetList()
}
