package middlewares

import (
	"strconv"
	"strings"

	"github.com/alidevjimmy/db-project-go/internal/transport/http/response"
	"github.com/alidevjimmy/db-project-go/pkg/jwt"
	"github.com/labstack/echo/v4"
)

type AccountMiddleware struct {
	JwtPkg jwt.Jwt
}

func NewAccountMiddleware(jwtPkg jwt.Jwt) AccountMiddleware {
	return AccountMiddleware{
		JwtPkg: jwtPkg,
	}
}

func (am *AccountMiddleware) OnlyAccountOwner(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get(echo.HeaderAuthorization)
		errResp := response.ErrorResp{
			Error: response.Error{
				Code:    echo.ErrUnauthorized.Code,
				Message: "Unauthenticated",
			},
		}
		accID, err := strconv.Atoi(c.Param("account_id"))
		if accID == 0 || err != nil {
			return c.JSON(echo.ErrUnauthorized.Code, errResp)
		}

		if auth == "" {
			return c.JSON(errResp.Error.Code, errResp)
		}
		token := strings.Split(auth, " ")
		if len(token) != 2 {
			return c.JSON(errResp.Error.Code, errResp)
		}

		if token[0] != "Bearer" {
			return c.JSON(errResp.Error.Code, errResp)
		}

		claims, err := am.JwtPkg.ParseToken(token[1])
		if err != nil {
			return c.JSON(errResp.Error.Code, errResp)
		}
		
		if claims["sub"] != float64(accID) {
			return c.JSON(errResp.Error.Code, errResp)
		}
		return next(c)
	}
}
