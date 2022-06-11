package models

type Mahasiswa struct {
	Nama  string `json:"nama"`
	NIM   string `json:"nim"`
	Kelas string `json:"kelas"`
}

type ArrMhs struct {
	ArrMahasiswa []*Mahasiswa
}

type NotifSlackParam struct {
	Text string
	Data []*Mahasiswa
}
