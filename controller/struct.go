package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Surat struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	No_surat     int                `bson:"no_surat,omitempty" json:"no_surat,omitempty"`
	Status_surat Status             `bson:"status_surat,omitempty" json:"status_surat,omitempty"`
	Perihal      string             `bson:"perihal,omitempty" json:"perihal,omitempty"`
	Id_pos       Kodepos            `bson:"id_pos,omitempty" json:"id_pos,omitempty"`
	Pengirim_srt Pengirim           `bson:"pengirim_srt,omitempty" json:"pengirim_srt,omitempty"`
	Penerima_srt Penerima           `bson:"penerima_srt,omitempty" json:"penerima_srt,omitempty"`
}

type Disposisi struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Kode_disposisi int                `bson:"kode_disposisi,omitempty" json:"kode_disposisi,omitempty"`
	Tgl_disposisi  string             `bson:"tgl_disposisi,omitempty" json:"tgl_disposisi,omitempty"`
	Penerima_surat Penerima           `bson:"penerima_surat,omitempty" json:"penerima_surat,omitempty"`
	Stat_disposisi Status             `bson:"status_disposisi,omitempty" json:"status_disposisi,omitempty"`
}
type Kodepos struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Kode_daerah int                `bson:"kode_daerah,omitempty" json:"kode_daerah,omitempty"`
	Nama_daerah string             `bson:"nama_daerah,omitempty" json:"nama_dareah,omitempty"`
}

type Status struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Id_status  int                `bson:"id_status,omitempty" json:"id_status,omitempty"`
	Keterangan string             `bson:"keterangan,omitempty" json:"keterangan,omitempty"`
}

type Pengirim struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_pengirim string             `bson:"nama_pengirim,omitempty" json:"nama_pengirim,omitempty"`
	Alamat        string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Tgl_kirim     string             `bson:"tgl_kirim,omitempty" json:"tgl_kirim,omitempty"`
}

type Penerima struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_penerima string             `bson:"nama_penerima,omitempty" json:"nama_penerima,omitempty"`
	Alamat        string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Tgl_terima    string             `bson:"tgl_terima,omitempty" json:"tgl_terima,omitempty"`
}
type Karyawan struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama        string             `bson:"nama,omitempty" json:"nama,omitempty"`
	PhoneNumber string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Jabatan     string             `bson:"jabatan,omitempty" json:"jabatan,omitempty"`
	Jam_kerja   []JamKerja         `bson:"jam_kerja,omitempty" json:"jam_kerja,omitempty"`
	Hari_kerja  []string           `bson:"hari_kerja,omitempty" json:"hari_kerja,omitempty"`
}

type JamKerja struct {
	Durasi     int      `bson:"durasi,omitempty" json:"durasi,omitempty"`
	Jam_masuk  string   `bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty"`
	Jam_keluar string   `bson:"jam_keluar,omitempty" json:"jam_keluar,omitempty"`
	Gmt        int      `bson:"gmt,omitempty" json:"gmt,omitempty"`
	Hari       []string `bson:"hari,omitempty" json:"hari,omitempty"`
	Shift      int      `bson:"shift,omitempty" json:"shift,omitempty"`
	Piket_tim  string   `bson:"piket_tim,omitempty" json:"piket_tim,omitempty"`
}

type Presensi struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Longitude    float64            `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Latitude     float64            `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Location     string             `bson:"location,omitempty" json:"location,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	// Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Checkin string   `bson:"checkin,omitempty" json:"checkin,omitempty"`
	Biodata Karyawan `bson:"biodata,omitempty" json:"biodata,omitempty"`
}

type Lokasi struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama     string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Batas    Geometry           `bson:"batas,omitempty" json:"batas,omitempty"`
	Kategori string             `bson:"kategori,omitempty" json:"kategori,omitempty"`
}

type Geometry struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
}