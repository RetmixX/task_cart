package main

import (
	"fmt"
	"task_cart/config"
	"task_cart/internal/model/entity"
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
	defer db.MustCloseDB(dbConn, nil)
	err := dbConn.SetupJoinTable(&entity.Cart{}, "Products", &entity.CartProduct{})
	if err != nil {
		panic(err)
	}
	if err := dbConn.AutoMigrate(&entity.Product{}, &entity.Status{}, &entity.Cart{}, &entity.Order{}); err != nil {
		panic(err)
	}

	fmt.Println("Success migrate")

}
