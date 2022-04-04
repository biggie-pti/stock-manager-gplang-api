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

func (serviceRepository *ServiceRepository) FindService(id int) (Service, error) {
	var service Service
	err := serviceRepository.database.Find(&service, id).Error
	if service.Name == "" {
		err = errors.New("Service not found")
	}
	return s, err
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
