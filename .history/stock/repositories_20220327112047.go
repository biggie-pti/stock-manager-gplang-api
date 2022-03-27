// stock/repositories.go
package stock

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type ProductRepository struct {
	database *gorm.DB
}

func (repository *ProductRepository) FindAll() []Product {
	var products []Product
	repository.database.Find(&products)
	return products
}

func (repository *ProductRepository) Find(id int) (Product, error) {
	var product Product
	err := repository.database.Find(&product, id).Error
	if product.Name == "" {
		err = errors.New("Product not found")
	}
	return product, err
}

func (repository *ProductRepository) Create(product Todo) (Todo, error) {
	err := repository.database.Create(&todo).Error
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (repository *TodoRepository) Save(user Todo) (Todo, error) {
	err := repository.database.Save(user).Error
	return user, err
}

func (repository *TodoRepository) Delete(id int) int64 {
	count := repository.database.Delete(&Todo{}, id).RowsAffected
	return count
}

func NewTodoRepository(database *gorm.DB) *TodoRepository {
	return &TodoRepository{
		database: database,
	}
}
