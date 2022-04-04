// stock/service-repository.go
package stock

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type ServiceRepository struct {
	database *gorm.DB
}

func (serviceRepository *ServiceRepository) FindAllServices() []Service {
	var services []Service
	serviceRepository.database.Find(&services)
	return services
}

func (serviceRepository *ServiceRepository) Find(id int) (S, error) {
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

func (repository *ProductRepository) Delete(id int) int64 {
	count := repository.database.Delete(&Product{}, id).RowsAffected
	return count
}

func NewProductRepository(database *gorm.DB) *ProductRepository {
	return &ProductRepository{
		database: database,
	}
}
