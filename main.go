package main

import (
	"github.com/Jac-Bo/gofigure/utils/db"
	"github.com/Jac-Bo/gofigure/utils/engine"
)

func main() {
	db.DatabaseConnection()
	engine.StartGin()
}
