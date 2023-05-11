package models

import (
	"database/sql"
	"fmt"

	"github.com/jeypc/go-crud/config"
	"github.com/jeypc/go-crud/entities"
)

type MahasiswaModel struct {
	conn *sql.DB
}

func NewMahasiswaModel() *MahasiswaModel {
	conn, err := config.DBCOnnection()
	if err != nil {
		panic(err)
	}

	return &MahasiswaModel{
		conn: conn,
	}
}

func (p *MahasiswaModel) FindAll() ([]entities.Mahasiswa, error) {
	
	rows, err := p.conn.Query("SELECT * FROM mahasiswa")
	if err != nil {
		return []entities.Mahasiswa{}, err
	}
	defer rows.Close()

	var dataMahasiswa []entities.Mahasiswa
	for rows.Next() {
		var mahasiswa entities.Mahasiswa
		rows.Scan(&mahasiswa.Id, 
			&mahasiswa.NamaLengkap, 
			&mahasiswa.Tentang)
		dataMahasiswa = append(dataMahasiswa, mahasiswa)
	}

	return dataMahasiswa, nil

}

func (p *MahasiswaModel) Create(mahasiswa entities.Mahasiswa) bool {
	result, err := p.conn.Exec("INSERT INTO mahasiswa (nama, tentang) values (?,?)", mahasiswa.NamaLengkap, mahasiswa.Tentang)
	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}