package entities

type Mahasiswa struct {
	Id          int64
	NamaLengkap string `validate:"required" label:"Nama Lengkap"`
	Tentang     string `validate:"required" label:"Tentang"`
}