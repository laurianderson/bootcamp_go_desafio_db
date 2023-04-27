package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laurianderson/bootcamp_go_desafio_db/internal/customers"
	"github.com/laurianderson/bootcamp_go_desafio_db/internal/domain"
)

type Customers struct {
	s customers.Service
}

func NewHandlerCustomers(s customers.Service) *Customers {
	return &Customers{s}
}

func (c *Customers) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customers, err := c.s.ReadAll()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, customers)
	}
}

func (c *Customers) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customer := domain.Customers{}
		err := ctx.ShouldBindJSON(&customer)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		err = c.s.Create(&customer)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, gin.H{"data": customer})
	}
}

func (c *Customers) LoadJson() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		customers, err := c.s.LoadJson()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusOK, gin.H{"data": customers})

	}
}
