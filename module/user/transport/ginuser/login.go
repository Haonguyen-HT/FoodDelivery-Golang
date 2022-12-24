package ginuser

import (
	"FoodDelivery/common"
	"FoodDelivery/components/appcontext"
	"FoodDelivery/components/hasher"
	"FoodDelivery/components/tokenprovider/jwt"
	userbiz "FoodDelivery/module/user/biz"
	usermodel "FoodDelivery/module/user/model"
	userstore "FoodDelivery/module/user/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		business := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*7)
		acc, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(acc))
	}
}