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

func (repository *ProductRepository) Create(product Product) (Product, error) {
	err := repository.database.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (repository *ProductRepository) Save(user Product) (Product, error) {
	err := repository.database.Save(user).Error
	return user, err
}

func (repository *ProductRepository) DeleteSer(id int) int64 {
	count := repository.database.Delete(&Product{}, id).RowsAffected
	return count
}

func NewProductRepository(database *gorm.DB) *ProductRepository {
	return &ProductRepository{
		database: database,
	}
}
