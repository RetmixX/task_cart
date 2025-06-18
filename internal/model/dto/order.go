package dto

type OrderDTO struct {
	Id     uint      `json:"id"`
	Status StatusDTO `json:"status"`
	InCart []InCart  `json:"in_cart"`
}

type UpdateOrderDTO struct {
	StatusId uint `json:"status_id" binding:"required"`
}
