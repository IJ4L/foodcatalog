package employee

import (
	"context"

	"github.com/ij4l/foodCatalog/apps"
	db "github.com/ij4l/foodCatalog/database/postgres/sqlc"
	typ "github.com/ij4l/foodCatalog/database/postgres/type"
	"github.com/ij4l/foodCatalog/graph/model"
)

type employeeRepository struct {
	repo apps.AppRepository
}

// findEmployeeById implements employeeRepositoryContract.
func (e employeeRepository) findEmployeeById(ctx context.Context, empId int) (me model.Employee, err error) {
	employee, err := e.repo.SelectEmployeeByID(ctx, int32(empId))
	if err != nil {
		return
	}

	me = *NewEmployee(&employee)
	return
}

// newEmployee implements employeeRepositoryContract.
func (e employeeRepository) newEmployee(ctx context.Context, ne model.NewEmployee) (me model.Employee, err error) {
	arg := db.InsertEmployeeParams{
		Name:    typ.Text(ne.Name),
		Address: typ.Text(ne.Address),
		Nip:     typ.Text(ne.Nip),
	}

	employee, err := e.repo.InsertEmployee(ctx, arg)
	if err != nil {
		return
	}

	me = *New(&employee)
	return
}

// findAllEmployees implements employeeRepositoryContract.
func (e employeeRepository) findAllEmployees(ctx context.Context) (me []*model.Employee, err error) {
	employees, err := e.repo.SelectAllEmployee(ctx)
	if err != nil {
		return
	}

	me = NewList(employees)
	return
}

// removeEmployeeById implements employeeRepositoryContract.
func (e employeeRepository) removeEmployeeById(ctx context.Context, empId int) (err error) {
	err = e.repo.RemoveEmployeeByID(ctx, int32(empId))
	return
}

func NewEmployeeRepository(repo *apps.AppRepository) employeeRepository {
	return employeeRepository{repo: *repo}
}
