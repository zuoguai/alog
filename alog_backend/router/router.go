package router

import (
	"alog/configs"
	"alog/internal/handlers"
	articleview "alog/internal/handlers/article"
	"alog/internal/pkg"
	"alog/internal/pkg/log"
	"alog/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func InitRouterAndServe() {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	// 设置路由
	r.Use(Cors())
	setRoutes(r)

	// 启动server
	port := configs.GetGlobalConfig().AppConfig.Port
	if err := r.Run(":" + strconv.Itoa(port)); err != nil {
		log.Errorf("start server err:" + err.Error())
	}
}

func setRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")
	apiGroup.POST("/login", handlers.Login)

	articleGroup := apiGroup.Group("/article", AddPrivilege())
	articleGroup.GET("/:article_id", articleview.View)
	articleGroup.GET("/list", articleview.List)

	adminGroup := apiGroup.Group("/admin", AuthAdmin())
	adminGroup.POST("", articleview.Create)
	adminGroup.PUT("", articleview.Modify)
	adminGroup.DELETE("/:article_id", articleview.Del)
}

func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := ""
		token = c.Request.Header.Get("Authorization")

		//token = c.Request.FormValue("token")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": -1, "msg": "请求未携带token，无权限访问"})
			c.Abort()
			return

		}
		claims, err := utils.ParseJwtToken(token, configs.GetGlobalConfig().JwtConfig.Secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": -1, "msg": err.Error()})
			c.Abort()
			return
		}
		//token超时
		if time.Now().Unix() > claims.StandardClaims.ExpiresAt {
			c.JSON(http.StatusUnauthorized, gin.H{"status": -1, "msg": "token过期"})
			c.Abort() //阻止执行
			return
		}
		c.Set(pkg.UserID, claims.UserID)
		c.Set(pkg.IsAdmin, true)
		c.Next()
	}
}
func AddPrivilege() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := ""
		token = c.Request.Header.Get("Authorization")
		if token == "" {
			c.Next()
			return
		}
		claims, err := utils.ParseJwtToken(token, configs.GetGlobalConfig().JwtConfig.Secret)
		if err != nil {
			c.Next()
			return
		}
		//token超时
		if time.Now().Unix() > claims.StandardClaims.ExpiresAt {
			c.Next()
			return
		}
		c.Set(pkg.UserID, claims.UserID)
		c.Set(pkg.IsAdmin, true)
		c.Next()
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin,Authorization, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length,Authorization"+
				"Access-Control-Allow-Origin, "+
				"Access-Control-Allow-Headers, "+
				"Cache-Control, "+
				",Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
