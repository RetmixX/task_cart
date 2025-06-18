package entity

type CartProduct struct {
	ProductId uint `gorm:"primaryKey"`
	CartId    uint `gorm:"primaryKey"`
	Quantity  int  `gorm:"not null"`
}
