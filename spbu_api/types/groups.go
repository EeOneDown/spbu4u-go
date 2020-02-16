package types

type GroupEvents struct {
	StudentGroupId                   int64
	StudentGroupDisplayName          string
	TimeTableDisplayName             string
	PreviousWeekMonday               JsonDate
	NextWeekMonday                   JsonDate
	IsPreviousWeekReferenceAvailable bool
	IsNextWeekReferenceAvailable     bool
	IsCurrentWeekReferenceAvailable  bool
	WeekDisplayText                  string
	WeekMonday                       JsonDate
	//Days (Array[SpbuEducation.TimeTable.Web.Api.v1.DataContracts.GroupEventsContract.EventsDay], optional): Events grouped by days
}
