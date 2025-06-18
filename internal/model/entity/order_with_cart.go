package entity

import "task_cart/internal/model/dto"

type OrderWithCart struct {
	Order    Order
	Status   Status
	Products []struct {
		Product
		Quantity int
	}
}

func (o OrderWithCart) ToDTO() *dto.OrderDTO {
	var mapStruct []dto.InCart

	for _, v := range o.Products {
		mapStruct = append(mapStruct, dto.InCart{
			Quantity: v.Quantity,
			Product:  *v.Product.ToDTO(),
		})
	}

	return &dto.OrderDTO{
		Id:     o.Order.ID,
		Status: *o.Status.ToDTO(),
		InCart: mapStruct,
	}
}
