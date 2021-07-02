package transaction

import (
	"gocatering/helper"
	"gocatering/model"
	"gocatering/payment"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type TransactionHandler struct {
	service Service
}

func NewTransactionHandler(s Service) *TransactionHandler {
	return &TransactionHandler{service: s}
}

func (h *TransactionHandler) GetAllTransaction(e echo.Context) error {
	transactions, err := h.service.GetAllTransaction()
	if err != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(err.Error(), http.StatusInternalServerError, "failed", nil))
	}

	var formatTransResponse []model.TransactionResponse
	for _, transaction := range transactions {
		transactionResponse := model.TransactionResponse{
			TransID:       transaction.ID,
			PaketName:     transaction.Paket.Name,
			Quantity:      transaction.Quantity,
			Total:         transaction.Total,
			Location:      transaction.Location,
			RegentName:    transaction.Regency.Name,
			CustomerName:  transaction.User.FullName,
			DeliveredTime: transaction.DeliverTime,
			PaymentURL:    "",
			Note:          transaction.Note,
		}

		formatTransResponse = append(formatTransResponse, transactionResponse)
	}

	return e.JSON(http.StatusCreated,
		helper.Apiresponse("List of Transaction", http.StatusOK, "success", formatTransResponse))

}

func (h *TransactionHandler) FindTransactionByUserId(e echo.Context) error {
	id := 1
	transactions, err := h.service.FindTransactionByUserId(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(err.Error(), http.StatusInternalServerError, "failed", nil))
	}

	var formatTransResponse []model.TransactionResponse
	for _, transaction := range transactions {
		transactionResponse := model.TransactionResponse{
			TransID:       transaction.ID,
			PaketName:     transaction.Paket.Name,
			Quantity:      transaction.Quantity,
			Total:         transaction.Total,
			Location:      transaction.Location,
			RegentName:    transaction.Regency.Name,
			CustomerName:  transaction.User.FullName,
			DeliveredTime: transaction.DeliverTime,
			PaymentURL:    "",
			Note:          transaction.Note,
		}

		formatTransResponse = append(formatTransResponse, transactionResponse)
		return e.JSON(http.StatusCreated,
			helper.Apiresponse("List of Transactiio user", http.StatusOK, "success", formatTransResponse))

	}

	return e.JSON(http.StatusCreated,
		helper.Apiresponse("List of Transaction", http.StatusOK, "success", formatTransResponse))

}

func (h *TransactionHandler) CreateTransaction(e echo.Context) error {

	var trans model.Transaction

	e.Bind(&trans)

	paket, err := h.service.FindPaketById(trans.PaketID)
	if err != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(err.Error(), http.StatusInternalServerError, "failed", nil))
	}

	user, err := h.service.FindUserById(trans.UserID)
	if err != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(err.Error(), http.StatusInternalServerError, "failed", nil))
	}

	regency, err := h.service.FindRegencyById(trans.RegencyID)

	if err != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(err.Error(), http.StatusInternalServerError, "failed", nil))
	}
	//	user, err := h.service.FindUserById(trans.UserID)

	qty := trans.Quantity
	price := paket.Price
	shipping_cost := regency.ShippingCost
	subtotal := (qty * price)
	if paket.Discount > 0 {
		discount := (paket.Discount / 100) * subtotal
		subtotal = subtotal - discount

	}
	total := subtotal + shipping_cost

	newTrans := model.Transaction{
		UserID:      trans.UserID,
		PaketID:     trans.PaketID,
		Quantity:    qty,
		Total:       total,
		Location:    trans.Location,
		RegencyID:   trans.RegencyID,
		Status:      "pending",
		DeliverTime: trans.DeliverTime,
		Note:        trans.Note,
	}

	errors := h.service.CreateTransaction(&newTrans)

	if errors != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(errors.Error(), http.StatusInternalServerError, "failed", nil))
	}

	paymentUrl, err := payment.GetPaymentUrl(&newTrans, user)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error during get payment url",
			"message": err.Error,
		})
	}

	newTrans.PaymentUrl = paymentUrl
	getErrors := h.service.UpdateTransaction(newTrans.ID, &newTrans)
	if getErrors != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(errors.Error(), http.StatusInternalServerError, "failed", nil))
	}

	transactionResponse := model.TransactionResponse{
		TransID:       newTrans.ID,
		PaketName:     paket.Name,
		Quantity:      newTrans.Quantity,
		Total:         newTrans.Total,
		Location:      newTrans.Location,
		RegentName:    regency.Name,
		CustomerName:  user.FullName,
		DeliveredTime: newTrans.DeliverTime,
		PaymentURL:    paymentUrl,
		Note:          newTrans.Note,
	}

	return e.JSON(http.StatusOK,
		helper.Apiresponse("transaction succes", http.StatusOK, "succes", transactionResponse))
}

func (h *TransactionHandler) GetNotif(e echo.Context) error {
	notif := model.MidtransNotification{}
	e.Bind(&notif)

	transaction_id, _ := strconv.Atoi(notif.OrderID)

	//config.DB.Where("id=?", transaction_id).First(&transaction)
	trans, err := h.service.FindTransactionById(transaction_id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(err.Error(), http.StatusInternalServerError, "failed", nil))
	}

	if notif.PaymentType == "credit_card" && notif.TransactionStatus == "capture" && notif.FraudStatus == "accept" {
		trans.Status = "paid"
	} else if notif.TransactionStatus == "settlement" {
		trans.Status = "paid"
	} else if notif.TransactionStatus == "deny" || notif.TransactionStatus == "expire" || notif.TransactionStatus == "cancel" {
		trans.Status = "cancelled"
	}

	errors := h.service.UpdateTransaction(trans.ID, trans)
	if errors != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(err.Error(), http.StatusInternalServerError, "failed", nil))
	}

	return e.JSON(http.StatusOK,
		helper.Apiresponse("update transaction success", http.StatusOK, "succes", notif))
}
