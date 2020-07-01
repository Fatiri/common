package validate

import (
	"github.com/microservices/serives/account/models"
	"github.com/sirupsen/logrus"
)

//AccountValueValidate validate value data account
func AccountValueValidate(data *models.Account) {
	if data.Fullname == "" {
		logrus.Error("fullname harus di isi")
		return
	}
	if data.Place == "" {
		logrus.Error("place harus di isi")
		return
	}
}