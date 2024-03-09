package router

import (
	"blog/controller"
	"blog/logger"
	"blog/middlewares"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	//_ "blog/docs" // 千万不要忘了导入把你上一步生成的docs

	"github.com/gin-contrib/pprof"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1))

	//r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	//注册业务路由

	v1 := r.Group("/api/v1")

	//不需要JWT认证的页面
	{
		//注册
		v1.POST("/signup", controller.SignUpHandler)
		//登录
		v1.POST("/login", controller.LoginHandler)
		//依据时间或分数获取帖子列表
		v1.GET("/posts", controller.GetPostListHandler)

	}

	v1.Use(middlewares.JWTAuthMiddleware()) //应用JWT认证中间件
	{
		//创建帖子
		v1.POST("/post", controller.CreatePostHandler)
		//删除帖子
		v1.POST("/deleteP/:id", controller.DeletePostHandler)
		//投票
		v1.POST("/vote", controller.PostVoteHandler)

	}

	//{
	//	v1.GET("/community", controller.CommunityHandler)
	//	v1.GET("/community/:id", controller.CommunityDetailHandler)
	//
	//	v1.POST("/post", controller.CreatePostHandler)
	//	v1.GET("/post/:id", controller.GetPostDetailHandler)
	//	//不分时间或者分数获取帖子列表
	//	v1.GET("/posts", controller.GetPostListHandler)
	//	//根据时间或分数获取帖子列表
	//	v1.GET("/posts2", controller.GetPostListHandler2)
	//	////根据社区获取帖子列表(默认时间，可以改为分数)
	//	//v1.GET("/posts3", controller.GetCommunityPostListHandler)
	//
	//	//投票
	//	v1.POST("/vote", controller.PostVoteHandler)
	//}

	pprof.Register(r) //注册pprof相关路由

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
