package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	ConfigRuntime()
	StartGin()
}

// ConfigRuntime sets the number of operating system threads.
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func index(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "this is a test, but also a good test!")
}

// StartGin starts gin web server with setting router.
func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.LoadHTMLGlob("resources/*.templ.html")
	router.Static("/static", "resources/static")
	router.GET("/", index)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
