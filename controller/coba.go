package controller

import (
	"errors"
	"fmt"
	"net/http"

	inimodelproyek1 "github.com/AdeCandra12/BE_proyek1/model"
	inimodulproyek1 "github.com/AdeCandra12/BE_proyek1/module"
	"github.com/AdeCandra12/pemrog3-ulbi/config"
	inimodul "github.com/AdeCandra12/surat/module"
	"github.com/aiteung/musik"
	cek "github.com/aiteung/presensi"
	"github.com/gofiber/fiber/v2"
	inimodellatihan "github.com/indrariksa/be_presensi/model"
	inimodullatihan "github.com/indrariksa/be_presensi/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetPresensi(c *fiber.Ctx) error {
	ps := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
	return c.JSON(ps)
}

// GetAllPresensi godoc
// @Summary Get All Data Presensi.
// @Description Mengambil semua data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Success 200 {object} Presensi
// @Router /presensi [get]
func GetAllPresensi(c *fiber.Ctx) error {
	ps := inimodullatihan.GetAllPresensi(config.Ulbimongoconn, "presensi")
	return c.JSON(ps)
}

// GetPresensiID godoc
// @Summary Get By ID Data Presensi.
// @Description Ambil per ID data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /presensi/{id} [get]
func GetPresensiID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := inimodullatihan.GetPresensiFromID(objID, config.Ulbimongoconn, "presensi")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

func GetAllSurat(c *fiber.Ctx) error {
	ps := inimodul.GetAllSurat(config.Ulbimongoconn2, "surat")
	return c.JSON(ps)
}
func GetAllDisposisi(c *fiber.Ctx) error {
	ps := inimodul.GetAllDisposisi(config.Ulbimongoconn2, "disposisi")
	return c.JSON(ps)
}

// InsertData godoc
// @Summary Insert data presensi.
// @Description Input data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param request body Presensi true "Payload Body [RAW]"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 500
// @Router /ins [post]
func InsertData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var presensi inimodellatihan.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodullatihan.InsertPresensi(db, "presensi",
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// update data
// UpdateData godoc
// @Summary Update data presensi.
// @Description Ubah data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Presensi true "Payload Body [RAW]"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 500
// @Router /upd/{id} [put]
func UpdateData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var presensi inimodellatihan.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = inimodullatihan.UpdatePresensi(db, "presensi",
		objectID,
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// delete data
// DeletePresensiByID godoc
// @Summary Delete data presensi.
// @Description Hapus data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete/{id} [delete]
func DeletePresensiByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = inimodullatihan.DeletePresensiByID(objID, config.Ulbimongoconn, "presensi")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}

// Get All monitoring

// GetAllMahasiswa godoc
// @Summary Get All Data Mahasiswa.
// @Description Mengambil semua data mahasiswa.
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Success 200 {object} Mahasiswa
// @Router /all-mahasiswa [get]
func GetAllMahasiswa(c *fiber.Ctx) error {
	ms := inimodulproyek1.GetAllMahasiswa(config.Ulbimongoconn, "mahasiswa")
	return c.JSON(ms)
}

// GetAllOrangTua godoc
// @Summary Get All Data OrangTua.
// @Description Mengambil semua data orangtua.
// @Tags OrangTua
// @Accept json
// @Produce json
// @Success 200 {object} OrangTua
// @Router /all-orangtua [get]
func GetAllOrangTua(c *fiber.Ctx) error {
	ot := inimodulproyek1.GetAllOrangTua(config.Ulbimongoconn, "orangtua")
	return c.JSON(ot)
}

// GetAllMatakuliah godoc
// @Summary Get All Data Matakuliah.
// @Description Mengambil semua data matakuliah.
// @Tags Matakuliah
// @Accept json
// @Produce json
// @Success 200 {object} Matakuliah
// @Router /all-matakuliah [get]
func GetAllMatakuliah(c *fiber.Ctx) error {
	mk := inimodulproyek1.GetAllMatakuliah(config.Ulbimongoconn, "matakuliah")
	return c.JSON(mk)
}

// GetAllAbsensi godoc
// @Summary Get All Data Absensi.
// @Description Mengambil semua data absensi.
// @Tags Absensi
// @Accept json
// @Produce json
// @Success 200 {object} Absensi
// @Router /all-absensi [get]
func GetAllAbsensi(c *fiber.Ctx) error {
	as := inimodulproyek1.GetAllAbsensi(config.Ulbimongoconn, "absensi")
	return c.JSON(as)
}

// GetAllNilai godoc
// @Summary Get All Data Nilai.
// @Description Mengambil semua data nilai.
// @Tags Nilai
// @Accept json
// @Produce json
// @Success 200 {object} Nilai
// @Router /all-nilai [get]
func GetAllNilai(c *fiber.Ctx) error {
	na := inimodulproyek1.GetAllNilai(config.Ulbimongoconn, "nilai")
	return c.JSON(na)
}

// Get From ID Monitoring

// GetMahasiswaFromID godoc
// @Summary Get By ID Data Mahasiswa.
// @Description Ambil per ID data mahasiswa.
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Mahasiswa
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /mahasiswa/{id} [get]
func GetMahasiswaFromID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ms, err := inimodulproyek1.GetMahasiswaFromID(objID, config.Ulbimongoconn, "mahasiswa")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ms)
}

