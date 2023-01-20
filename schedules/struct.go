package schedules

type Schedule struct {
	Index int `json:"urutan"`
	TeacherCode int `json:"code"`
	ClockBegin string `json:"beginAt"`
	ClockEnd string `json:"endAt"`
}

type ScheduleMaps map[string]Schedule