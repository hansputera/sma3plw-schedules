package teachers

import (
	"encoding/json"
	"io"
	"os"
	"path"
)

func GetTeachers(teachers_path string) (TeacherMaps, error) {
	teachers := TeacherMaps{}

	file, err := os.OpenFile(path.Join(teachers_path, "raw.json"), os.O_RDONLY, os.ModeDevice)
	if err != nil {
		return teachers, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return teachers, err
	}

	if err = json.Unmarshal(data, &teachers); err != nil {
		return teachers, err
	}

	return teachers, nil
}