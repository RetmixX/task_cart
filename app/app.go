package app

import (
	"context"
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

func Run(ctx context.Context) {
	cfg := config.MustLoad()
	dbConn := db.MustStartDB(&cfg.DbConf)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	defer db.MustCloseDB(dbConn)
	err := dbConn.SetupJoinTable(&entity.Cart{}, "Products", &entity.CartProduct{})

	if err != nil {
		panic(err)
	}

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
	case <-ctx.Done():
	case <-sgn:
	}
}
