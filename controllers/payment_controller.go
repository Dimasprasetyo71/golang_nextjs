package controllers

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"

	// "github.com/veritrans/go-midtrans"

	"DimasPrasetyo/backend-next/config"
	"DimasPrasetyo/backend-next/models"
)

// CreatePayment - Membuat transaksi pembayaran baru

// GetPaymentByID - Mendapatkan pembayaran berdasarkan ID
func GetPaymentByID(c *gin.Context) {
	id := c.Param("id")

	var payment models.Payment
	if err := config.ConnectDB().First(&payment, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	c.JSON(http.StatusOK, payment)
}

// GetAllPayments - Mendapatkan semua pembayaran
func GetAllPayments(c *gin.Context) {
	var payments []models.Payment
	if err := config.ConnectDB().Find(&payments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch payments"})
		return
	}

	c.JSON(http.StatusOK, payments)
}
// CreatePayment - Membuat transaksi pembayaran baru
func CreatePayment(c *gin.Context) {
	var request struct {
		UserID        int     `json:"user_id" binding:"required"`
		Amount        float64 `json:"amount" binding:"required"`
		PaymentMethod string  `json:"payment_method" binding:"required"`
		ItemID        string  `json:"item_id" binding:"required"`
		ItemName      string  `json:"item_name" binding:"required"`
	}

	// Bind JSON input
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input (optional additional logic)

	// Create a new payment record



	// Midtrans Transaction
	var snapClient = snap.Client{}
	snapClient.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	// Generate transaction details
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "MID-User-" + strconv.Itoa(request.UserID) + "-" + request.ItemID,
			GrossAmt: int64(request.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: "John",
			LName: "Doe",
			Email: "john@doe.com",
			Phone: "081234567890",
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				ID:    "Property-" + request.ItemID,
				Qty:   1,
				Price: int64(request.Amount),
				Name:  request.ItemName,
			},
		},
	}

	// Create Midtrans transaction
	response, errSnap := snapClient.CreateTransaction(req)
	if errSnap != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Midtrans transaction failed", "details": errSnap.GetRawError()})
		return
	}

	// Respond with payment and Midtrans data
	c.JSON(http.StatusOK, gin.H{
		"midtrans": gin.H{
			"token":        response.Token,
			"redirect_url": response.RedirectURL,
		},
	})
}


// UpdatePaymentStatus - Memperbarui status pembayaran
func UpdatePaymentStatus(c *gin.Context) {
	id := c.Param("id")

	var payment models.Payment
	if err := config.ConnectDB().First(&payment, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	var input struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment.Status = input.Status
	if err := config.ConnectDB().Save(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update payment"})
		return
	}

	c.JSON(http.StatusOK, payment)
}

// DeletePayment - Menghapus pembayaran
func DeletePayment(c *gin.Context) {
	id := c.Param("id")

	if err := config.ConnectDB().Delete(&models.Payment{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete payment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment deleted successfully"})
}
