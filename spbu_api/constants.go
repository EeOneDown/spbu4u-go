package spbu_api

const (
	BaseUrl = "https://timetable.spbu.ru/api/v1"
	// Educators
	//EducatorSemesterUrl = BaseUrl + "/educators/%d/events"
	EducatorCustomUrl = BaseUrl + "/educators/%d/events/%d-%v-%d/%d-%v-%d"
	//EducatorSearchUrl = BaseUrl + "educators/search/%v"
	// Groups
	//GroupWeekUrl = BaseUrl + "/groups/%d/events"
	//GroupCustomWeekUrl = BaseUrl + "/groups/%d/events/%d-%v-%d"
	GroupCustomUrl = BaseUrl + "/groups/%d/events/%d-%v-%d/%d-%v-%d"
)
