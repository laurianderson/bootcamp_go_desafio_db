package sales

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/laurianderson/bootcamp_go_desafio_db/internal/domain"
)

type Service interface {
	Create(sales *domain.Sales) error
	ReadAll() ([]*domain.Sales, error)
	LoadJson() ([]*domain.Sales, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(sales *domain.Sales) error {
	_, err := s.r.Create(sales)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ReadAll() ([]*domain.Sales, error) {
	return s.r.ReadAll()
}

func (s *service) LoadJson() (sales []*domain.Sales, err error) {
	var file []byte
	file, err = os.ReadFile("../datos/sales.json")
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(file), &sales)
	if err != nil {
		return
	}

	for _, sale := range sales {
		var id int64
		id, err = s.r.Create(sale)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		sale.Id = int(id)
	}

	return

}