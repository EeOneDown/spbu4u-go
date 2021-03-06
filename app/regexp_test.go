package app

import (
	"testing"
)

func isEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestFindScheduleFromGroupUrl(t *testing.T) {
	t.Parallel()
	for _, url := range []string{
		"https://timetable.spbu.ru/BIOL/StudentGroupEvents/Online/247986",
		"timetable.spbu.ru/BIOL/StudentGroupEvents/Attestation/247986",
		"https://timetable.spbu.ru/BIOL/StudentGroupEvents/Final/247986",
		"https://timetable.spbu.ru/BIOL/StudentGroupEvents/Semester/247986/1",
		"timetable.spbu.ru/BIOL/StudentGroupEvents/Primary/247986",
		"timetable.spbu.ru/BIOL/StudentGroupEvents/Primary/247986/2020-02-17",
		"https://timetable.spbu.ru/StudentGroupEvents/Semester?publicDivisionAlias=BIOL&studentGroupId=247986",
	} {
		if !isEq(
			RegExpRegisterUrl.FindStringSubmatch(url),
			[]string{url, "StudentGroupEvents", "247986"},
		) {
			t.Fail()
			continue
		}
	}
}

func TestFindScheduleFromEducatorUrl(t *testing.T) {
	t.Parallel()
	for _, url := range []string{
		"timetable.spbu.ru/EducatorEvents/1420",
		"https://timetable.spbu.ru/EducatorEvents/1420/1",
	} {
		if !isEq(
			RegExpRegisterUrl.FindStringSubmatch(url),
			[]string{url, "EducatorEvents", "1420"},
		) {
			t.Fail()
			continue
		}
	}
}

func TestFindScheduleFromWeekEducatorUrl(t *testing.T) {
	t.Parallel()
	for _, url := range []string{
		"timetable.spbu.ru/WeekEducatorEvents/1420",
		"https://timetable.spbu.ru/WeekEducatorEvents/1420/2020-02-17",
	} {
		if !isEq(
			RegExpRegisterUrl.FindStringSubmatch(url),
			[]string{url, "WeekEducatorEvents", "1420"},
		) {
			t.Fail()
			continue
		}
	}
}

func TestRegExpSubjectNoComma(t *testing.T) {
	t.Parallel()
	message := "Вычислительные методы в гидродинамике и теории волн"
	if !isEq(
		RegExpSubject.FindStringSubmatch(message),
		[]string{message, "Вычислительные методы в гидродинамике и теории волн", ""},
	) {
		t.Fail()
	}
}

func TestRegExpSubjectOneComma(t *testing.T) {
	t.Parallel()
	message := "Вычислительные методы в гидродинамике и теории волн, практическое занятие в присутствии преподавателя"
	if !isEq(
		RegExpSubject.FindStringSubmatch(message),
		[]string{
			message,
			"Вычислительные методы в гидродинамике и теории волн",
			"практическое занятие в присутствии преподавателя",
		},
	) {
		t.Fail()
	}
}

func TestRegExpSubjectSeveralCommas(t *testing.T) {
	t.Parallel()
	message := "Вычислительные методы в гидродинамике и теории волн, лекция, лекция, лекция"
	if !isEq(
		RegExpSubject.FindStringSubmatch(message),
		[]string{message, "Вычислительные методы в гидродинамике и теории волн, лекция, лекция", "лекция"},
	) {
		t.Fail()
	}
}

func TestRegExpStart(t *testing.T) {
	t.Parallel()
	for _, message := range []string{"/start", "релогин"} {
		if !isEq(
			RegExpStart.FindStringSubmatch(message),
			[]string{message},
		) {
			t.Fail()
			continue
		}
	}
}

func TestRegExpStartBad(t *testing.T) {
	t.Parallel()
	for _, message := range []string{"asdasdads/startasdasd", "тестыы"} {
		if RegExpStart.FindStringSubmatch(message) != nil {
			t.Fail()
			continue
		}
	}
}

func TestRegExpWhoAmI(t *testing.T) {
	t.Parallel()
	for _, message := range []string{"/me", "/whoami", "я", "кто я", "группа"} {
		if !isEq(
			RegExpWhoAmI.FindStringSubmatch(message),
			[]string{message},
		) {
			t.Fail()
			continue
		}
	}
}

func TestRegExpWhoAmIBad(t *testing.T) {
	t.Parallel()
	for _, message := range []string{"/measdasd", "asd/whoamiaasd", "сессияя", "кто группа", "группа Б"} {
		if RegExpWhoAmI.FindStringSubmatch(message) != nil {
			t.Fail()
			continue
		}
	}
}

func TestRegExpMainMenu(t *testing.T) {
	t.Parallel()
	for _, message := range []string{"/menu", EmojiBack, "назад"} {
		if !isEq(
			RegExpMainMenu.FindStringSubmatch(message),
			[]string{message},
		) {
			t.Fail()
			continue
		}
	}
}

