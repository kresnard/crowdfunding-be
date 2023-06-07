package payment

import (
	"crowdfunding/transaction"
	"crowdfunding/user"
	"fmt"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type service struct {
}

type Service interface {
	GetToken(transaction transaction.Transaction, user user.User) string
}

func NewService() *service {
	return &service{}
}

func (s *service) GetToken(transaction transaction.Transaction, user user.User) string {
	midtrans.ServerKey = "YOUR-SERVER-KEY"
	midtrans.Environment = midtrans.Sandbox

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			Email: user.Email,
			FName: user.Name,
		},
	}

	snapResp := CreateTransaction(req)
	fmt.Println("Response :", snapResp)
}
