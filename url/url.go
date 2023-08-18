package url

import (
	"github.com/AdeCandra12/pemrog3-ulbi/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth

	page.Get("/", controller.Homepage) //ujicoba panggil package musik
	// page.Get("/presensi", controller.GetPresensi)

	page.Get("/presensi", controller.GetAllPresensi)    //menampilkan seluruh data presensi
	page.Get("/presensi/:id", controller.GetPresensiID) //menampilkan data presensi berdasarkan id
	page.Get("/surat", controller.GetAllSurat)          //menampilkan seluruh data presensi
	page.Get("/disposisi", controller.GetAllDisposisi)  //menampilkan seluruh data presensi
	// get monitoring
	page.Get("/all-mahasiswa", controller.GetAllMahasiswa)
	page.Get("/all-orangtua", controller.GetAllOrangTua)
	page.Get("/all-matakuliah", controller.GetAllMatakuliah)
	page.Get("/all-absensi", controller.GetAllAbsensi)
	page.Get("/all-nilai", controller.GetAllNilai)
	// get from id monitoring
	page.Get("/mahasiswa/:id", controller.GetMahasiswaFromID)
	page.Get("/orangtua/:id", controller.GetOrangTuaFromID)
	page.Get("/matakuliah/:id", controller.GetMatakuliahFromID)
	page.Get("/absensi/:id", controller.GetAbsensiFromID)
	page.Get("/nilai/:id", controller.GetNilaiFromID)
	// insert monitoring
	page.Post("/ins-mahasiswa", controller.InsertMahasiswa)
	page.Post("/ins-orangtua", controller.InsertOrangTua)
	page.Post("/ins-matakuliah", controller.InsertMatakuliah)
	page.Post("/ins-absensi", controller.InsertAbsensi)
	page.Post("/ins-nilai", controller.InsertNilai)
	// Update Monitoring
	page.Put("/upd-mahasiswa/:id", controller.UpdateMahasiswa)
	page.Put("/upd-orangtua/:id", controller.UpdateOrangTua)
	page.Put("/upd-matakuliah/:id", controller.UpdateMatakuliah)
	page.Put("/upd-absensi/:id", controller.UpdateAbsensi)
	page.Put("/upd-nilai/:id", controller.UpdateNilai)
	// Delete Monitoring
	page.Delete("/delete-mahasiswa/:id", controller.DeleteMahasiswaByID)
	page.Delete("/delete-orangtua/:id", controller.DeleteOrangTuaByID)
	page.Delete("/delete-matakuliah/:id", controller.DeleteMatakuliahByID)
	page.Delete("/delete-absensi/:id", controller.DeleteAbsensiByID)
	page.Delete("/delete-nilai/:id", controller.DeleteNilaiByID)
	// surat
	page.Get("/surat", controller.GetAllSurat)                //menampilkan seluruh data presensi
	page.Get("/disposisi", controller.GetAllDisposisi)        //menampilkan seluruh data presens
	page.Post("/ins", controller.InsertData)                  // insert data
	page.Put("/upd/:id", controller.UpdateData)               // update data
	page.Delete("/delete/:id", controller.DeletePresensiByID) // delete data
	page.Get("/docs/*", swagger.HandlerDefault)
}
