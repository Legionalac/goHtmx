package api

import (
	customMiddlewares "goHtmx/api/middlewares"
	libs "goHtmx/internal/core"
	authservice "goHtmx/internal/services/authService"
	views "goHtmx/web/templates"

	"net/http"

	"github.com/a-h/templ"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitHttpServer() *echo.Echo {
	logger := libs.GetLogger()
	//-------------DB AND SERVICES---------
	db := libs.CockroachDb{}
	db.CreateCockRoachConnection()
	defer db.CloseConnection()

	authserviceInstance := authservice.AuthService{Db: db}

	e := echo.New()
	//-------------MIDDLEWARES-------------
	customMiddlewares.LoggerMiddleware(e, logger)
	//---------------ROUTES----------------
	e.GET("/", func(c echo.Context) error {
		return Render(c, http.StatusOK, views.Index(views.Form()))
	})
	e.GET("/register", func(c echo.Context) error {
		return Render(c, http.StatusOK, views.Index(views.Register(false)))
	})
	e.POST("/login", func(c echo.Context) error {
		email := c.FormValue("email")
		pass := c.FormValue("password")
		if email == "" || pass == "" {
			return echo.NewHTTPError(http.StatusInternalServerError,"BAD PARAMS")
		}

		cookie := libs.GenerateCookie(email)
		c.SetCookie(cookie)
		return c.Redirect(http.StatusSeeOther,"/dashboard")
	})
	e.POST("/createAccount", func(c echo.Context) error {
		email := c.FormValue("email")
		pass := c.FormValue("password")
		err := authserviceInstance.CreateAccount(email,pass)
		if err != nil {
			logger.Error().Msg("server.go::CreateAccount " + err.Error())
			return Render(c, http.StatusOK, views.Index(views.Register(true)))
		}
		return Render(c, http.StatusOK, views.Index(views.Form()))
	})
	e.Static("/static/*", "web/static")
	users := e.Group("/")
	{
		config := echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(libs.JwtCustomClaims)
			},
			SigningKey: []byte("secret"),
			TokenLookup: "cookie:token",
		}
		users.Use(echojwt.WithConfig(config))
		users.GET("dashboard", func(c echo.Context) error {
			return Render(c, http.StatusOK, views.Index(views.Content()))
		})
	}

	//-------------------------------------
	logger.Info().Msg("Starting server on port 3000")
	e.Start(":3000")

	return e
}


func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	ctx.Response().Header().Set(echo.HeaderAccessControlAllowCredentials, "true")
	ctx.Response().Header().Set(echo.HeaderAccessControlAllowMethods, "true")
	ctx.Response().Header().Set(echo.HeaderAccessControlAllowHeaders , "true")
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}
