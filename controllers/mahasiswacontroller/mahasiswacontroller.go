package mahasiswacontroller

import (
	"net/http"
	"text/template"

	"github.com/jeypc/go-crud/entities"
	"github.com/jeypc/go-crud/libraries"
	"github.com/jeypc/go-crud/models"
)

var validation = libraries.NewValidation()
var mahasiswaModel = models.NewMahasiswaModel()

func Index(response http.ResponseWriter, request *http.Request){

	mahasiswa, _ := mahasiswaModel.FindAll()

	data := map[string]interface{}{
		"mahasiswa": mahasiswa,
	}

	temp, err := template.ParseFiles("views/mahasiswa/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}
func Add(response http.ResponseWriter, request *http.Request){

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/mahasiswa/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	}else if request.Method == http.MethodPost {
		
		request.ParseForm()

		var mahasiswa entities.Mahasiswa
		mahasiswa.NamaLengkap = request.Form.Get("nama")
		mahasiswa.Tentang = request.Form.Get("tentang")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(mahasiswa)

		if vErrors != nil {
			data["mahasiswa"] = mahasiswa
			data["validation"] = vErrors
		}else {
			data["pesan"] = "Data berhasil disimpan"
			mahasiswaModel.Create(mahasiswa)
		}

		temp, _ := template.ParseFiles("views/mahasiswa/add.html")
		temp.Execute(response, data)
	}

}
func Edit(response http.ResponseWriter, request *http.Request){
	
}
func Delete(response http.ResponseWriter, request *http.Request){
	
}
