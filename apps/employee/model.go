package employee

import (
	db "github.com/ij4l/foodCatalog/database/postgres/sqlc"
	"github.com/ij4l/foodCatalog/graph/model"
)

func New(e *db.Employee) *model.Employee {
	return &model.Employee{
		ID:        int(e.ID),
		Name:      e.Name.String,
		Nip:       e.Nip.String,
		Address:   e.Address.String,
		CreatedAt: e.CreatedAt.Time.String(),
		UpdatedAt: e.UpdatedAt.Time.String(),
	}
}

func NewEmployee(e *db.SelectEmployeeByIDRow) *model.Employee {
	return &model.Employee{
		ID:        int(e.ID),
		Name:      e.Name.String,
		Nip:       e.Nip.String,
		Address:   e.Address.String,
		CreatedAt: e.CreatedAt.Time.String(),
		UpdatedAt: e.UpdatedAt.Time.String(),
	}
}

func NewList(es []db.SelectAllEmployeeRow) (me []*model.Employee) {
	for _, e := range es {
		me = append(me, &model.Employee{
			ID:        int(e.ID),
			Name:      e.Name.String,
			Nip:       e.Nip.String,
			Address:   e.Address.String,
			CreatedAt: e.CreatedAt.Time.String(),
			UpdatedAt: e.UpdatedAt.Time.String(),
		})
	}

	return
}
