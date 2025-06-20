package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"task_cart/internal/model/entity"
	"task_cart/pkg/db"
)

type CartInterface interface {
	AddProductCart(product *entity.Product, quantity int) (*entity.CartWithProducts, error)
	ViewCart() (*entity.CartWithProducts, error)
	DeleteProductFromCart(product *entity.Product) (*entity.CartWithProducts, error)
}

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(conn *gorm.DB) *CartRepository {
	return &CartRepository{db: conn}
}

func (c *CartRepository) AddProductCart(product *entity.Product, quantity int) (*entity.CartWithProducts, error) {
	const op = "repository.cart.AddProductCart"
	var currentCart entity.Cart
	if err := c.db.Where("is_ordered", false).First(&currentCart).Error; err != nil {
		return nil, fmt.Errorf("%s: can't get current cart: %w", op, err)
	}

	var cartProduct entity.CartProduct

	existsProductInCart := c.db.Where("cart_id = ? and product_id = ?", currentCart.ID, product.ID).First(&cartProduct).Error
	if existsProductInCart != nil && errors.Is(existsProductInCart, gorm.ErrRecordNotFound) {
		addProduct := entity.CartProduct{
			ProductId: product.ID,
			CartId:    currentCart.ID,
			Quantity:  quantity,
		}

		if err := c.db.Create(&addProduct).Error; err != nil {
			return nil, fmt.Errorf("%s: can't add product to cart: %w", op, err)
		}
	} else {
		newQuantity := cartProduct.Quantity + quantity
		if err := c.db.Model(&cartProduct).Update("quantity", newQuantity).Error; err != nil {
			return nil, fmt.Errorf("%s: can't update quntity: %w", op, err)
		}
	}

	if err := c.db.Model(product).Update("count", product.Count-quantity).Error; err != nil {
		return nil, fmt.Errorf("%s: can't update count in product: %w", op, err)
	}

	return c.ViewCart()
}

func (c *CartRepository) ViewCart() (*entity.CartWithProducts, error) {
	const op = "repository.cart.ViewCart"
	var result entity.CartWithProducts

	if err := c.db.Where("is_ordered", false).First(&result.Cart).Error; err != nil {
		return nil, fmt.Errorf("%s: can't get current cart: %w", op, err)
	}

	err := c.db.Table("cart_products").Select("products.*, cart_products.quantity").
		Joins("join products on products.id = cart_products.product_id").
		Where("cart_products.cart_id = ?", result.Cart.ID).Scan(&result.Products).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("%s: can't get cart data: %w", op, err)
	}

	return &result, nil
}

func (c *CartRepository) DeleteProductFromCart(product *entity.Product) (*entity.CartWithProducts, error) {
	const op = "repository.cart.DeleteProductFromCart"
	var currentCart entity.Cart
	if err := c.db.Where("is_ordered", false).First(&currentCart).Error; err != nil {
		return nil, fmt.Errorf("%s: can't get current cart: %w", op, err)
	}

	var selectCartProduct entity.CartProduct
	if err := c.db.Where("product_id = ? and cart_id = ?",
		product.ID, currentCart.ID).First(&selectCartProduct).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, db.EntityNotFoundErr
		}
		return nil, fmt.Errorf("%s: can't get position in cart: %w", op, err)
	}
	quantity := selectCartProduct.Quantity
	if err := c.db.Delete(&selectCartProduct).Error; err != nil {
		return nil, fmt.Errorf("%s: can't remove product from cart: %w", op, err)
	}

	if err := c.db.Model(&product).Update("count", product.Count+quantity).Error; err != nil {
		return nil, fmt.Errorf("%s: error update info in product: %w", op, err)
	}

	return c.ViewCart()
}
