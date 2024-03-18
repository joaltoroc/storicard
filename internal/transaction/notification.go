package transaction

type Notification interface {
	SendMail(destination string, fileName string, executionID string, body []map[string]string) error
}
