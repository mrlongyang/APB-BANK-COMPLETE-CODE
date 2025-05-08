
package main

import (
    "github.com/labstack/echo/v4"
    echojwt "github.com/labstack/echo-jwt/v4"
    "github.com/mrlongyang/apbbank/config"
    "github.com/mrlongyang/apbbank/handlers"
    "github.com/mrlongyang/apbbank/utils"
)

func main() {
    db := config.InitDB()
    e := echo.New()

    e.Static("/", "public") // Serve HTML files from public/

    userHandler := &handlers.UserHandler{DB: db}
    productHandler := &handlers.ProductHandler{DB: db}

    e.POST("/register", userHandler.Register)
    e.POST("/login", userHandler.Login)

    r := e.Group("/products")
    r.Use(echojwt.WithConfig(echojwt.Config{
        SigningKey: utils.SecretKey,
    }))
    r.POST("", productHandler.Create)
    r.GET("", productHandler.Read)
    r.PUT("", productHandler.Update)
    r.DELETE("/:id", productHandler.Delete)

    e.Logger.Fatal(e.Start(":8080"))
}
