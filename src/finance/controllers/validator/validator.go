package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/riyan-eng/api-finance/src/finance/controllers/dto"
)

var validate = validator.New()

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func CashReceipt(cashReceipt dto.CashReceiptReq) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(cashReceipt)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func CashPayment(cashPayment dto.CashPaymentReq) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(cashPayment)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
