package rabbithole

import "fmt"

type ErrorResponse struct {
	StatusCode int
	Message    string `json:"error"`
	Reason     string `json:"reason"`
}

func (rme ErrorResponse) Error() string {
	return fmt.Sprintf("Error %d (%s): %s", rme.StatusCode, rme.Message, rme.Reason)
}
