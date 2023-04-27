package products

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/laurianderson/bootcamp_go_desafio_db/internal/domain"
)

type Service interface {
	Create(product *domain.Product) error
	ReadAll() ([]*domain.Product, error)
	LoadJson() ([]*domain.Product, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(product *domain.Product) error {
	_, err := s.r.Create(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ReadAll() ([]*domain.Product, error) {
	return s.r.ReadAll()
}

func (s *service) LoadJson() (products []*domain.Product, err error) {
	var file []byte
	file, err = os.ReadFile("../datos/products.json")
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(file), &products)
	if err != nil {
		return
	}

	for _, product := range products {
		var id int64
		id, err = s.r.Create(product)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		product.Id = int(id)
	}

	return

}