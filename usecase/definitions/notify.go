package definitions

import "go-csv/domain/models"

type Notify interface {
	DoNotifySlackSend(param models.NotifSlackParam) error
}
