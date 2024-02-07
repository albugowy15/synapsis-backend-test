package repositories

import (
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/db"
	"github.com/albugowy15/synapsis-backend-test/internal/pkg/models"
)

type PaymentMethodRepository struct{}

var patmentMethodRepository *PaymentMethodRepository

func GetPaymentMethodRepository() *PaymentMethodRepository {
	if patmentMethodRepository == nil {
		patmentMethodRepository = &PaymentMethodRepository{}
	}
	return patmentMethodRepository
}

func (r *PaymentMethodRepository) GetById(paymentMethodId int) (models.PaymentMethod, error) {
	var paymentMethod models.PaymentMethod
	err := db.GetDB().QueryRow("SELECT id, type, merchant, fee FROM payment_methods WHERE id = $1", paymentMethodId).
		Scan(&paymentMethod.ID, &paymentMethod.Type, &paymentMethod.Merchant, &paymentMethod.Fee)
	if err != nil {
		return paymentMethod, err
	}
	return paymentMethod, nil
}
