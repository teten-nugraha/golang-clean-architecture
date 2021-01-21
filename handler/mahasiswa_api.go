package handler

import (
	"github.com/labstack/echo"
	"github.com/teten-nugraha/golang-crud/dto"
	_ "github.com/teten-nugraha/golang-crud/dto"
	"github.com/teten-nugraha/golang-crud/service"
	_ "github.com/labstack/echo"
	"net/http"
)

type MahasiswaAPI struct {
	MahasiswaService service.MahasiswaService
}

func ProviderMahasiswaAPI(k service.MahasiswaService) MahasiswaAPI {
	return MahasiswaAPI{MahasiswaService: k}
}

// implementasi
func (m *MahasiswaAPI) FindAll(e echo.Context) error {

	mahasiswas := m.MahasiswaService.FindAll()

	if len(mahasiswas) == 0 {
		return SuccessResponse(e, http.StatusNoContent, mahasiswas)
	}

	return SuccessResponse(e, http.StatusOK, mahasiswas)
}

func (m *MahasiswaAPI) SaveOrUpdate(e echo.Context) error {
	var newDto dto.MahasiswaDTO

	newDto.Nim = e.FormValue("Nim")
	newDto.Nama = e.FormValue("Nama")
	newDto.Phone = e.FormValue("Phone")
	newDto.Alamat = e.FormValue("Alamat")
	newDto.Email = e.FormValue("Email")

	res, err := m.MahasiswaService.SaveOrUpdate(newDto)
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(e, http.StatusOK, res)
}
