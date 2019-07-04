package messages

import "time"

//AddDailyRateRequest : request message to add daily exchange
type AddDailyRateRequest struct {
	From         string `json:"from" binding:"required,len=3"`
	To           string `json:"to" binding:"required,len=3"`
	ExchangeDate string `json:"date" binding:"required"`
	Rate         string `json:"rate" binding:"required"`
}

//DeleteRequest : request message to delete list of exchanges
type DeleteRequest struct {
	Exchanges []*ExchangeRequest `json:"exchanges" binding:"required"`
}

//Exchange7DaysResponse :request model to show 7 days exchange
type Exchange7DaysResponse struct {
	From     string         `json:"from" binding:"required,len=3"`
	To       string         `json:"to" binding:"required,len=3"`
	Average  float64        `json:"average"`
	Variance float64        `json:"variance"`
	Data     []ExchangeData `json:"data"`
}

//ExchangeData : model to show date and rate
type ExchangeData struct {
	Date *time.Time `json:"Date"`
	Rate float64    `json:"Rate"`
}

//ExchangeRequest :request model for add exchange
type ExchangeRequest struct {
	From string `json:"from" binding:"required,len=3"`
	To   string `json:"to" binding:"required,len=3"`
}

//ExchangeRequest7Day :request model for add exchange
type ExchangeRequest7Day struct {
	From string `json:"from" binding:"required,len=3"`
	To   string `json:"to" binding:"required,len=3"`
	Date string `json:"date" binding:"required"`
}

type TrackedRequest struct {
	Date      string             `json:"date" binding:"required"`
	Exchanges []*ExchangeRequest `json:"exchanges" binding:"required"`
}

//TrackedResponse : Response struct for tracked request
type TrackedResponse struct {
	From string `json:"from" binding:"required,len=3"`
	To   string `json:"to" binding:"required,len=3"`
	Rate string `json:"rate" binding:"required"`
	Avg  string `json:"7-day-avg" binding:"required"`
}
