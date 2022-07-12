package main

import (
	admin_models "blog/admin/models"
	"blog/config"
	"net/http"
)

func main() {
	admin_models.Post{}.Migrate()
	http.ListenAndServe(":8080", config.Routes())

}
