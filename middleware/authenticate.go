package middleware

import (
	"FoodDelivery/common"
	"FoodDelivery/components/appcontext"
	"FoodDelivery/components/tokenprovider/jwt"
	userstore "FoodDelivery/module/user/store"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"wrong authentication header",
		"ErrWrongAuthHeader")
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")

	if parts[0] != "Bearer" || len(parts) != 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(errors.New("invalid permission"))
	}

	return parts[1], nil
}

func RequireAuth(appCtx appcontext.AppContext) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()
		store := userstore.NewSQLStore(db)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("client has been delete or banned")))
		}

		user.Mask(false)
		c.Set(common.CurrentUser, user)
		c.Next()
	}
}