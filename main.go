package main

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"task_cart/config"
	"task_cart/internal/model/entity"
	"task_cart/internal/repository"
	"task_cart/internal/service"
	v1 "task_cart/internal/transport/rest/controllers/v1"
	"task_cart/internal/transport/server"
	"task_cart/pkg/db"
)

type Test struct {
	Name  string `form:"name" json:"name" binding:"required,min=3"`
	Age   int    `form:"age" json:"age" binding:"required,min=0,max=110"`
	Email string `form:"email" json:"email" binding:"required,email"`
}

func main() {
	/*r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		var body Test

		if err := c.ShouldBindJSON(&body); err != nil {
			fmt.Println(err.(validator.ValidationErrors)[0].Field())
			fmt.Println(err.(validator.ValidationErrors)[0].Tag())
			fmt.Println()
			c.JSON(400, gin.H{
				"errors": parseError(err.(validator.ValidationErrors)),
			})
			return
		}

		c.JSON(200, gin.H{
			"msg": "Ok",
		})

	})
	r.Run()*/
	cfg := config.MustLoad()
	dbConn := db.MustStartDB(&cfg.DbConf, nil)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	defer db.MustCloseDB(dbConn, nil)
	err := dbConn.SetupJoinTable(&entity.Cart{}, "Products", &entity.CartProduct{})
	if err != nil {
		panic(err)
	}
	if err := dbConn.AutoMigrate(&entity.Product{}, &entity.Status{}, &entity.Cart{}, &entity.Order{}); err != nil {
		panic(err)
	}

	fmt.Println("Success migrate")

	//seedData(dbConn)
	//seedStatus(dbConn)
	logger.Info("Start service...")
	productRepo := repository.NewProductRepository(dbConn)
	statusRepo := repository.NewStatusRepository(dbConn)
	cartRepo := repository.NewCartRepository(dbConn)
	orderRepo := repository.NewOrderRepository(dbConn)

	productService := service.NewProductService(productRepo, logger)
	statusService := service.NewStatusService(statusRepo, logger)
	cartService := service.NewCartService(cartRepo, productRepo, logger)
	orderService := service.NewOrderService(orderRepo, statusRepo, cartRepo, logger)

	statusCtrl := v1.NewStatusController(statusService)
	productCtrl := v1.NewProductController(productService)
	cartCtrl := v1.NewCartController(cartService)
	orderCtrl := v1.NewOrderController(orderService)

	restServer := server.NewRestServer(cfg.Port, cfg.Mode, statusCtrl, productCtrl, cartCtrl, orderCtrl)

	go restServer.StartServer()
	defer restServer.StopServer()

	sgn := make(chan os.Signal, 1)

	signal.Notify(sgn, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-context.Background().Done():
	case <-sgn:
	}

}

func seedData(con *gorm.DB) {
	con.Create(&entity.Product{
		Title: "Pizza",
		Price: 100,
		Count: 999,
	})

	con.Create(&entity.Product{
		Title: "Sushi",
		Price: 354,
		Count: 999,
	})

	con.Create(&entity.Product{
		Title: "Kebab",
		Price: 50,
		Count: 999,
	})

	con.Create(&entity.Product{
		Title: "Pepsi",
		Price: 80,
		Count: 999,
	})

	con.Create(&entity.Product{
		Title: "Добрый кола 1L",
		Price: 129,
		Count: 999,
	})

	con.Create(&entity.Cart{})

}

func seedStatus(conn *gorm.DB) {
	conn.Create(&entity.Status{Name: "Issued"})
	conn.Create(&entity.Status{Name: "Paid"})
	conn.Create(&entity.Status{Name: "Sent"})
	conn.Create(&entity.Status{Name: "Delivered"})
}
