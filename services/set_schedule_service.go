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

func SetSchedules(schedules_path string, class string, day string, schedules []*ScheduleSetPayload) error {
	day = strings.ToLower(day)
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
