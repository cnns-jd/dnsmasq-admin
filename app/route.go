package app

import "dnsmasq-admin.io/app/core"

func init() {
	// engine.Use(core.Auth())
	engine.GET("/",core.List)
	engine.GET("/item/:name",core.Item)
	engine.POST("/item/edit/:name",core.Edit)
	engine.GET("/add",core.Add)
	engine.POST("/item/add",core.ItemAdd)
}
