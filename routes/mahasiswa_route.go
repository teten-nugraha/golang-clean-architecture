package routes

import (
	"github.com/labstack/echo"
	"github.com/teten-nugraha/golang-crud/handler"
)

func MahasiswaRoute(routes *echo.Echo, api handler.MahasiswaAPI) {

	mhs := routes.Group("/mahasiswa")
	{
		mhs.GET("/", api.FindAll)
		mhs.POST("/", api.SaveOrUpdate)
	}
}
