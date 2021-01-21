package routes

import (
	"github.com/labstack/echo"
	"github.com/teten-nugraha/golang-crud/handler"
)

func MahasiswaRoute(routes *echo.Echo, api handler.MahasiswaAPI) {

	mhs := routes.Group("/mahasiswa")
	{
		mhs.GET("/list", api.FindAll)
		mhs.POST("/save", api.SaveOrUpdate)
		mhs.GET("/findByNIM/:nim", api.FindByNIM)
	}
}
