package types

type EducatorEvents struct {
	EducatorMasterId                 int64
	EducatorDisplayText              string
	EducatorLongDisplayText          string
	PreviousWeekMonday               JsonDate
	NextWeekMonday                   JsonDate
	IsPreviousWeekReferenceAvailable bool
	IsNextWeekReferenceAvailable     bool
	IsCurrentWeekReferenceAvailable  bool
	WeekDisplayText                  string
	WeekMonday                       JsonDate
	//EducatorEventsDays (Array[SpbuEducation.TimeTable.Web.Api.v1.DataContracts.EducatorEventsContract.EventsDay], optional): Events grouped by days of week
}
