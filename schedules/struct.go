package schedules

import "smanti_schedules/teachers"

type Schedule struct {
	Index int `json:"urutan"`
	TeacherCode int `json:"code"`
	ClockBegin string `json:"beginAt"`
	ClockEnd string `json:"endAt"`

	Teacher *teachers.Teacher `json:"teacher"`
}

type ScheduleMaps map[string][]Schedule