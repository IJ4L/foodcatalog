package employee

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ij4l/foodCatalog/graph/model"
)

type employeeRepositoryContract interface {
	newEmployee(ctx context.Context, ne model.NewEmployee) (me model.Employee, err error)
	findAllEmployees(ctx context.Context) (me []*model.Employee, err error)
	findEmployeeById(ctx context.Context, empId int) (me model.Employee, err error)
	removeEmployeeById(ctx context.Context, empId int) (err error)
}

type employeeService struct {
	ctx *gin.Context
	er  employeeRepositoryContract
}

func NewEmployeeService(es employeeRepositoryContract, ctx *gin.Context) employeeService {
	return employeeService{er: es, ctx: ctx}
}

func (es employeeService) createEmployee(ne model.NewEmployee) (me model.Employee, err error) {
	me, err = es.er.newEmployee(es.ctx, ne)
	return
}

func (es employeeService) listEmployee() (me []*model.Employee, err error) {
	me, err = es.er.findAllEmployees(es.ctx)
	return
}

func (es employeeService) removeEmployeeByID(empID int) (err error) {
	err = es.er.removeEmployeeById(es.ctx, empID)
	return
}

func (es employeeService) employee(empID int) (me model.Employee, err error) {
	me, err = es.er.findEmployeeById(es.ctx, empID)
	return
}
