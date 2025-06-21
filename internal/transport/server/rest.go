package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "task_cart/internal/transport/rest/api/v1"
	ctrl "task_cart/internal/transport/rest/controllers/v1"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "task_cart/docs"
)

type TransportInterface interface {
	StartServer()
	StopServer()
}

type RestServer struct {
	server *http.Server
}

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.

// @host		localhost:3000
// @BasePath	/api/v1
func NewRestServer(port string, serverMode string,
	statusCtrl *ctrl.StatusController, productCtrl *ctrl.ProductController,
	cartCtrl *ctrl.CartController, orderCtrl *ctrl.OrderController) *RestServer {

	engine := gin.Default()
	gin.SetMode(serverMode)

	v1.RegisterStatusRoutes(engine, statusCtrl)
	v1.RegisterProductRoutes(engine, productCtrl)
	v1.RegisterCartRoutes(engine, cartCtrl)
	v1.RegisterOrderRoutes(engine, orderCtrl)

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server := &http.Server{
		Addr:    port,
		Handler: engine,
	}

	return &RestServer{server: server}
}

func (r *RestServer) StartServer() {
	if err := r.server.ListenAndServe(); err != nil {
		fmt.Println("[ERR]: can't start server: ", err)
		panic(err)
	}

	fmt.Println("Server start")
}

func (r *RestServer) StopServer() {
	if err := r.server.Shutdown(context.Background()); err != nil {
		fmt.Println("[ERR]: can't stop server: ", err)
		panic(err)
	}

	fmt.Println("Server stop")
}
