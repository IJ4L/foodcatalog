package menu

import (
	db "github.com/ij4l/foodCatalog/database/postgres/sqlc"
	"github.com/ij4l/foodCatalog/graph/model"
)

func NewRow(menu *db.SelectMenuByIDRow) *model.Menu {
	return &model.Menu{
		ID:        int(menu.ID),
		Name:      menu.Name.String,
		Category:  menu.Category.String,
		Desc:      menu.Description.String,
		Price:     int(menu.Price.Int32),
		CreatedAt: menu.CreatedAt.Time.String(),
		UpdatedAt: menu.UpdatedAt.Time.String(),
	}
}

func NewMenu(menu *db.Menu) *model.Menu {
	return &model.Menu{
		ID:        int(menu.ID),
		Name:      menu.Name.String,
		Category:  menu.Category.String,
		Desc:      menu.Description.String,
		Price:     int(menu.Price.Int32),
		CreatedAt: menu.CreatedAt.Time.String(),
		UpdatedAt: menu.UpdatedAt.Time.String(),
	}
}

func NewList(menus *[]db.SelectAllMenuRow) (mm []*model.Menu) {
	for _, menu := range *menus {
		mm = append(mm, &model.Menu{
			ID:        int(menu.ID),
			Name:      menu.Name.String,
			Category:  menu.Category.String,
			Desc:      menu.Description.String,
			Price:     int(menu.Price.Int32),
			CreatedAt: menu.CreatedAt.Time.String(),
			UpdatedAt: menu.UpdatedAt.Time.String(),
		})
	}

	return
}