func TestRegExpMainMenuBad(t *testing.T) {
	t.Parallel()
	for _, message := range []string{"/menuxxx", EmojiAlarmClock, "дазан"} {
		if RegExpMainMenu.FindStringSubmatch(message) != nil {
			t.Fail()
			continue
		}
	}
}

func TestRegExpSchedule(t *testing.T) {
	t.Parallel()
	for _, message := range []string{"/schedule", "расписание"} {
		if !isEq(
			RegExpSchedule.FindStringSubmatch(message),
			[]string{message},
		) {
			t.Fail()
			continue
		}
	}
}

func TestRegExpScheduleBad(t *testing.T) {
	t.Parallel()
	for _, message := range []string{"/schedule  zzz", "расписание сессии"} {
		if RegExpSchedule.FindStringSubmatch(message) != nil {
			t.Fail()
			continue
		}
	}
}

func TestRegExpToday(t *testing.T) {
	t.Parallel()
	for _, message := range []string{"/today", "сегодня"} {
		if !isEq(
			RegExpToday.FindStringSubmatch(message),
			[]string{message},
		) {
			t.Fail()
			continue
		}
	}
}

func TestRegExpTodayBad(t *testing.T) {
	t.Parallel()
	for _, message := range []string{"/today  zzz", "сегодня спим"} {
		if RegExpToday.FindStringSubmatch(message) != nil {
			t.Fail()
			continue
		}
	}
}

func TestRegExpTomorrow(t *testing.T) {
	t.Parallel()
	for _, message := range []string{"/week", "вся неделя"} {
		if !isEq(
			RegExpWeek.FindStringSubmatch(message),
			[]string{message},
		) {
			t.Fail()
			continue
		}
	}
}

func TestRegExpTomorrowBad(t *testing.T) {
	t.Parallel()
	for _, message := range []string{"/weeвся неделя", "/week неделя", "вся неделя след"} {
		if RegExpWeek.FindStringSubmatch(message) != nil {
			t.Fail()
			continue
		}
	}
}

func TestRegExpWeekNext(t *testing.T) {
	t.Parallel()
	for _, message := range []string{"/weeknext", "вся неделя след", "вся неделя следующая"} {
		if !isEq(
			RegExpWeekNext.FindStringSubmatch(message),
			[]string{message},
		) {
			t.Fail()
			continue
		}
	}
}

func TestRegExpWeekNextBad(t *testing.T) {
	t.Parallel()
	for _, message := range []string{"/week", "вся неделя следxxx"} {
		if RegExpWeekNext.FindStringSubmatch(message) != nil {
			t.Fail()
			continue
		}
	}
}

func TestRegExpSettings(t *testing.T) {
	t.Parallel()
	for _, message := range []string{"/settings", "настройки", EmojiGear} {
		if !isEq(
			RegExpSettings.FindStringSubmatch(message),
			[]string{message},
		) {
			t.Fail()
			continue
		}
	}
}

func TestRegExpSettingsBad(t *testing.T) {
	t.Parallel()
	for _, message := range []string{"not /settings", "настройки2", EmojiBookmark} {
		if RegExpSettings.FindStringSubmatch(message) != nil {
			t.Fail()
			continue
		}
	}
}

func TestRegExpExit(t *testing.T) {
	t.Parallel()
	for _, message := range []string{"/exit", "завершить", "выход"} {
		if !isEq(
			RegExpExit.FindStringSubmatch(message),
			[]string{message},
		) {
			t.Fail()
			continue
		}
	}
}

func TestRegExpExitBad(t *testing.T) {
	t.Parallel()
	for _, message := range []string{"not /exit", "завершить все", "вход"} {
		if RegExpExit.FindStringSubmatch(message) != nil {
			t.Fail()
			continue
		}
	}
}

func TestRegExpSupport(t *testing.T) {
	t.Parallel()
	for _, message := range []string{"/support", "поддержка"} {
		if !isEq(
			RegExpSupport.FindStringSubmatch(message),
			[]string{message},
		) {
			t.Fail()
			continue
		}
	}
}

func TestRegExpSupportBad(t *testing.T) {
	t.Parallel()
	for _, message := range []string{"/support2", "поддержка с воздуха"} {
		if RegExpSupport.FindStringSubmatch(message) != nil {
			t.Fail()
			continue
		}
	}
}

func TestRegExpTrains(t *testing.T) {
	t.Parallel()
	for _, message := range []string{"/trains", EmojiStation, "элекрички", "элекрон"} {
		if !isEq(
			RegExpTrains.FindStringSubmatch(message),
			[]string{message},
		) {
			t.Fail()
			continue
		}
	}
}

func TestRegExpTrainsBad(t *testing.T) {
	t.Parallel()
	for _, message := range []string{"/trans", EmojiStar, "лекрички", "протон"} {
		if RegExpTrains.FindStringSubmatch(message) != nil {
			t.Fail()
			continue
		}
	}
}
