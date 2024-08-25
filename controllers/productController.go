package controllers

import (
	"net/http"

	"example.com/authentication/initializers"
	"example.com/authentication/models"
	"example.com/authentication/utils"
	"github.com/gin-gonic/gin"
)

type productInput struct {
	Product_name string  `json:"product_name" binding:"required"`
	Price        float64 `json:"price" binding:"required,numeric"`
	Stock        int     `json:"stock" binding:"numeric"`
}

func CreateProduct(c *gin.Context) {
	var productData productInput

	// bind data
	if err := utils.BindData(&productData, c); !err {
		return
	}

	product := models.Product{
		Product_name: productData.Product_name,
		Price:        productData.Price,
		Stock:        productData.Stock,
	}

	if result := initializers.DB.Create(&product); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed create", "error", result.Error.Error(), c)
		return
	}

	utils.ReturnResponse(http.StatusCreated, "ok", "data", product, c)
}

func GetAllProduct(c *gin.Context) {
	var products []models.Product

	if result := initializers.DB.Find(&products); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed get data", "error", result.Error.Error(), c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "data", products, c)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if result := initializers.DB.First(&product, id); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed get data", "error", result.Error.Error(), c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "data", product, c)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var productData productInput

	if err := utils.BindData(&productData, c); !err {
		return
	}

	var product models.Product
	if result := initializers.DB.First(&product, id); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to get data", "error", result.Error.Error(), c)
		return
	}

	// new data
	newData := models.Product{
		Product_name: productData.Product_name,
		Price:        productData.Price,
		Stock:        productData.Stock,
	}

	if result := initializers.DB.Model(&product).Updates(newData); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to update data", "error", result.Error.Error(), c)
		return
	}

	var updatedProduct models.Product
	if result := initializers.DB.First(&updatedProduct, id); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to update data", "error", result.Error.Error(), c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "new_data", updatedProduct, c)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	if result := initializers.DB.Delete(&product, id); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed delete data", "error", result.Error.Error(), c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "", nil, c)
}
