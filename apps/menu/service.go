package menu

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ij4l/foodCatalog/graph/model"
)

type menuRepositoryContract interface {
	insertMenu(ctx context.Context, mmi *model.NewMenu) (mm *model.Menu, err error)
	findAll(ctx context.Context) (mm []*model.Menu, err error)
	findById(ctx context.Context, id int) (mm *model.Menu, err error)
}

type menuService struct {
	ctx *gin.Context
	mr  menuRepositoryContract
}

func NewMenuService(mr menuRepositoryContract, ctx *gin.Context) menuService {
	return menuService{ctx: ctx, mr: mr}
}

func (ms menuService) InsertMenu(mmi *model.NewMenu) (mm *model.Menu, err error) {
	mm, err = ms.mr.insertMenu(ms.ctx, mmi)
	return
}

func (ms menuService) FindAll() (mm []*model.Menu, err error) {
	mm, err = ms.mr.findAll(ms.ctx)
	return
}

func (ms menuService) FindById(id int) (mm *model.Menu, err error) {
	mm, err = ms.mr.findById(ms.ctx, id)
	return
}
