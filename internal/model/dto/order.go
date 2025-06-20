package dto

type OrderDTO struct {
	Id     uint      `json:"id"`
	Status StatusDTO `json:"status"`
	InCart []InCart  `json:"in_cart"`
	Amount float32   `json:"amount"`
}

type UpdateOrderDTO struct {
	StatusId uint `json:"status_id" binding:"required"`
}
