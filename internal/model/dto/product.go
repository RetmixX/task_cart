package dto

type ProductDTO struct {
	Id    uint    `json:"id"`
	Title string  `json:"title"`
	Price float32 `json:"price"`
}

type InCart struct {
	Quantity int        `json:"quantity"`
	Product  ProductDTO `json:"product"`
}

type ProductWithCountDTO struct {
	ProductDTO
	Count int `json:"count"`
}
