package entity

import "task_cart/internal/model/dto"

type CartWithProducts struct {
	Cart     Cart
	Products []struct {
		Product
		Quantity int
	}
}

func (p CartWithProducts) ToDTO() *dto.CartDTO {
	var mapStruct []dto.InCart

	for _, v := range p.Products {
		mapStruct = append(mapStruct, dto.InCart{
			Quantity: v.Quantity,
			Product:  *v.Product.ToDTO(),
		})
	}

	return &dto.CartDTO{
		Id:     p.Cart.ID,
		InCart: mapStruct,
	}
}
