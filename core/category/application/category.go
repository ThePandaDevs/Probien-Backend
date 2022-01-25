package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/category/domain"
	"github.com/JairDavid/Probien-Backend/core/category/infrastructure/persistance"
	"github.com/gin-gonic/gin"
)

type CategoryInteractor struct {
}

func (CI *CategoryInteractor) GetById(c *gin.Context) (domain.Category, error) {
	repository := persistance.NewCategoryRepositoryImpl(config.ConnectDB())
	return repository.GetById(c)
}

func (CI *CategoryInteractor) GetAll() ([]domain.Category, error) {
	repository := persistance.NewCategoryRepositoryImpl(config.ConnectDB())
	return repository.GetAll()
}

func (CI *CategoryInteractor) Create(c *gin.Context) (domain.Category, error) {
	repository := persistance.NewCategoryRepositoryImpl(config.ConnectDB())
	return repository.Create(c)
}

func (CI *CategoryInteractor) Delete(c *gin.Context) (domain.Category, error) {
	repository := persistance.NewCategoryRepositoryImpl(config.ConnectDB())
	return repository.Delete(c)
}

func (CI *CategoryInteractor) Update(c *gin.Context) (domain.Category, error) {
	repository := persistance.NewCategoryRepositoryImpl(config.ConnectDB())
	return repository.Update(c)
}
