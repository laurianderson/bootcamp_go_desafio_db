package invoices

import (
	"database/sql"
	"github.com/laurianderson/bootcamp_go_desafio_db/internal/domain"
)

const (
	QueryGetTotal = `
		SELECT i.id, SUM(p.price * s.quantity) AS total
		FROM invoices AS i
		INNER JOIN sales AS s ON i.id = s.invoice_id
		INNER JOIN products AS p ON s.product_id = p.id
		GROUP BY i.id;`
	QueryUpdateTotal = `UPDATE invoices SET total = ? WHERE id = ?;`
)

type Repository interface {
	Create(invoices *domain.Invoices) (int64, error)
	ReadAll() ([]*domain.Invoices, error)
	Update() error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(invoices *domain.Invoices) (int64, error) {
	query := `INSERT INTO invoices (customer_id, datetime, total) VALUES (?, ?, ?)`
	row, err := r.db.Exec(query, &invoices.CustomerId, &invoices.Datetime, &invoices.Total)
	if err != nil {
		return 0, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) ReadAll() ([]*domain.Invoices, error) {
	query := `SELECT id, customer_id, datetime, total FROM invoices`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	invoices := make([]*domain.Invoices, 0)
	for rows.Next() {
		invoice := domain.Invoices{}
		err := rows.Scan(&invoice.Id, &invoice.CustomerId, &invoice.Datetime, &invoice.Total)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, &invoice)
	}
	return invoices, nil
}

func (r *repository) Update() error {
	// prepare statement
	rows, err := r.db.Query(QueryGetTotal)
	if err != nil {
		return err
	}
	defer rows.Close()

	// iterate over rows and update invoices
	for rows.Next() {
		var id int
		var total float64
		err = rows.Scan(&id, &total)
		if err != nil {
			return err
		}
		_, err = r.db.Exec(QueryUpdateTotal, total, id)
		if err != nil {
			return err
		}
	}
	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}