package main

import (
	"github.com/samrat-rm/OrderService-GO.git/product"
	"github.com/samrat-rm/OrderService-GO.git/user"
)

func main() {
	user.InitialMigration()
	product.InitialMigrationProduct()
}
