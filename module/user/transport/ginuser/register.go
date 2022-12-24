package ginuser

import (
	"FoodDelivery/common"
	"FoodDelivery/components/appcontext"
	"FoodDelivery/components/hasher"
	userbiz "FoodDelivery/module/user/biz"
	usermodel "FoodDelivery/module/user/model"
	userstore "FoodDelivery/module/user/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx appcontext.AppContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := ctx.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(store, md5)

		if err := biz.Register(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.UID.String()))

	}
}