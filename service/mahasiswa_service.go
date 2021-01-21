package service

import (
	"errors"
	"github.com/teten-nugraha/golang-crud/domain"
	"github.com/teten-nugraha/golang-crud/dto"
	"github.com/teten-nugraha/golang-crud/mapper"
	"github.com/teten-nugraha/golang-crud/repository"
)

type MahasiswaRepositoryContract interface {
	SaveOrUpdate(dto dto.MahasiswaDTO) (dto.MahasiswaDTO, error)
	FindAll() [] dto.MahasiswaDTO
	FindByNim(nim string) dto.MahasiswaDTO
	DeleteMahasiswa(id string) error
}

type MahasiswaService struct {
	MahasiswaRepository repository.MahasiswaRepository
}

func ProviderMahasiswaService(m repository.MahasiswaRepository) MahasiswaService {
	return MahasiswaService{
		MahasiswaRepository: m,
	}
}

// implementation
func (m *MahasiswaService) SaveOrUpdate(dto dto.MahasiswaDTO) (dto.MahasiswaDTO, error) {

	mahasiswaEntity := mapper.ToMahasiswaDomain(dto)

	mahasiswa, err := m.MahasiswaRepository.SaveOrUpdate(mahasiswaEntity)

	return mapper.ToMahasiswaDto(mahasiswa), err
}

func (m *MahasiswaService) FindAll() [] dto.MahasiswaDTO {

	datas := m.MahasiswaRepository.FindAll()

	return mapper.ToMahasiswaDtoList(datas)
}

func (m *MahasiswaService) FindByNim(nim string) dto.MahasiswaDTO {
	mahasiswa := m.MahasiswaRepository.FindByNim(nim)

	return mapper.ToMahasiswaDto(mahasiswa)
}

func (m *MahasiswaService) DeleteMahasiswa(id string) error {

	mahasiswa := m.MahasiswaRepository.FindByID(id)

	if(mahasiswa == (domain.Mahasiswa{})) {
		return errors.New("Mahasiswa Tidak ada")
	}

	err := m.MahasiswaRepository.DeleteMahasiswa(mahasiswa)
	if err != nil {
		return err
	}

	return nil
}