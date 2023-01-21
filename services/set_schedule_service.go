package services

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
)

type ScheduleSetPayload struct {
	Begin string `json:"begin"`
	End   string `json:"end"`
	Code  int    `json:"code"`
}

func valid_day(day string) bool {
	return day == "senin" || day == "selasa" || day == "rabu" || day == "kamis" || day == "jumat"
}

func SetSchedules(schedules_path string, class string, day string, schedules []*ScheduleSetPayload) error {
	day = strings.ToLower(day)

	if !valid_day(day) {
		return errors.New("unsupported day")
	}
	file, err := os.OpenFile(path.Join(schedules_path, class, fmt.Sprintf("%s.txt", day)), os.O_RDWR, os.ModeDevice)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			file, err = os.Create(path.Join(schedules_path, class, fmt.Sprintf("%s.txt", day)))
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	for _, s := range schedules {
		file.WriteString(fmt.Sprintf("%s - %s | %d\n", s.Begin, s.End, s.Code))
	}

	file.Close()

	return err
}
