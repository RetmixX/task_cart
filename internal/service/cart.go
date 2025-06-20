package service

import (
	"errors"
	"fmt"
	"task_cart/internal/model/consts"
	"task_cart/internal/model/dto"
	"task_cart/internal/repository"
	"task_cart/pkg/db"
)

type CartInterface interface {
	AddProduct(idProduct uint, quantity int) (*dto.CartDTO, error)
	SeeCart() (*dto.CartDTO, error)
	DeleteProduct(idProduct uint) (*dto.CartDTO, error)
}

// NEED LOGGER!!!!!
type CartService struct {
	cartRepo    repository.CartInterface
	productRepo repository.ProductInterface
}

func NewCartService(cartRepo repository.CartInterface,
	productRepo repository.ProductInterface) *CartService {

	return &CartService{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (c *CartService) AddProduct(idProduct uint, quantity int) (*dto.CartDTO, error) {
	findProduct, err := c.productRepo.ById(idProduct)
	if err != nil {
		if errors.Is(err, db.EntityNotFoundErr) {
			return nil, consts.NotFoundErr
		}
	}

	if findProduct.Count < quantity {
		return nil, consts.InvalidRequest
	}

	cart, err := c.cartRepo.AddProductCart(findProduct, quantity)

	if err != nil {
		return nil, consts.ServerErr
	}

	return cart.ToDTO(), nil

}

func (c *CartService) SeeCart() (*dto.CartDTO, error) {
	cartData, err := c.cartRepo.ViewCart()
	if err != nil {
		return nil, consts.ServerErr
	}

	return cartData.ToDTO(), nil
}

func (c *CartService) DeleteProduct(idProduct uint) (*dto.CartDTO, error) {
	product, err := c.productRepo.ById(idProduct)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, db.EntityNotFoundErr) {
			return nil, consts.NotFoundErr
		}

		return nil, consts.ServerErr
	}

	cart, err := c.cartRepo.DeleteProductFromCart(product)
	if err != nil {
		if errors.Is(err, db.EntityNotFoundErr) {
			return nil, consts.NotFoundErr
		}
		return nil, consts.ServerErr
	}

	return cart.ToDTO(), nil
}
