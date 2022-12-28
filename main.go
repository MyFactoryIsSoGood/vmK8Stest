package vmK8Stest

import (
	"github.com/gin-gonic/gin"
	"os"
	"vmK8Stest/controllers"
)

func main() {
	router := gin.Default()
	router.GET("/test", controllers.HelloWorlder)

	err := router.Run(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}
