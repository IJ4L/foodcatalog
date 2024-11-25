package employee

import "github.com/ij4l/foodCatalog/graph/model"

type EmployeeHandler struct {
	es employeeService
}

func NewEmployeeHandler(es employeeService) EmployeeHandler {
	return EmployeeHandler{es: es}
}

func (eh EmployeeHandler) CreateEmployee(ne model.NewEmployee) (me model.Employee, err error) {
	me, err = eh.es.createEmployee(ne)
	return
}

func (eh EmployeeHandler) ListEmployee() (me []*model.Employee, err error) {
	me, err = eh.es.listEmployee()
	return
}

func (eh EmployeeHandler) RemoveEmployee(empID int) (err error) {
	err = eh.es.removeEmployeeByID(empID)
	return
}

func (eh EmployeeHandler) GetEmployeeByID(empID int) (me model.Employee, err error) {
	me, err = eh.es.employee(empID)
	return
}
