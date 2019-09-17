package controller

import "github.com/kataras/iris"

func (c *Controller) GetBlog(ctx iris.Context) {
	// Viết code ở đây 

	ctx.View("/blog/index.html")
}

func (c *Controller) GetPostById(ctx iris.Context) {
	// Viết code ở đây
	
	ctx.View("/blog/post.html")
}

func (c *Controller) GetHome(ctx iris.Context) {
	ctx.View("index.html")
}

func (c *Controller) GetShop(ctx iris.Context) {
	ctx.View("/shop/index.html")
}