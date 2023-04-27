package invoices

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/laurianderson/bootcamp_go_desafio_db/internal/domain"
)

type Service interface {
	Create(invoices *domain.Invoices) error
	ReadAll() ([]*domain.Invoices, error)
	LoadJson() ([]*domain.Invoices, error)
	Update() error

}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(invoices *domain.Invoices) error {
	_, err := s.r.Create(invoices)
	if err != nil {
		return err
	}
	return nil

}

func (s *service) ReadAll() ([]*domain.Invoices, error) {
	return s.r.ReadAll()
}

func (s *service) LoadJson() (invoices []*domain.Invoices, err error) {
	var file []byte
	file, err = os.ReadFile("../datos/invoices.json")
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(file), &invoices)
	if err != nil {
		return
	}

	for _, invoice := range invoices {
		var id int64
		id, err = s.r.Create(invoice)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		invoice.Id = int(id)
	}

	return

}

func (s *service) Update() error {
	return s.r.Update()
}