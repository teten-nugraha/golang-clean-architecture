package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	db2 "github.com/teten-nugraha/golang-crud/db"
	"github.com/teten-nugraha/golang-crud/injection"
)

func Init() *echo.Echo {
	dbConfig := db2.InitDB()

	mahasiswaAPI := injection.InitMahasiswaAPI(dbConfig)

	routes := echo.New()

	// set logger
	routes.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, time:${time_unix}, uri=${uri}, status=${status}, error=${error}, latency_human=${latency}, bytes_in=${bytes_in}, bytes_out=${bytes_out} \n",
	}))

	// Gzip Compression
	routes.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	//MahasiswaRoute(routes, mahasiswaAPI)

	routes.GET("/mahasiswalist", mahasiswaAPI.FindAll)
	routes.POST("/mahasiswa", mahasiswaAPI.SaveOrUpdate)

	return routes
}