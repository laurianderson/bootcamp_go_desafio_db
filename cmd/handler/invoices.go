package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/laurianderson/bootcamp_go_desafio_db/internal/domain"
	"github.com/laurianderson/bootcamp_go_desafio_db/internal/invoices"
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

func (i *Invoices) LoadJson() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		invoices, err := i.s.LoadJson()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusOK, gin.H{"data": invoices})

	}
}

func (i *Invoices) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := i.s.Update()
		if err != nil {
			return
		}
		return
	}
}


