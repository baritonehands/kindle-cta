package domain

type TrainArrivalsResponse struct {
	Root TrainArrivalsRoot `json:"ctatt"`
}

type TrainArrivalsRoot struct {
	Timestamp TrainTimestamp `json:"tmst"`
	ErrorCode string         `json:"errCd"`
	ErrorName string         `json:"errNm"`
	Etas      []TrainEta     `json:"eta"`
}

type TrainEta struct {
	StationId     string         `json:"staId"`
	StopId        string         `json:"stpId"`
	StationName   string         `json:"staNm"`
	StopDesc      string         `json:"stpDe"`
	Run           string         `json:"rn"`
	Route         string         `json:"rt"`
	DestStop      string         `json:"destSt"`
	DestName      string         `json:"destNm"`
	DirCode       string         `json:"trDr"`
	PredictedAt   TrainTimestamp `json:"prdt"`
	ArrivalTime   TrainTimestamp `json:"arrT"`
	IsApproaching string         `json:"isApp"`
	IsSchedule    string         `json:"isSch"`
	IsDelayed     string         `json:"isDly"`
	IsFault       string         `json:"isFlt"`
	Lat           string         `json:"lat"`
	Lon           string         `json:"lon"`
	Heading       string         `json:"heading"`
}

type TrainGroupKey struct {
	Route, DestName string
}
