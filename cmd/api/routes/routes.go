package routes

import (
	"fmt"
	"shadow/cmd/api/routes/v1"
	"shadow/cmd/api/routes/v2"
	"strconv"
	"time"

	"shadow/internal/platform/mongo"

	"github.com/gin-gonic/gin"

	"shadow/internal/config"
)

type Request struct {
	ClientIP     string
	TimeStamp    string
	Method       string
	Path         string
	Protocol     string
	StatusCode   string
	Latency      string
	UserAgent    string
	ErrorMessage string
}

func Init() {
	config := config.GetConfig()
	col := mongo.Session.C("Logs")

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		col.Insert(Request{
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			strconv.Itoa(param.StatusCode),
			param.Latency.String(),
			param.Request.UserAgent(),
			param.ErrorMessage,
		})
		statusColor := param.StatusCodeColor()
		methodColor := param.MethodColor()
		resetColor := param.ResetColor()
		return fmt.Sprintf("[GIN] %v |%s %3d %s| %13v | %15s |%s %-7s %s %s\n%s",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			statusColor, param.StatusCode, resetColor,
			param.Latency,
			param.ClientIP,
			methodColor, param.Method, resetColor,
			param.Path,
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	router.LoadHTMLGlob("cmd/api/views/*.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	v1.Init(router.Group("v1"))
	v2.Init(router.Group("v2"))

	router.Run(":" + config.GetString("server.port"))
}
