package initialize

import (
	"gin-server-cli/global"
	"gin-server-cli/middleware"
	"gin-server-cli/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var r = gin.Default()
	//Router.StaticFS(global.Config.Local.Path, http.Dir(global.Config.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了

	//使用zap作为gin的日志中间件，但是会出现一条请求控制台打印两条日志，解决办法是关闭zap的控制台打印或者使用gin.New()
	//r.Use(ginzap.Ginzap(global.ZapLog, time.RFC3339, true))
	//r.Use(ginzap.RecoveryWithZap(global.ZapLog, true))

	// 跨域
	r.Use(middleware.Cors())
	global.ZapLog.Info("use middleware cors")

	//swagger处理
	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//global.ZapLog.Info("register swagger handler")

	v1 := r.Group("/api/v1")
	PublicGroup := v1.Group("")
	{
		router.InitPublicRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}
	//privateGroup := v1.Group("")
	//PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//{
	//router.InitApiRouter(PrivateGroup)                   // 注册功能api路由
	//router.InitJwtRouter(PrivateGroup)                   // jwt相关路由
	//router.InitUserRouter(PrivateGroup)                  // 注册用户路由
	//router.InitMenuRouter(PrivateGroup)                  // 注册menu路由
	//router.InitEmailRouter(PrivateGroup)                 // 邮件相关路由
	//router.InitSystemRouter(PrivateGroup)                // system相关路由
	//router.InitCasbinRouter(PrivateGroup)                // 权限相关路由
	//router.InitCustomerRouter(PrivateGroup)              // 客户路由
	//router.InitAutoCodeRouter(PrivateGroup)              // 创建自动化代码
	//router.InitAuthorityRouter(PrivateGroup)             // 注册角色路由
	//router.InitSimpleUploaderRouter(PrivateGroup)        // 断点续传（插件版）
	//router.InitSysDictionaryRouter(PrivateGroup)         // 字典管理
	//router.InitSysOperationRecordRouter(PrivateGroup)    // 操作记录
	//router.InitSysDictionaryDetailRouter(PrivateGroup)   // 字典详情管理
	//router.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由
	//router.InitWorkflowProcessRouter(PrivateGroup)       // 工作流相关接口
	//}
	global.ZapLog.Info("router register success")
	return r
}
