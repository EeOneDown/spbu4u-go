package yandex_rasp_api

type SearchResponse struct {
	Pagination       Pagination `json:"pagination"`
	IntervalSegments []Segment  `json:"interval_segments"`
	Segments         []Segment  `json:"segments"`
	Search           Search     `json:"search"`
}

type Pagination struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Segment struct {
	Arrival           string     `json:"arrival"`
	From              Station    `json:"from"`
	Thread            Thread     `json:"thread"`
	DeparturePlatform string     `json:"departure_platform"`
	Departure         string     `json:"departure"`
	Stops             string     `json:"stops"`
	DepartureTerminal string     `json:"departure_terminal"`
	To                Station    `json:"to"`
	HasTransfers      bool       `json:"has_transfers"`
	TicketsInfo       TicketInfo `json:"tickets_info"`
	Duration          int        `json:"duration"`
	ArrivalTerminal   string     `json:"arrival_terminal"`
	StartDate         string     `json:"start_date"`
	ArrivalPlatform   string     `json:"arrival_platform"`
}

type Station struct {
	Code            string `json:"code"`
	Title           string `json:"title"`
	StationType     string `json:"station_type"`
	StationTypeName string `json:"station_type_name"`
	PopularTitle    string `json:"popular_title"`
	ShortTitle      string `json:"short_title"`
	TransportType   string `json:"transport_type"`
	Type            string `json:"type"`
}

type Thread struct {
	UID              string           `json:"uid"`
	Title            string           `json:"title"`
	Interval         Interval         `json:"interval"`
	Number           string           `json:"number"`
	ShortTitle       string           `json:"short_title"`
	ThreadMethodLink string           `json:"thread_method_link"`
	Carrier          Carrier          `json:"carrier"`
	TransportType    string           `json:"transport_type"`
	Vehicle          string           `json:"vehicle"`
	TransportSubtype TransportSubtype `json:"transport_subtype"`
	ExpressType      string           `json:"express_type"`
}

type Interval struct {
	Density   string `json:"density"`
	EndTime   string `json:"end_time"`
	BeginTime string `json:"begin_time"`
}

type TicketInfo struct {
	EtMarker bool    `json:"et_marker"`
	Places   []Place `json:"places"`
}

type Place struct {
	Currency string `json:"currency"`
	Price    Price  `json:"price"`
	Name     string `json:"name"`
}

type Price struct {
	Cents int `json:"cents"`
	Whole int `json:"whole"`
}

type Carrier struct {
	Code     int    `json:"code"`
	Contacts string `json:"contacts"`
	URL      string `json:"url"`
	LogoSVG  string `json:"logo_svg"`
	Title    string `json:"title"`
	Phone    string `json:"phone"`
	Codes    Codes  `json:"codes"`
	Address  string `json:"address"`
	Logo     string `json:"logo"`
	Email    string `json:"email"`
}

type TransportSubtype struct {
	Color string `json:"color"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

type Codes struct {
	ICAO   string `json:"icao"`
	Sirena string `json:"sirena"`
	IATA   string `json:"iata"`
}

type Search struct {
	Date string  `json:"date"`
	To   Station `json:"to"`
	From Station `json:"from"`
}
