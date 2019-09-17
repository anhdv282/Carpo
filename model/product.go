package model

type Product struct {
	// Tao shema blog
	TableName []byte `sql:"product.product"`
	// ID
	Id int `sql:",pk"`
	// Tên sản phẩm
	Name string
	// Mô tả
	Description string
	// Giá
	Price int
	// Ảnh
	Image string
}

type CreateProductForm struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	Price       int    `form:"price"`
}
