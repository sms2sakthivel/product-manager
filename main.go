package main

import (
	_ "github.com/sms2sakthivel/product-manager/docs"
	"github.com/sms2sakthivel/product-manager/products"
)

func main() {
	// Step 1: Create a New User Service Application
	app := products.NewApp()

	// Step 2: Start Server and Listen on the Port 8001
	app.Listen(":8002")
}
