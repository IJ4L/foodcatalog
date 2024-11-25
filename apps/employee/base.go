package employee

import (
	"github.com/gin-gonic/gin"
	"github.com/ij4l/foodCatalog/apps"
)

func InitializeAuthHandler(repo *apps.AppRepository) (em EmployeeHandler) {
	ctx := &gin.Context{}
	er := NewEmployeeRepository(repo)
	as := NewEmployeeService(er, ctx)
	em = NewEmployeeHandler(as)
	return em
}