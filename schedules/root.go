package schedules

import (
	"bufio"
	"os"
	"path"
	"smanti_schedules/teachers"
	"strconv"
	"strings"
)

func GetSchedules(schedules_path string, class string, teachers *teachers.TeacherMaps) (*ScheduleMaps, error) {
	maps := ScheduleMaps{}

	dirs, err := os.ReadDir(path.Join(schedules_path, class))
	if err != nil {
		return &maps, err
	}

	for _, entry := range dirs {
		if !entry.Type().IsDir() {
			file, err := os.OpenFile(path.Join(schedules_path, class, entry.Name()), os.O_RDONLY, os.ModeDevice)
			if err == nil {
				scanner := bufio.NewScanner(file)
				scanner.Split(bufio.ScanLines)
				index := 0

				for scanner.Scan() {
					texts := strings.Split(scanner.Text(), "|")
					clockTexts := strings.Split(strings.TrimSpace(texts[0]), "-")

					code, _ := strconv.Atoi(strings.TrimSpace(texts[1]))
					t := (*teachers)[strconv.Itoa(code)]
					data := Schedule{
						Index:       index + 1,
						TeacherCode: code,
						ClockBegin:  clockTexts[0],
						ClockEnd:    clockTexts[1],
						Teacher:     &t,
					}

					if len(t.Nama) < 1 {
						data.Teacher = nil
					}

					maps[strings.Replace(entry.Name(), ".txt", "", -1)] = append(maps[strings.Replace(entry.Name(), ".txt", "", -1)], data)
					index++
				}
			}
		}
	}

	return &maps, nil
}
