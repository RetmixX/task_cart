package dto

type CartDTO struct {
	Id     uint     `json:"id"`
	InCart []InCart `json:"in_cart"`
}

type AddProductCartDTO struct {
	ProductId uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

type DeleteProductDTO struct {
	ProductId uint `json:"product_id" binding:"required"`
}
