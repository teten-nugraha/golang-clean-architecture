package mapper

import (
	"github.com/teten-nugraha/golang-crud/domain"
	"github.com/teten-nugraha/golang-crud/dto"
)

// From DTO to Domain
func ToMahasiswaDomain(dto dto.MahasiswaDTO) domain.Mahasiswa {
	return domain.Mahasiswa{
		Nim: dto.Nim,
		Nama: dto.Nama,
		Phone: dto.Phone,
		Alamat: dto.Alamat,
		Email: dto.Email,
	}
}

func ToMahasiswaDomainList(dtos []dto.MahasiswaDTO) []domain.Mahasiswa {
	mahasiswas := make([]domain.Mahasiswa, len(dtos))

	for i, itm := range dtos {
		mahasiswas[i] = ToMahasiswaDomain(itm)
	}
	return mahasiswas
}

// from domain to DTO
// From DTO to Domain
func ToMahasiswaDto(mahasiswa domain.Mahasiswa) dto.MahasiswaDTO {
	return dto.MahasiswaDTO{
		Nim: mahasiswa.Nim,
		Nama: mahasiswa.Nama,
		Phone: mahasiswa.Phone,
		Alamat: mahasiswa.Alamat,
		Email: mahasiswa.Email,
	}
}

func ToMahasiswaDtoList(mahasiswas []domain.Mahasiswa) []dto.MahasiswaDTO {
	dtos := make([]dto.MahasiswaDTO, len(mahasiswas))

	for i, itm := range mahasiswas {
		dtos[i] = ToMahasiswaDto(itm)
	}

	return dtos
}
