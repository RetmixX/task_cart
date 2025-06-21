package service

import (
	"errors"
	"fmt"
	"log/slog"
	"task_cart/internal/model/consts"
	"task_cart/internal/model/dto"
	"task_cart/internal/repository"
	"task_cart/pkg/db"
	"task_cart/pkg/log/sl"
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
	log         *slog.Logger
}

func NewCartService(cartRepo repository.CartInterface,
	productRepo repository.ProductInterface, log *slog.Logger) *CartService {

	return &CartService{
		cartRepo:    cartRepo,
		productRepo: productRepo,
		log:         log,
	}
}

func (c *CartService) AddProduct(idProduct uint, quantity int) (*dto.CartDTO, error) {
	const op = "service.Cart.AddProduct"
	log := c.log.With(slog.String("op", op))
	log.Info("Start AddProduct()")
	findProduct, err := c.productRepo.ById(idProduct)
	if err != nil {
		if errors.Is(err, db.EntityNotFoundErr) {
			log.Warn("product not found", sl.Err(err))
			return nil, consts.NotFoundErr
		}
		log.Error("fail find product by id", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if findProduct.Count < quantity {
		return nil, consts.InvalidRequest
	}

	cart, err := c.cartRepo.AddProductCart(findProduct, quantity)

	if err != nil {
		log.Error("fail add product to cart", sl.Err(err))
		return nil, consts.ServerErr
	}
	log.Info("End AddProduct()")
	return cart.ToDTO(), nil

}

func (c *CartService) SeeCart() (*dto.CartDTO, error) {
	const op = "service.Cart.SeeCart"
	log := c.log.With(slog.String("op", op))
	log.Info("Start SeeCart()")
	cartData, err := c.cartRepo.ViewCart()
	if err != nil {
		log.Error("fail get data cart", sl.Err(err))
		return nil, consts.ServerErr
	}
	log.Info("End SeeCart()")
	return cartData.ToDTO(), nil
}

func (c *CartService) DeleteProduct(idProduct uint) (*dto.CartDTO, error) {
	const op = "service.Cart.DeleteProduct"
	log := c.log.With(slog.String("op", op))
	log.Info("Start DeleteProduct()")
	product, err := c.productRepo.ById(idProduct)
	if err != nil {
		log.Error("fail ")
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
