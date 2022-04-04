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
	return service, err
}

func (serviceRepository *ServiceRepository) AddService(service Service) (Service, error) {
	err := serviceRepository.database.Create(&service).Error
	if err != nil {
		return service, err
	}

	return service, nil
}



func (serviceRepository *ServiceRepository) DeleteService(id int) int64 {
	count := serviceRepository.database.Delete(&Service{}, id).RowsAffected
	return count
}

func NewRepository(database *gorm.DB) *ProductRepository {
	return &ProductRepository{
		database: database,
	}
}
