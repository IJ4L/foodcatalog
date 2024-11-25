package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/ij4l/foodCatalog/apps"
)

func InitializeMenuHandler(repo *apps.AppRepository) (mh MenuHandler) {
	ctx := &gin.Context{}
	mr := NewMenuRepository(repo)
	ms := NewMenuService(mr, ctx)
	mh = NewMenuHandler(ms)
	return mh
}
