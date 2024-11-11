package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	"fmt"
)

type CategoryRepository interface {
	Store(category *model.Category) error
	Update(id int, category model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]model.Category, error)
}

type categoryRepository struct {
	filebasedDb *filebased.Data
}

func NewCategoryRepo(filebasedDb *filebased.Data) *categoryRepository {
	return &categoryRepository{filebasedDb}
}

func (c *categoryRepository) Store(category *model.Category) error {
	// Simpan kategori menggunakan StoreCategory dari filebased
	return c.filebasedDb.StoreCategory(*category)
}

func (c *categoryRepository) Update(id int, category model.Category) error {
	// Ambil kategori berdasarkan ID untuk memastikan ada
	existingCategory, err := c.filebasedDb.GetCategoryByID(id)
	if err != nil {
		return fmt.Errorf("category not found: %w", err)
	}

	// Update field yang diperlukan
	existingCategory.Name = category.Name

	// Simpan kembali kategori yang sudah diperbarui
	return c.filebasedDb.StoreCategory(*existingCategory)
}

func (c *categoryRepository) Delete(id int) error {
	// Hapus kategori berdasarkan ID menggunakan DeleteCategory dari filebased
	return c.filebasedDb.DeleteCategory(id)
}

func (c *categoryRepository) GetByID(id int) (*model.Category, error) {
	// Ambil kategori berdasarkan ID menggunakan GetCategoryByID dari filebased
	return c.filebasedDb.GetCategoryByID(id)
}

func (c *categoryRepository) GetList() ([]model.Category, error) {
	// Ambil semua kategori menggunakan GetCategories dari filebased
	return c.filebasedDb.GetCategories()
}
