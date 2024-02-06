package validator

import (
	"fmt"

	"github.com/albugowy15/synapsis-backend-test/internal/pkg/models"
)

func ValidateAddToCartRequest(body models.AddCartRequest) error {
	if body.Quantity <= 0 {
		return fmt.Errorf("quantity cannot be 0")
	}
	return nil
}
