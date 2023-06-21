package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mahendrabagul/demo-app/invoice-service/pkg/rest/server/models"
	"github.com/mahendrabagul/demo-app/invoice-service/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type InvoiceController struct {
	invoiceService *services.InvoiceService
}

func NewInvoiceController() (*InvoiceController, error) {
	invoiceService, err := services.NewInvoiceService()
	if err != nil {
		return nil, err
	}
	return &InvoiceController{
		invoiceService: invoiceService,
	}, nil
}

func (invoiceController *InvoiceController) CreateInvoice(context *gin.Context) {
	// validate input
	var input models.Invoice
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger invoice creation
	if err := invoiceController.invoiceService.CreateInvoice(input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Invoice created successfully"})
}

func (invoiceController *InvoiceController) UpdateInvoice(context *gin.Context) {
	// validate input
	var input models.Invoice
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger invoice update
	if err := invoiceController.invoiceService.UpdateInvoice(id, input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Invoice updated successfully"})
}

func (invoiceController *InvoiceController) FetchInvoice(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger invoice fetching
	invoice, err := invoiceController.invoiceService.GetInvoice(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, invoice)
}

func (invoiceController *InvoiceController) DeleteInvoice(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger invoice deletion
	if err := invoiceController.invoiceService.DeleteInvoice(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Invoice deleted successfully",
	})
}

func (invoiceController *InvoiceController) ListInvoices(context *gin.Context) {
	// trigger all invoices fetching
	invoices, err := invoiceController.invoiceService.ListInvoices()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, invoices)
}

func (*InvoiceController) PatchInvoice(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*InvoiceController) OptionsInvoice(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*InvoiceController) HeadInvoice(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
