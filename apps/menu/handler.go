package menu

import "github.com/ij4l/foodCatalog/graph/model"

type MenuHandler struct {
	ms menuService
}

func NewMenuHandler(ms menuService) MenuHandler {
	return MenuHandler{ms: ms}
}

func (mh MenuHandler) InsertMenu(mmi *model.NewMenu) (mm *model.Menu, err error) {
	mm, err = mh.ms.InsertMenu(mmi)
	return
}

func (mh MenuHandler) FindAll() (mm []*model.Menu, err error) {
	mm, err = mh.ms.FindAll()
	return
}

func (mh MenuHandler) FindById(id int) (mm *model.Menu, err error) {
	mm, err = mh.ms.FindById(id)
	return
}
