package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/teten-nugraha/golang-crud/domain"
)

type MahasiswaRepositoryContract interface {
	SaveOrUpdate(mahasiawa domain.Mahasiswa) (domain.Mahasiswa, error)
	FindAll() [] domain.Mahasiswa
	FindByID(id string) domain.Mahasiswa
	FindByNim(nim string) domain.Mahasiswa
	DeleteMahasiswa(mahasiswa domain.Mahasiswa) error
}

type MahasiswaRepository struct {
	DB *gorm.DB
}

func ProviderMahasiswaRepository(DB *gorm.DB) MahasiswaRepository {
	return MahasiswaRepository{DB: DB}
}

// implementation
func (m *MahasiswaRepository) SaveOrUpdate(mahasiswa domain.Mahasiswa) (domain.Mahasiswa, error) {
	if err := m.DB.Create(&mahasiswa).Error; err != nil {
		return mahasiswa, err
	}
	return mahasiswa, nil
}

func (m *MahasiswaRepository) FindAll() [] domain.Mahasiswa {
	var mahasiswas []domain.Mahasiswa

	m.DB.Find(&mahasiswas)

	return mahasiswas
}

func(m *MahasiswaRepository) FindByID(id string) domain.Mahasiswa {
	var mahasiswa domain.Mahasiswa

	m.DB.Where("id =? ", id).Find(&mahasiswa)

	return mahasiswa
}


func(m *MahasiswaRepository) FindByNim(nim string) domain.Mahasiswa {
	var mahasiswa domain.Mahasiswa

	m.DB.Where("nim =? ", nim).Find(&mahasiswa)

	return mahasiswa
}

func (m *MahasiswaRepository) DeleteMahasiswa(mahasiswa domain.Mahasiswa) error {
	if err := m.DB.Delete(&mahasiswa).Error; err!= nil  {
		return err
	}
	return nil
}