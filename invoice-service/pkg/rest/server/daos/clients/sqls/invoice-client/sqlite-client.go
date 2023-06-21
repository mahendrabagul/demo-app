package invoice_client

import (
	"database/sql"
	"errors"
	"github.com/mahendrabagul/demo-app/invoice-service/pkg/rest/server/daos/clients/sqls"
	"github.com/mahendrabagul/demo-app/invoice-service/pkg/rest/server/models"
)

func Migrate(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS invoices(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
        Amount TEXT NOT NULL,
        
        Items TEXT NOT NULL,
        
        Name TEXT NOT NULL,
        
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err := r.DB.Exec(query)
	return err
}

func Create(r *sqls.SQLiteClient, m models.Invoice) (*models.Invoice, error) {
	insertQuery := "INSERT INTO invoices(Amount, Items, Name)values(?, ?, ?)"
	res, err := r.DB.Exec(insertQuery, m.Amount, m.Items, m.Name)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	return &m, nil
}

func All(r *sqls.SQLiteClient) ([]models.Invoice, error) {
	selectQuery := "SELECT * FROM invoices"
	rows, err := r.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var all []models.Invoice
	for rows.Next() {
		var m models.Invoice
		if err := rows.Scan(&m.Id, &m.Amount, &m.Items, &m.Name); err != nil {
			return nil, err
		}
		all = append(all, m)
	}
	if all == nil {
		all = []models.Invoice{}
	}
	return all, nil
}

func Get(r *sqls.SQLiteClient, id int64) (*models.Invoice, error) {
	selectQuery := "SELECT * FROM invoices WHERE Id = ?"
	row := r.DB.QueryRow(selectQuery, id)

	var m models.Invoice
	if err := row.Scan(&m.Id, &m.Amount, &m.Items, &m.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}
	return &m, nil
}

func Update(r *sqls.SQLiteClient, id int64, m models.Invoice) (*models.Invoice, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	updateQuery := "UPDATE invoices SET Amount = ?, Items = ?, Name = ? WHERE Id = ?"
	res, err := r.DB.Exec(updateQuery, m.Amount, m.Items, m.Name, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	return &m, nil
}

func Delete(r *sqls.SQLiteClient, id int64) error {
	deleteQuery := "DELETE FROM invoices WHERE Id = ?"
	res, err := r.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	return err
}
