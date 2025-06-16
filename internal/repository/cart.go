package repository

type CartInterface interface {
	AddProductCart(idProduct uint, quantity int)
}
