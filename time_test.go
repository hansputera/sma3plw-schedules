package main

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	days := []string{"senin", "selasa", "kamis", "jumat"}
	// senin = 9
	// selasa, dan kamis = 10
	// jumat = 5

	for _, day := range days {
		times := []string{}
		max_times := 0

		if day == "jumat" {
			max_times = 7
			times = append(times, "8.30")
		} else {
			max_times = 13
			if day != "senin" {
				times = append(times, "7.15")
			} else {
				max_times = 12
				times = append(times, "08.00")
			}
		}

		timeParse, err := time.Parse("15:04", strings.Replace(times[0], ".", ":", -1))
		if err != nil {
			t.Error(err)
		} else {

			for {
				if len(times) >= max_times {
					break
				}

				end_time := timeParse

				if day == "senin" && len(times) == 4 {
					end_time = end_time.Add(25 * time.Minute)
				} else if day == "jumat" {
					if len(times) == 4 {
						end_time = end_time.Add(20 * time.Minute)
					} else {
						end_time = end_time.Add(35 * time.Minute)
					}
				} else {
					if len(times) == 7 {
						end_time = end_time.Add(50 * time.Minute)
					} else {
						end_time = end_time.Add(35 * time.Minute)
					}
				}

				times = append(times, fmt.Sprintf("%d.%d - %d.%d", timeParse.Hour(), timeParse.Minute(), end_time.Hour(), end_time.Minute()))
				timeParse = end_time
			}

			times = times[1:]
			fmt.Println(times)
		}
	}
}
