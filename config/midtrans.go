package config

import (
	"log"
	"os"

	"github.com/veritrans/go-midtrans"
)

var MidtransClient midtrans.Client

// InitMidtrans - Inisialisasi client Midtrans
func InitMidtrans() {
	// Ambil Server Key dari environment
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	if serverKey == "" {
		log.Fatal("MIDTRANS_SERVER_KEY is not set in the environment variables")
	}

	// Inisialisasi Midtrans client
	MidtransClient = midtrans.NewClient()
	MidtransClient.ServerKey = serverKey
	MidtransClient.ClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
	MidtransClient.APIEnvType = midtrans.Sandbox // Ubah ke midtrans.Production untuk produksi

	log.Println("Midtrans client initialized successfully")
}
