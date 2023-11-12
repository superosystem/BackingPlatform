package service

import (
	"strconv"

	"github.com/superosystem/BackingPlatform/backend/src/config"
	"github.com/superosystem/BackingPlatform/backend/src/entity"
	"github.com/superosystem/BackingPlatform/backend/src/model"
	"github.com/veritrans/go-midtrans"
)

type PaymentService interface {
	GetPaymentURL(payment model.Payment, user entity.User) (string, error)
}

type paymentService struct{}

func NewPaymentService() *paymentService {
	return &paymentService{}
}

func (s *paymentService) GetPaymentURL(payment model.Payment, user entity.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = config.LoadPayGateConfig().PGServerKey
	midclient.ClientKey = config.LoadPayGateConfig().PGClientKey
	// Sandbox for dev || Production for Prod
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(payment.ID),
			GrossAmt: int64(payment.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
