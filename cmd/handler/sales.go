package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laurianderson/bootcamp_go_desafio_db/internal/domain"
	"github.com/laurianderson/bootcamp_go_desafio_db/internal/sales"
)

type Sales struct {
	s sales.Service
}

func NewHandlerSales(s sales.Service) *Sales {
	return &Sales{s}
}

func (s *Sales) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		invoices, err := s.s.ReadAll()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, invoices)
	}
}

func (s *Sales) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sale := domain.Sales{}
		err := ctx.ShouldBindJSON(&sale)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		err = s.s.Create(&sale)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, gin.H{"data": sale})
	}
}

func (s *Sales) LoadJson() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		sales, err := s.s.LoadJson()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusOK, gin.H{"data": sales})

	}
}