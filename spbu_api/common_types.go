package spbu_api

type EventsDay struct {
	Day            ZLessTime
	DayString      string
	DayStudyEvents []Event
}

type Event struct {
	StudyEventsTimeTableKindCode     int
	Start                            ZLessTime
	End                              ZLessTime
	Subject                          string
	TimeIntervalString               string
	DateWithTimeIntervalString       string
	DisplayDateAndTimeIntervalString string
	LocationsDisplayText             string
	EducatorsDisplayText             string
	HasEducators                     bool
	IsCancelled                      bool
	ContingentUnitName               string
	DivisionAndCourse                string
	IsAssigned                       bool
	TimeWasChanged                   bool
	LocationsWereChanged             bool
	EducatorsWereReassigned          bool
	ElectiveDisciplinesCount         int
	IsElective                       bool
	HasTheSameTimeAsPreviousItem     bool
	IsStudy                          bool
	AllDay                           bool
	WithinTheSameDay                 bool
	EventLocations                   []EventLocation
	EducatorIds                      []EducatorId
}

type EventLocation struct {
	IsEmpty                  bool
	DisplayName              string
	HasGeographicCoordinates bool
	Latitude                 float32
	Longitude                float32
	LatitudeValue            string
	LongitudeValue           string
	EducatorsDisplayText     string
	HasEducators             bool
	EducatorIds              []EducatorId
}

type EducatorId struct {
	Id   int64  `json:"Item1"`
	Name string `json:"Item2"`
}
