package daos

import (
	"errors"
	"github.com/mahendrabagul/demo-app/invoice-service/pkg/rest/server/daos/clients/sqls"
	invoiceClient "github.com/mahendrabagul/demo-app/invoice-service/pkg/rest/server/daos/clients/sqls/invoice-client"
	"github.com/mahendrabagul/demo-app/invoice-service/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
)

type InvoiceDao struct {
	sqlClient *sqls.SQLiteClient
}

func NewInvoiceDao() (*InvoiceDao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = invoiceClient.Migrate(sqlClient)
	if err != nil {
		return nil, err
	}
	return &InvoiceDao{
		sqlClient,
	}, nil
}

func (invoiceDao *InvoiceDao) CreateInvoice(invoice models.Invoice) error {
	_, err := invoiceClient.Create(invoiceDao.sqlClient, invoice)
	if err != nil {
		return err
	}
	log.Debugf("invoice created")
	return nil
}

func (invoiceDao *InvoiceDao) UpdateInvoice(id int64, invoice models.Invoice) error {
	if id != invoice.Id {
		return errors.New("id and payload don't match")
	}
	_, err := invoiceClient.Update(invoiceDao.sqlClient, id, invoice)
	if err != nil {
		return err
	}
	log.Debugf("invoice updated")
	return nil
}

func (invoiceDao *InvoiceDao) DeleteInvoice(id int64) error {
	err := invoiceClient.Delete(invoiceDao.sqlClient, id)
	if err != nil {
		return err
	}
	log.Debugf("invoice deleted")
	return nil
}

func (invoiceDao *InvoiceDao) ListInvoices() ([]models.Invoice, error) {
	invoices, err := invoiceClient.All(invoiceDao.sqlClient)
	if err != nil {
		return invoices, err
	}
	log.Debugf("invoice listed")
	return invoices, nil
}

func (invoiceDao *InvoiceDao) GetInvoice(id int64) (models.Invoice, error) {
	invoice, err := invoiceClient.Get(invoiceDao.sqlClient, id)
	if err != nil {
		return models.Invoice{}, err
	}
	log.Debugf("invoice retrieved")
	return *invoice, nil
}
