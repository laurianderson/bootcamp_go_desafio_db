package customers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/laurianderson/bootcamp_go_desafio_db/internal/domain"
)

type Service interface {
	Create(customers *domain.Customers) error
	ReadAll() ([]*domain.Customers, error)
	LoadJson() (customers []*domain.Customers, err error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(customers *domain.Customers) error {
	_, err := s.r.Create(customers)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ReadAll() ([]*domain.Customers, error) {
	return s.r.ReadAll()
}

//Función para cargar el file y setear el id 
func (s *service) LoadJson() (customers []*domain.Customers, err error) {
	//Leer el archivo y parsearlo a customers
	var file []byte
	file, err = os.ReadFile("../datos/customers.json")
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(file), &customers)
	if err != nil {
		return
	}

	//Mejoras: Intentar aplicar esto 
	////guardar cada uno de los registros
	//var savedCustomers []domain.Customers
	//var notSavedCUstomers []domain.Customers

	for _, customer := range customers {
		var id int64
		id, err = s.r.Create(customer)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		customer.Id = int(id)
	}

	return

}
