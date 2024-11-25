package menu

import (
	"context"

	"github.com/ij4l/foodCatalog/apps"
	db "github.com/ij4l/foodCatalog/database/postgres/sqlc"
	typ "github.com/ij4l/foodCatalog/database/postgres/type"
	"github.com/ij4l/foodCatalog/graph/model"
)

type menuRepository struct {
	repo apps.AppRepository
}

// findAll implements menuRepositoryContract.
func (m menuRepository) findAll(ctx context.Context) (mm []*model.Menu, err error) {
	menus, err := m.repo.SelectAllMenu(ctx)
	if err != nil {
		return
	}

	mm = NewList(&menus)
	return
}

// findById implements menuRepositoryContract.
func (m menuRepository) findById(ctx context.Context, id int) (mm *model.Menu, err error) {
	menu, err := m.repo.SelectMenuByID(ctx, int32(id))
	if err != nil {
		return
	}

	mm = NewRow(&menu)
	return
}

// insertMenu implements menuRepositoryContract.
func (m menuRepository) insertMenu(ctx context.Context, mmi *model.NewMenu) (mm *model.Menu, err error) {
	arg := db.InsertMenuParams{
		Name:        typ.Text(mmi.Name),
		Category:    typ.Text(mmi.Category),
		Description: typ.Text(mmi.Desc),
		Price:       typ.Int4(mmi.Price),
	}

	menu, err := m.repo.InsertMenu(ctx, arg)
	mm = NewMenu(&menu)
	return
}

func NewMenuRepository(repo *apps.AppRepository) menuRepository {
	return menuRepository{repo: *repo}
}
