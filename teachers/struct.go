package teachers

type Teacher struct {
	Nama        string   `json:"nama"`
	Mapel       []string `json:"mapel"`
	JamMengajar int      `json:"jam_mengajar"`
	Keterangan  string   `json:"keterangan"`
}

type TeacherMaps map[string]Teacher
