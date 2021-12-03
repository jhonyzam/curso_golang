package http

import (
	"github.com/jhonyzam/curso_golang/aula_8/products/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	productService domain.ProductService
}

func NewHandler(productService domain.ProductService) http.Handler {
	handler := &handler{
		productService: productService,
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	v1 := router.Group("/v1")

	v1.POST("/products", handler.postProduct)
	v1.GET("/products/:id", handler.getProduct)
	v1.PUT("/products/:id", handler.updateProduct)
	v1.DELETE("/products/:id", handler.deleteProduct)

	return router
}
