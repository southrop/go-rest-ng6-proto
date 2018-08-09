package main

import (
    "github.com/gin-gonic/gin"
    "github.com/delsner/go-rest-ng6-proto/backend/utils"
    "github.com/delsner/go-rest-ng6-proto/backend/controllers"
)

func setupRouter() *gin.Engine {
    r := gin.Default() // init router with default mw (e.g. logging)

    user := r.Group("/api/user")
    {
        user.POST("/", controllers.CreateUser)
        user.GET("/:userId", controllers.GetUserById)
        user.GET("/", controllers.GetAllUsers)
    }

    return r
}

func main() {
    // setup router
    r := setupRouter()

    // get port from env or use default 8080
    port := utils.GetEnv("PORT", "8080")

    // run http server
    r.Run(":" + port)
}
