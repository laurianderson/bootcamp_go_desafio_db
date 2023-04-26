package handler

import (
	"github.com/laurianderson/bootcamp_go_desafio_db/internal/domain"
	"github.com/laurianderson/bootcamp_go_desafio_db/internal/invoices"
	"github.com/gin-gonic/gin"
)

type Invoices struct {
	s invoices.Service
}

func NewHandlerInvoices(s invoices.Service) *Invoices {
	return &Invoices{s}
}

func (i *Invoices) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		invoices, err := i.s.ReadAll()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, invoices)
	}
}

func (i *Invoices) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		invoices := domain.Invoices{}
		err := ctx.ShouldBindJSON(&invoices)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		err = i.s.Create(&invoices)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, gin.H{"data": invoices})
	}
}
