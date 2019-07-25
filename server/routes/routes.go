// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package routes

import (
	"github.com/valasek/quasar-starter-kit-go-gin/server/api"
	"github.com/valasek/quasar-starter-kit-go-gin/server/logger"

	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"github.com/spf13/viper"
)

var w *tabwriter.Writer

func noRoute(c *gin.Context) {
	path := strings.Split(c.Request.URL.Path, "/")
	if (path[1] != "") && (path[1] == "api") {
		c.JSON(http.StatusNotFound, gin.H{"msg": "no route", "body": nil})
	} else {
		c.HTML(http.StatusOK, "index.html", "")
	}
}

// Logger provides logrus logger middleware
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		// after request
		latency := time.Since(t)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		path := c.Request.URL.Path
		message := fmt.Sprintf("[server] | %3d | %12v |%s | %-7s %s %s",
			statusCode,
			latency,
			clientIP,
			method,
			path,
			c.Errors.String(),
		)
		switch {
		case statusCode >= 400 && statusCode <= 499:
			logger.Log.Warning(message)
		case statusCode >= 500:
			logger.Log.Error(message)
		default:
			logger.Log.Info(message)
		}
	}
}

// SetupRouter builds the routes for the api
func SetupRouter(api *api.API) *gin.Engine {

	gin.DefaultWriter = colorable.NewColorableStdout()
	if viper.GetString("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(Logger())

	// set CORS
	// router.Use(cors.Default())
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	// no route, bad url
	router.NoRoute(noRoute)

	router.Use(static.Serve("/", static.LocalFile("./client/dist", true)))

	a := router.Group("/api")
	{
		// download all data
		a.GET("/download/data", api.Download)
		// upload all data
		a.POST("/upload/data", api.Upload)
	}
	// handle 404 and due to Vue history mode return home page
	router.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(".", "client", "dist", "index.html"))
	})

	return router
}

// PrintRoutes prints all set routes
func PrintRoutes(c *gin.Engine) {
	fmt.Println("gin.DebugPrintRouteFunc()")
}
