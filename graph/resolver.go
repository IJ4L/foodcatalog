package graph

import (
	"github.com/ij4l/foodCatalog/apps/auth"
	"github.com/ij4l/foodCatalog/apps/employee"
	"github.com/ij4l/foodCatalog/apps/menu"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AuthHandler auth.AuthHandler
	EmployeeHandler employee.EmployeeHandler
	MenuHandler menu.MenuHandler
}