// GetOrangTuaFromID godoc
// @Summary Get By ID Data OrangTua.
// @Description Ambil per ID data orangtua.
// @Tags OrangTua
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} OrangTua
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /orangtua/{id} [get]
func GetOrangTuaFromID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ot, err := inimodulproyek1.GetOrangTuaFromID(objID, config.Ulbimongoconn, "orangtua")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ot)
}

// GetMatakuliahFromID godoc
// @Summary Get By ID Data Matakuliah.
// @Description Ambil per ID data matakuliah.
// @Tags Matakuliah
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Matakuliah
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /matakuliah/{id} [get]
func GetMatakuliahFromID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	mk, err := inimodulproyek1.GetMatakuliahFromID(objID, config.Ulbimongoconn, "matakuliah")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(mk)
}

// GetAbsensiFromID godoc
// @Summary Get By ID Data Absensi.
// @Description Ambil per ID data absensi.
// @Tags Absensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Absensi
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /absensi/{id} [get]
func GetAbsensiFromID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	as, err := inimodulproyek1.GetAbsensiFromID(objID, config.Ulbimongoconn, "absensi")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(as)
}

// GetNilaiFromID godoc
// @Summary Get By ID Data Nilai.
// @Description Ambil per ID data nilai.
// @Tags Nilai
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Nilai
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /nilai/{id} [get]
func GetNilaiFromID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	na, err := inimodulproyek1.GetNilaiFromID(objID, config.Ulbimongoconn, "nilai")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(na)
}

// Insert Function Monitoring

