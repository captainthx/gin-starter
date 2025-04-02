package main

import (
	"fmt"
	"gin-starter/config"
	"gin-starter/core/domain"
	"gin-starter/docs"
	"gin-starter/logs"
	"gin-starter/midleware"
	"gin-starter/routes"
	"net/http"
	"time"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// @description Gin APIs
// @version 1.0
// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization
func main() {

	// กำหนดค่า timezone
	initTimezone()

	// รับค่า config  จาก env
	err := godotenv.Load(".env")
	if err != nil {
		logs.Error(err)
	}

	// เริ่มต้นกำหนดค่าต่างๆ
	config.Init()

	// ตั้งค่า redis client
	config.InitRedisClient(config.RedisBaseUrl)

	// สร้าง dsn (domain name server) ของ msql เพื่อใช้ กับ gorm
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DbUsername, config.DbPassword, config.DbHost, config.DbPort, config.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logs.Error(err)
	}

	// migrate database
	err = db.Set("gorm.table_options", "ENGINE=InnoDB").AutoMigrate(
		&domain.User{},
		&domain.File{},
	)
	if err != nil {
		logs.Error(err)
		return
	}

	// เริ่มต้นกำหนดค่า router ต่างๆ
	router := initRoute(db)
	err = router.Run(":" + config.ServerPort)
	if err != nil {
		logs.Error(err)
	}

}

// กำหนด route และตั้งค่า route ต่างๆ
func initRoute(db *gorm.DB) *gin.Engine {

	gin.SetMode(config.Mode)

	router := gin.Default()
	router.MaxMultipartMemory = 2 << 20 // 2 MB
	docs.SwaggerInfo.BasePath = "/"

	// สร้าง router สำหรับ swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// สรา้ง router สำหรับ scalar api document
	router.GET("/docs", func(c *gin.Context) {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			// SpecURL: "https://generator3.swagger.io/openapi.json",// เรียกใช้ข้อมูลจาก swagger doc
			SpecURL: config.FullBaseUrl + "/swagger/doc.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Gin APIs",
			},
			Layout: "modern",
		})
		if err != nil {
			fmt.Printf("%v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// routesต่างๆ

	// auth routes
	routes.RegisAuthRoutes(router, db)

	// file routes
	router.Use(midleware.RequireAuth)
	routes.RegisFileRoutes(router, db)

	// user routes

	router.Use(midleware.RequireAuth)
	routes.RegisUserRoutes(router, db)

	return router
}

func initTimezone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}
