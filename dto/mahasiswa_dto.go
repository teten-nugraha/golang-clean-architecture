package dto

type MahasiswaDTO struct {

	Nim    string `json:"nim" validate:"required"`
	Nama   string `json:"nama" validate:"required"`
	Phone  string `json:"phone" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	Email  string `json:"email" validate:"required"`

}
