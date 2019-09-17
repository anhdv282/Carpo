package controller

import (
	"io"
	"log"
	"os"
	"sample-project/model"

	"github.com/kataras/iris"
)

func (c *Controller) AdminGetPosts(ctx iris.Context) {
	// Viết code ở đây

	ctx.View("/admin/blog/index.html")
}

func (c *Controller) AdminGetCreatePostPage(ctx iris.Context) {
	// Viết code ở đây

	ctx.View("/admin/blog/create.html")
}

func (c *Controller) AdminCreateProduct(ctx iris.Context) {
	ctx.View("/admin/product/create.html")
}

func (c *Controller) AdminCreateProductPost(ctx iris.Context) {
	// Đọc form dữ liệu
	var form model.CreateProductForm
	err := ctx.ReadForm(&form)
	if err != nil {
		log.Println(err)
		return
	}

	// Upload ảnh
	file, info, err := ctx.FormFile("image")
	if err != nil {
		log.Println("-- 1: ",err)
		return
	}
	defer file.Close()
	fname := info.Filename
	out, err := os.OpenFile("./uploads/"+fname, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("-- 2: ",err)
		return
	}
	defer out.Close()
	io.Copy(out, file)

	// Insert dữ liệu
	var product model.Product
	product.Name = form.Name
	product.Description = form.Description
	product.Price = form.Price 
	product.Image = "/uploads/"+fname
	
	tx, err := c.DB.Begin()
	if err != nil {
		log.Println(err)
		return 
	}

	err = tx.Insert(&product)
	if err != nil {
		log.Println(err)
		return 
	}

	tx.Commit()
	
	log.Println("----------------------------------------")
	log.Println("thêm sản phẩm thành công rồi")
}
