package persistence

import (
	"encoding/json"

	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	database *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) repository.IProductRepository {
	return &ProductRepositoryImpl{database: db}
}

func (r *ProductRepositoryImpl) GetById(c *gin.Context) (*domain.Product, error) {
	var product domain.Product

	if err := r.database.Model(&domain.Product{}).Find(&product, c.Param("id")).Error; err != nil {
		return nil, ErrorProcess
	}

	if product.ID == 0 {
		return nil, ProductNotFound
	}
	return &product, nil
}

func (r *ProductRepositoryImpl) GetAll(c *gin.Context) (*[]domain.Product, error) {
	var products []domain.Product

	if err := r.database.Model(&domain.Product{}).Find(&products).Error; err != nil {
		return nil, ErrorProcess
	}

	return &products, nil
}

func (r *ProductRepositoryImpl) Create(c *gin.Context) (*domain.Product, error) {
	var product domain.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		return nil, ErrorBinding
	}

	if err := r.database.Model(&domain.Product{}).Create(&product).Error; err != nil {
		return nil, ErrorProcess
	}

	data, _ := json.Marshal(&product)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), SpInsert, SpNoPrevData, string(data[:]))
	return &product, nil
}

func (r *ProductRepositoryImpl) Update(c *gin.Context) (*domain.Product, error) {
	patch, product, productOld := map[string]interface{}{}, domain.Product{}, domain.Product{}

	if err := c.Bind(&patch); err != nil {
		return nil, ErrorBinding
	}

	_, errID := patch["id"]

	if !errID {
		return nil, ErrorBinding
	}

	r.database.Model(&domain.Product{}).Find(&productOld, patch["id"])

	if err := r.database.Model(&domain.Product{}).Where("id = ?", patch["id"]).Updates(&patch).Find(&product).Error; err != nil {
		return nil, ErrorProcess
	}

	if product.ID == 0 {
		return nil, ProductNotFound
	}

	old, _ := json.Marshal(&productOld)
	current, _ := json.Marshal(&product)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), SpUpdate, string(old[:]), string(current[:]))
	return &product, nil
}
