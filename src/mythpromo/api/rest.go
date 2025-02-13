package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func loadRoutes(r *gin.Engine) {
	r.GET("/ping", ping)
	r.GET("/products", getProducts)
	// add other handlers here
}

// Request handlers

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong!")
}

func getProducts(c *gin.Context) {
	var req ProductsRequest
	products := make([]ProductResponse, 0)
	err := c.BindQuery(&req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	result, err := manager.FindProducts(req.Category, req.PriceLessThan)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	for _, product := range result {
		productResponse := ProductResponse{
			Sku:      product.Sku,
			Name:     product.Name,
			Category: product.Category,
			Price: Price{
				Original:           product.Price.Original,
				Final:              product.Price.Final,
				DiscountPercentage: strconv.Itoa(product.Price.DiscountPercentage) + "%",
				Currency:           product.Price.Currency,
			},
		}
		products = append(products, productResponse)
	}
	c.JSON(http.StatusOK, products)
}
