package injection

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/teten-nugraha/golang-crud/handler"
	"github.com/teten-nugraha/golang-crud/repository"
	"github.com/teten-nugraha/golang-crud/service"
)

func initMahasiswaAPI(db *gorm.DB) handler.MahasiswaAPI {
	wire.Build(repository.ProviderMahasiswaRepository, service.ProviderMahasiswaService, handler.ProviderMahasiswaAPI)

	return handler.MahasiswaAPI{}
}