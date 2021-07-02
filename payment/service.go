package payment

import (
	"gocatering/model"
	"strconv"

	midtrans "github.com/veritrans/go-midtrans"
)

func GetPaymentUrl(trans *model.Transaction, user *model.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = ""
	midclient.ClientKey = ""
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.FullName,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(trans.ID),
			GrossAmt: int64(trans.Total),
		},
	}

	respon, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "something wrong guysss.... midtrans, error get token", err
	}

	return respon.RedirectURL, nil

}
