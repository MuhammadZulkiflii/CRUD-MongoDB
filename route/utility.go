package route

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/MuhammadZulkiflii/CRUD-MongoDB/errors"

	"github.com/gin-gonic/gin"
)

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Auth-User, x-callback-token, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT,DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatusJSON(204, gin.H{"code": 204, "error": "cors origin", "data": nil})
			return
		}
		c.Next()
	}
}

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // Start timer
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Fill the params
		param := gin.LogFormatterParams{}

		param.TimeStamp = time.Now() // Stop timer
		param.Latency = param.TimeStamp.Sub(start)
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

		param.BodySize = c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path
		errMsg := strings.Join(c.Errors.Errors(), "\n")
		log.Printf(`{"ip": "%s","method": "%s", "status": "%d", "path": "%s", "latency": "%s", "err": "%s"}`, param.ClientIP, param.Method, param.StatusCode, param.Path, param.Latency.String(), errMsg)
	}
}
func NewRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "8080")
	}
	engine := gin.New()
	engine.Use(cors())
	engine.Use(logger())
	return engine
}

func Response(ctx *gin.Context, data interface{}, err error) {
	c, e := errors.FromError(err)
	if e != nil {
		ctx.AbortWithStatusJSON(c, gin.H{"code": c, "data": nil, "errors": e.Error()})
		return
	}
	ctx.JSON(c, gin.H{"code": c, "data": data, "errors": nil})
}
