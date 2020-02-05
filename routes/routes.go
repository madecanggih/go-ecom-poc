package routes

import (
	"architect/saras-go-poc/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	api := e.Group("/api/v1")
	{
		api.POST("/auth/login", handlers.PostLogin)
		api.POST("/auth/register", handlers.PostRegister)

		// api.GET("/book/checkout/:user_id", handlers.GetCheckout)
		// api.POST("/book/checkout/:user_id", handlers.PostCheckout)

		// api.POST("/cart", handlers.PostCart)
		// api.GET("/cart/:id", handlers.GetCart)
		// api.PUT("/cart/:id", handlers.PutCart)
		// api.DELETE("/cart/:id", handlers.DeleteCart)

		// api.GET("/category", handlers.GetCategory)
		// api.POST("/category", handlers.PutCategory)
		// api.GET("/category/:id", handlers.GetCategory)

		// api.GET("/invoice/:id", handlers.GetInvoice)
		// api.GET("/invoice/history/:user_id", handlers.GetInvoiceHistory)

		// api.GET("/product", handlers.GetProduct)
		// api.POST("/product", handlers.PostProduct)
		// api.GET("/product/:id", handlers.GetProduct)

		// api.GET("/promo", handlers.GetPromo)
		// api.POST("/promo", handlers.PostPromo)
		// api.GET("/promo/:code", handlers.GetPromo)
		// api.PUT("/promo/:code", handlers.PutPromo)
		// api.DELETE("/promo/:code", handlers.DeletePromo)

		// api.GET("/store", handlers.GetStore)
		// api.POST("/store", handlers.PostStore)
		// api.GET("/store/:id", handlers.GetStore)
		// api.PUT("/store/:id", handlers.PutStore)
		// api.DELETE("/store/:id", handlers.DeleteStore)

		// api.GET("/trolley/:user_id", handlers.GetTrolley)

		api.GET("/users", handlers.GetUsers)
		api.GET("/users/:id", handlers.GetUsers)
		api.GET("/userlist", handlers.UserList)
		api.GET("/userlist/:id", handlers.UserList)

		// api.POST("/wishlist", handlers.PostWishlist)
		// api.GET("/wishlist/:id", handlers.GetWishlist)
		// api.DELETE("/wishlist/:id", handlers.DeleteWishlist)
	}

	return e
}
