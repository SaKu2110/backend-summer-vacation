package main

import(
	// import gin library
	"github.com/gin-gonic/gin"

	// import gorm library
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

	// import sample API packages
	"github.com/miraikeitai2020/backend-summer-vacation/pkg/debug"
	"github.com/miraikeitai2020/backend-summer-vacation/pkg/config"
	"github.com/miraikeitai2020/backend-summer-vacation/pkg/server/controller"
)

const(
	DB_DRIVER_MYSQL = "mysql"
)

func initializeDataBase() (*gorm.DB, error) {
	token := config.GetConnectionToken()
	for i:=0; i<5; i++ {
		if db, err := gorm.Open(DB_DRIVER_MYSQL, token);err == nil {
			return db, nil
		}
		config.BackOff(i)
	}
	return nil, debug.Err("Faild connection database")
}

func initializeController(db *gorm.DB) controller.Controller {
	return controller.Controller{
		DB:	db,
	}
}

func setupRooter(ctrl controller.Controller) *gin.Engine {
	router := gin.Default()
	// API Handlers List
	router.GET("/", ctrl.HelloWorld)
	router.POST("/sayhello", ctrl.SayHello)
	// 以下， 課題の進行状況に応じてコメントアウトを外す
	router.GET("/task1", ctrl.Task1)
	router.POST("/task2", ctrl.Task2)
	router.POST("/task3", ctrl.SignUp)
	router.POST("/task4", ctrl.SignIn)
	return router
}

func main() {
	db, err := initializeDataBase()
	if err != nil {
		panic(err)
	}
	ctrl := initializeController(db)
	router := setupRooter(ctrl)
	err = router.Run(":8080")
	if err != nil {
		panic("server Painc")
	}
}