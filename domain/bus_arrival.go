package domain

type BusResponse[T any] struct {
	Root T `json:"bustime-response"`
}

type BusRoutesRoot struct {
	Routes []BusRoute `json:"routes"`
}

type BusRoute struct {
	RouteId    string `json:"rt"`
	RouteName  string `json:"rtnm"`
	RouteRun   string `json:"rtdd"`
	RouteColor string `json:"rtclr"`
}

type BusArrivalsRoot struct {
	Etas []BusEta `json:"prd"`
}

type BusEta struct {
	PredictedAt       BusTimestamp `json:"tmstmp"`
	Typ               string       `json:"typ"`
	StopName          string       `json:"stpnm"`
	StopId            string       `json:"stpid"`
	VehicleId         string       `json:"vid"`
	Distance          int          `json:"dstp"`
	RouteId           string       `json:"rt"`
	RouteDesc         string       `json:"rtdd"`
	RouteDir          string       `json:"rtdir"`
	DestName          string       `json:"des"`
	ArrivalTime       BusTimestamp `json:"prdtm"`
	ArrivalPrediction string       `json:"prdctdn"`
	Zone              string       `json:"zone"`
	PassengerLoad     string       `json:"psgld"`
}

type BusGroupKey struct {
	RouteId, RouteDir, DestName string
}
