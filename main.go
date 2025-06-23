package main

import (
	"context"
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
	"task_cart/pkg/log/sl"
)

func main() {

	cfg := config.MustLoad()
	dbConn := db.MustStartDB(&cfg.DbConf)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	defer db.MustCloseDB(dbConn, nil)
	err := dbConn.SetupJoinTable(&entity.Cart{}, "Products", &entity.CartProduct{})

	if err != nil {
		panic(err)
	}
	if err := dbConn.AutoMigrate(&entity.Product{}, &entity.Status{}, &entity.Cart{}, &entity.Order{}); err != nil {
		panic(err)
	}

	logger.Info("Success migrate")

	seedData(dbConn, logger)
	seedStatus(dbConn, logger)
	logger.Info("Start service...")
	productRepo := repository.NewProductRepository(dbConn)
	statusRepo := repository.NewStatusRepository(dbConn)
	cartRepo := repository.NewCartRepository(dbConn)
	orderRepo := repository.NewOrderRepository(dbConn)

	productService := service.NewProductService(productRepo, logger)
	statusService := service.NewStatusService(statusRepo, logger)
	cartService := service.NewCartService(cartRepo, productRepo, logger)
	orderService := service.NewOrderService(orderRepo, statusRepo, cartRepo, logger)

	statusCtrl := v1.NewStatusController(statusService, logger)
	productCtrl := v1.NewProductController(productService, logger)
	cartCtrl := v1.NewCartController(cartService, logger)
	orderCtrl := v1.NewOrderController(orderService, logger)

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

func seedData(con *gorm.DB, log *slog.Logger) {
	var products []entity.Product
	log.Info("Start seed products data")
	if err := con.Find(&products).Error; err != nil {
		log.Error("fail seed, error: ", sl.Err(err))
		return
	}

	if len(products) > 0 {
		log.Info("Table is fill, skip....")
		return
	}
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
	log.Info("End seed data products")

}

func seedStatus(conn *gorm.DB, log *slog.Logger) {
	var statuses []entity.Status
	log.Info("Start seed statuses order")
	if err := conn.Find(&statuses).Error; err != nil {
		log.Error("fail get data from table status: ", sl.Err(err))
		return
	}

	if len(statuses) > 0 {
		log.Info("Table is fill, skip....")
		return
	}
	conn.Create(&entity.Status{Name: "Issued"})
	conn.Create(&entity.Status{Name: "Paid"})
	conn.Create(&entity.Status{Name: "Sent"})
	conn.Create(&entity.Status{Name: "Delivered"})
	log.Info("End seed data status")
}
