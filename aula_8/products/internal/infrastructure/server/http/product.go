package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jhonyzam/curso_golang/aula_8/products/domain"
	"net/http"
	"strconv"
)

func (h *handler) postProduct(c *gin.Context) {
	product := &domain.Product{}
	if err := c.BindJSON(&product); err != nil {
		return
	}

	if !h.productService.IsValid(product) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	product, err := h.productService.Create(product)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *handler) getProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	product, err := h.productService.Get(productID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *handler) updateProduct(c *gin.Context) {
	product := &domain.Product{}
	if err := c.BindJSON(&product); err != nil {
		return
	}

	product.ID, _ = strconv.Atoi(c.Param("id"))

	if !h.productService.IsValid(product) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := h.productService.Update(product)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusNoContent, product)
}

func (h *handler) deleteProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	if err := h.productService.Delete(productID); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
