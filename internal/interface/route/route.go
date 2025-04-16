package route

import (
	"api/internal/interface/controller"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRoutes(router *mux.Router, db *gorm.DB) {
	absensiController := controller.NewAbsensiController(db)
	beasiswaController := controller.NewBeasiswaController(db)
	dosenController := controller.NewDosenController(db)
	jadwalKuliahController := controller.NewJadwalKuliahController(db)
	mahasiswaController := controller.NewMahasiswaController(db)
	mataKuliahController := controller.NewMataKuliahController(db)
	nilaiController := controller.NewNilaiController(db)
	pengumumanController := controller.NewPengumumanController(db)
	ruangKelasController := controller.NewRuangKelasController(db)
	tugasController := controller.NewTugasController(db)

	router.HandleFunc("/api/absensi", absensiController.GetAbsensis).Methods("GET")
	router.HandleFunc("/api/absensi", absensiController.CreateAbsensi).Methods("POST")
	router.HandleFunc("/api/absensi/{id}", absensiController.UpdateAbsensi).Methods("PUT")
	router.HandleFunc("/api/absensi/{id}", absensiController.DeleteAbsensi).Methods("DELETE")

	router.HandleFunc("/api/beasiswa", beasiswaController.GetBeasiswas).Methods("GET")
	router.HandleFunc("/api/beasiswa", beasiswaController.CreateBeasiswa).Methods("POST")
	router.HandleFunc("/api/beasiswa/{id}", beasiswaController.UpdateBeasiswa).Methods("PUT")
	router.HandleFunc("/api/beasiswa/{id}", beasiswaController.DeleteBeasiswa).Methods("DELETE")
	
    router.HandleFunc("/api/dosen", dosenController.GetDosens).Methods("GET")
	router.HandleFunc("/api/dosen", dosenController.CreateDosen).Methods("POST")
	router.HandleFunc("/api/dosen/{id}", dosenController.UpdateDosen).Methods("PUT")
	router.HandleFunc("/api/dosen/{id}", dosenController.DeleteDosen).Methods("DELETE")

    router.HandleFunc("/api/jadwalKuliah", jadwalKuliahController.GetJadwalKuliahs).Methods("GET")
	router.HandleFunc("/api/jadwalKuliah", jadwalKuliahController.CreateJadwalKuliah).Methods("POST")
	router.HandleFunc("/api/jadwalKuliah/{id}", jadwalKuliahController.UpdateJadwalKuliah).Methods("PUT")
	router.HandleFunc("/api/jadwalKuliah/{id}", jadwalKuliahController.DeleteJadwalKuliah).Methods("DELETE")

    router.HandleFunc("/api/mahasiswa", mahasiswaController.GetMahasiswas).Methods("GET")
	router.HandleFunc("/api/mahasiswa", mahasiswaController.CreateMahasiswa).Methods("POST")
	router.HandleFunc("/api/mahasiswa/{id}", mahasiswaController.UpdateMahasiswa).Methods("PUT")
	router.HandleFunc("/api/mahasiswa/{id}", mahasiswaController.DeleteMahasiswa).Methods("DELETE")

    router.HandleFunc("/api/mataKuliah", mataKuliahController.GetMataKuliahs).Methods("GET")
	router.HandleFunc("/api/mataKuliah", mataKuliahController.CreateMataKuliah).Methods("POST")
	router.HandleFunc("/api/mataKuliah/{id}", mataKuliahController.UpdateMataKuliah).Methods("PUT")
	router.HandleFunc("/api/mataKuliah/{id}", mataKuliahController.DeleteMataKuliah).Methods("DELETE")

    router.HandleFunc("/api/nilai", nilaiController.GetNilais).Methods("GET")
	router.HandleFunc("/api/nilai", nilaiController.CreateNilai).Methods("POST")
	router.HandleFunc("/api/nilai/{id}", nilaiController.UpdateNilai).Methods("PUT")
	router.HandleFunc("/api/nilai/{id}", nilaiController.DeleteNilai).Methods("DELETE")

    router.HandleFunc("/api/pengumuman", pengumumanController.GetPengumumans).Methods("GET")
	router.HandleFunc("/api/pengumuman", pengumumanController.CreatePengumuman).Methods("POST")
	router.HandleFunc("/api/pengumuman/{id}", pengumumanController.UpdatePengumuman).Methods("PUT")
	router.HandleFunc("/api/pengumuman/{id}", pengumumanController.DeletePengumuman).Methods("DELETE")
    
    router.HandleFunc("/api/ruangKelas", ruangKelasController.GetRuangKelass).Methods("GET")
	router.HandleFunc("/api/ruangKelas", ruangKelasController.CreateRuangKelas).Methods("POST")
	router.HandleFunc("/api/ruangKelas/{id}", ruangKelasController.UpdateRuangKelas).Methods("PUT")
	router.HandleFunc("/api/ruangKelas/{id}", ruangKelasController.DeleteRuangKelas).Methods("DELETE")

    router.HandleFunc("/api/tugas", tugasController.GetTugass).Methods("GET")
	router.HandleFunc("/api/tugas", tugasController.CreateTugas).Methods("POST")
	router.HandleFunc("/api/tugas/{id}", tugasController.UpdateTugas).Methods("PUT")
	router.HandleFunc("/api/tugas/{id}", tugasController.DeleteTugas).Methods("DELETE")
}
