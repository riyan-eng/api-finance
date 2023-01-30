package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/riyan-eng/api-finance/constant"
	"github.com/riyan-eng/api-finance/module/finance/controllers/dto"
	"github.com/riyan-eng/api-finance/module/finance/services"
	"github.com/riyan-eng/api-finance/module/finance/services/entities"
	"github.com/riyan-eng/api-finance/util"
)

type JournalController interface {
	CashReceipt(c *fiber.Ctx) error
	CashPayment(c *fiber.Ctx) error
	Sales(c *fiber.Ctx) error
	Purchase(c *fiber.Ctx) error
	General(c *fiber.Ctx) error
}

type journalService struct {
	JournalService services.JournalService
}

func NewJournalController(jS services.JournalService) JournalController {
	return &journalService{
		JournalService: jS,
	}
}

func (service *journalService) CashReceipt(c *fiber.Ctx) error {
	cashReceiptBody := new(dto.CashReceiptReq)

	// parsing body json
	if err := c.BodyParser(&cashReceiptBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	// validate body json
	if err := util.Validate(*cashReceiptBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err,
			"message": "fail",
		})
	}

	// convert dto to entity
	journal := entities.JournalEntity{
		ID:          uuid.NewString(),
		UserID:      "58dd4ecc-8cde-4ca3-b4f0-7451b7b59ce8",
		Amount:      cashReceiptBody.Amount,
		Description: cashReceiptBody.Description,
		Position: entities.Position{
			Debet: entities.Transaction{
				Code:   constant.COA_CASH,
				Amount: cashReceiptBody.Amount,
			},
			Credit: entities.Transaction{
				Code:   cashReceiptBody.Code,
				Amount: cashReceiptBody.Amount,
			},
		},
	}

	// comunicate with service
	if err := service.JournalService.CashReceipt(journal); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "success inserted cash receipt",
		"message": "ok",
	})
}

func (service *journalService) CashPayment(c *fiber.Ctx) error {
	cashPaymentBody := new(dto.CashPaymentReq)

	// parsing body json
	if err := c.BodyParser(&cashPaymentBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	// validate  body json
	if err := util.Validate(*cashPaymentBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err,
			"message": "fail",
		})
	}

	// convert dto to entity
	journal := entities.JournalEntity{
		ID:          uuid.NewString(),
		UserID:      "58dd4ecc-8cde-4ca3-b4f0-7451b7b59ce8",
		Amount:      cashPaymentBody.Amount,
		Description: cashPaymentBody.Description,
		Position: entities.Position{
			Debet: entities.Transaction{
				Code:   cashPaymentBody.Code,
				Amount: cashPaymentBody.Amount,
			},
			Credit: entities.Transaction{
				Code:   constant.COA_CASH,
				Amount: cashPaymentBody.Amount,
			},
		},
	}

	// communicate with service
	if err := service.JournalService.CashPayment(journal); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "success inserted cash payment",
		"message": "ok",
	})
}

func (jS *journalService) Sales(c *fiber.Ctx) error {
	salesBody := new(dto.SalesReq)

	// parse body json
	if err := c.BodyParser(&salesBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	// validate body json
	if err := util.Validate(*salesBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err,
			"message": "fail",
		})
	}

	// fmt.Println(salesBody)

	// convert dto to entity
	// journal := entities.JournalEntity{
	// 	ID:          uuid.NewString(),
	// 	UserID:      "58dd4ecc-8cde-4ca3-b4f0-7451b7b59ce8",
	// 	Amount:      salesBody.Amount,
	// 	Description: salesBody.Description,
	// 	Position: entities.Position{
	// 		Debet: entities.Transaction{
	// 			Code:   salesBody.Code,
	// 			Amount: salesBody.Amount,
	// 		},
	// 		Credit: entities.Transaction{
	// 			Code:   constant.COA_SALES,
	// 			Amount: salesBody.Amount,
	// 		},
	// 	},
	// }
	data, err := jS.JournalService.Sales(*salesBody)
	// communicate with service
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    data,
		"message": "ok",
	})
}

func (service *journalService) Purchase(c *fiber.Ctx) error {
	purchaseBody := new(dto.PurchaseReq)

	// parse body json
	if err := c.BodyParser(&purchaseBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	// validate body json
	if err := util.Validate(*purchaseBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err,
			"message": "fail",
		})
	}

	// convert dto to entity
	journal := entities.JournalEntity{
		ID:          uuid.NewString(),
		UserID:      "58dd4ecc-8cde-4ca3-b4f0-7451b7b59ce8",
		Amount:      purchaseBody.Amount,
		Description: purchaseBody.Description,
		Position: entities.Position{
			Debet: entities.Transaction{
				Code:   constant.COA_PURCHASE,
				Amount: purchaseBody.Amount,
			},
			Credit: entities.Transaction{
				Code:   constant.COA_ACCOUNT_PAYABLE,
				Amount: purchaseBody.Amount,
			},
		},
	}

	// communicate with service
	if err := service.JournalService.Purchase(journal); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "success inserted purchase",
		"message": "ok",
	})
}

func (service *journalService) General(c *fiber.Ctx) error {
	generalBody := new(dto.GeneralReq)

	// parse body json
	if err := c.BodyParser(&generalBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	// validate body json
	if err := util.Validate(*generalBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err,
			"message": "fail",
		})
	}

	// convert dto to entity
	journal := entities.JournalEntity{
		ID:          uuid.NewString(),
		UserID:      "58dd4ecc-8cde-4ca3-b4f0-7451b7b59ce8",
		Amount:      generalBody.Amount,
		Description: generalBody.Description,
		Position: entities.Position{
			Debet: entities.Transaction{
				Code:   generalBody.CodeDebet,
				Amount: generalBody.Amount,
			},
			Credit: entities.Transaction{
				Code:   generalBody.CodeCredit,
				Amount: generalBody.Amount,
			},
		},
	}

	// communicate with service
	if err := service.JournalService.General(journal); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "success inserted general",
		"message": "ok",
	})
}
