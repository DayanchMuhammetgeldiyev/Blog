package config

import (
	admin "blog/admin/controllers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	r.GET("/admin", admin.Dashboard{}.Index)
	//Add-BLOG
	r.GET("/admin/content-add", admin.Dashboard{}.NewItem)
	r.POST("/admin", admin.Dashboard{}.Add)
	r.GET("/admin/delete/:id", admin.Dashboard{}.Delete)
	r.GET("/admin/edit/:id", admin.Dashboard{}.Edit)
	r.POST("/admin/update/:id", admin.Dashboard{}.Update)
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads"))
	return r
}
