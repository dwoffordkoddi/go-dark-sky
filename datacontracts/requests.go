package datacontracts

type CompletedJobResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  []string    `json:"errors"`
}