// InsertMahasiswa godoc
// @Summary Insert data mahasiswa.
// @Description Input data mahasiswa.
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Param request body Mahasiswa true "Payload Body [RAW]"
// @Success 200 {object} Mahasiswa
// @Failure 400
// @Failure 500
// @Router /ins-mahasiswa [post]
func InsertMahasiswa(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var mahasiswa inimodelproyek1.Mahasiswa
	if err := c.BodyParser(&mahasiswa); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodulproyek1.InsertMahasiswa(db, "mahasiswa",
		mahasiswa.Nama_mhs,
		mahasiswa.NPM,
		mahasiswa.Jurusan,
		mahasiswa.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// InsertOrangTua godoc
// @Summary Insert data orangtua.
// @Description Input data orangtua.
// @Tags OrangTua
// @Accept json
// @Produce json
// @Param request body OrangTua true "Payload Body [RAW]"
// @Success 200 {object} OrangTua
// @Failure 400
// @Failure 500
// @Router /ins-orangtua [post]
func InsertOrangTua(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var orangtua inimodelproyek1.OrangTua
	if err := c.BodyParser(&orangtua); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodulproyek1.InsertOrangTua(db, "orangtua",
		orangtua.Nama_ortu,
		orangtua.Phone_number,
		orangtua.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// InsertMatakuliah godoc
// @Summary Insert data matakuliah.
// @Description Input data matakuliah.
// @Tags Matakuliah
// @Accept json
// @Produce json
// @Param request body Matakuliah true "Payload Body [RAW]"
// @Success 200 {object} Matakuliah
// @Failure 400
// @Failure 500
// @Router /ins-matakuliah [post]
func InsertMatakuliah(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var matakuliah inimodelproyek1.Matakuliah
	if err := c.BodyParser(&matakuliah); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodulproyek1.InsertMatakuliah(db, "matakuliah",
		matakuliah.Nama_matkul,
		matakuliah.SKS,
		matakuliah.Dosen_pengampu,
		matakuliah.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// InsertAbsensi godoc
// @Summary Insert data absensi.
// @Description Input data absensi.
// @Tags Absensi
// @Accept json
// @Produce json
// @Param request body Absensi true "Payload Body [RAW]"
// @Success 200 {object} Absensi
// @Failure 400
// @Failure 500
// @Router /ins-absensi [post]
func InsertAbsensi(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var absensi inimodelproyek1.Absensi
	if err := c.BodyParser(&absensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodulproyek1.InsertAbsensi(db, "absensi",
		absensi.Nama_mk,
		absensi.Tanggal,
		absensi.Checkin)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// InsertNilai godoc
// @Summary Insert data nilai.
// @Description Input data nilai.
// @Tags Nilai
// @Accept json
// @Produce json
// @Param request body Nilai true "Payload Body [RAW]"
// @Success 200 {object} Nilai
// @Failure 400
// @Failure 500
// @Router /ins-nilai [post]
func InsertNilai(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var nilai inimodelproyek1.Nilai
	if err := c.BodyParser(&nilai); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodulproyek1.InsertNilai(db, "nilai",
		nilai.NPM_ms,
		nilai.Presensi,
		nilai.Nilai_akhir,
		nilai.Grade)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// Update Function Monitoring

// UpdateMahasiswa godoc
// @Summary Update data mahasiswa.
// @Description Ubah data mahasiswa.
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Mahasiswa true "Payload Body [RAW]"
// @Success 200 {object} Mahasiswa
// @Failure 400
// @Failure 500
// @Router /upd-mahasiswa/{id} [put]
func UpdateMahasiswa(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var mahasiswa inimodelproyek1.Mahasiswa
	if err := c.BodyParser(&mahasiswa); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = inimodulproyek1.UpdateMahasiswa(db, "mahasiswa",
		objectID,
		mahasiswa.Nama_mhs,
		mahasiswa.NPM,
		mahasiswa.Jurusan,
		mahasiswa.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// UpdateOrangTua godoc
// @Summary Update data orangtua.
// @Description Ubah data orangtua.
// @Tags OrangTua
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body OrangTua true "Payload Body [RAW]"
// @Success 200 {object} OrangTua
// @Failure 400
// @Failure 500
// @Router /upd-orangtua/{id} [put]
func UpdateOrangTua(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var orangtua inimodelproyek1.OrangTua
	if err := c.BodyParser(&orangtua); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = inimodulproyek1.UpdateOrangTua(db, "orangtua",
		objectID,
		orangtua.Nama_ortu,
		orangtua.Phone_number,
		orangtua.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// UpdateMatakuliah godoc
// @Summary Update data matakuliah.
// @Description Ubah data matakuliah.
// @Tags Matakuliah
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Matakuliah true "Payload Body [RAW]"
// @Success 200 {object} Matakuliah
// @Failure 400
// @Failure 500
// @Router /upd-matakuliah/{id} [put]
func UpdateMatakuliah(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var matakuliah inimodelproyek1.Matakuliah
	if err := c.BodyParser(&matakuliah); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = inimodulproyek1.UpdateMatakuliah(db, "Matakuliah",
		objectID,
		matakuliah.Nama_matkul,
		matakuliah.SKS,
		matakuliah.Dosen_pengampu,
		matakuliah.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// UpdateAbsensi godoc
// @Summary Update data absensi.
// @Description Ubah data absensi.
// @Tags Absensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Absensi true "Payload Body [RAW]"
// @Success 200 {object} Absensi
// @Failure 400
// @Failure 500
// @Router /upd-absensi/{id} [put]
func UpdateAbsensi(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var absensi inimodelproyek1.Absensi
	if err := c.BodyParser(&absensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = inimodulproyek1.UpdateAbsensi(db, "absensi",
		objectID,
		absensi.Nama_mk,
		absensi.Tanggal,
		absensi.Checkin)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// UpdateNilai godoc
// @Summary Update data nilai.
// @Description Ubah data nilai.
// @Tags Nilai
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Nilai true "Payload Body [RAW]"
// @Success 200 {object} Nilai
// @Failure 400
// @Failure 500
// @Router /upd-nilai/{id} [put]
func UpdateNilai(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var nilai inimodelproyek1.Nilai
	if err := c.BodyParser(&nilai); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = inimodulproyek1.UpdateNilai(db, "nilai",
		objectID,
		nilai.NPM_ms,
		nilai.Presensi,
		nilai.Nilai_akhir,
		nilai.Grade)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// Delete Function Monitoring

// Delete Function

// DeleteMahasiswaByID godoc
// @Summary Delete data mahasiswa.
// @Description Hapus data mahasiswa.
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete-mahasiswa/{id} [delete]
func DeleteMahasiswaByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = inimodulproyek1.DeleteMahasiswaByID(objID, config.Ulbimongoconn, "mahasiswa")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}

// DeleteOrangTuaByID godoc
// @Summary Delete data orangtua.
// @Description Hapus data orangtua.
// @Tags OrangTua
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete-orangtua/{id} [delete]
func DeleteOrangTuaByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = inimodulproyek1.DeleteOrangTuaByID(objID, config.Ulbimongoconn, "orangtua")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}

// DeleteMatakuliahByID godoc
// @Summary Delete data matakuliah.
// @Description Hapus data matakuliah.
// @Tags Matakuliah
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete-matakuliah/{id} [delete]
func DeleteMatakuliahByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = inimodulproyek1.DeleteMatakuliahByID(objID, config.Ulbimongoconn, "matakuliah")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}

// DeleteAbsensiByID godoc
// @Summary Delete data absensi.
// @Description Hapus data absensi.
// @Tags Absensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete-absensi/{id} [delete]
func DeleteAbsensiByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = inimodulproyek1.DeleteAbsensiByID(objID, config.Ulbimongoconn, "absensi")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}

// DeleteNilaiByID godoc
// @Summary Delete data nilai.
// @Description Hapus data nilai.
// @Tags Nilai
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete-nilai/{id} [delete]
func DeleteNilaiByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = inimodulproyek1.DeleteNilaiByID(objID, config.Ulbimongoconn, "nilai")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}
