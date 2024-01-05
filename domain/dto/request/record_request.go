package request

type GetRecordRequest struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int64  `json:"minCount"`
	MaxCount  int64  `json:"maxCount"`
}
