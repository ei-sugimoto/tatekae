package model

import (
	"sort"

	"github.com/ei-sugimoto/tatekae/api/infrastructure/ent"
)

type Bill struct {
	ID        int    `json:"id"`
	Price     int    `json:"price"`
	ProjectID int    `json:"project_id"`
	SrcUser   int    `json:"src_user"`
	DstUser   int    `json:"dst_user"`
	SrcName   string `json:"src_name"`
	DstName   string `json:"dst_name"`
}

type Transaction struct {
	DstUser int
	SrcUser int
	Amount  int
}

type SumarizeBill struct {
	DstUserName string
	SrcUserName string
	Amount      int
}

type SumarizeBills []SumarizeBill

type PersonBlance struct {
	UserID  int
	Balance int
}

func CalcucateSummaries(bills []*Bill) ([]Transaction, error) {

	if len(bills) == 0 {
		return nil, &ent.NotFoundError{}
	}

	balances := make(map[int]int)
	for _, bill := range bills {
		balances[bill.SrcUser] += bill.Price
		balances[bill.DstUser] -= bill.Price
	}

	crediter := make([]PersonBlance, 0)
	debiter := make([]PersonBlance, 0)

	for userID, balance := range balances {
		if balance > 0 {
			crediter = append(crediter, PersonBlance{UserID: userID, Balance: balance})
		} else if balance < 0 {
			debiter = append(debiter, PersonBlance{UserID: userID, Balance: balance})
		}
	}

	sort.Slice(crediter, func(i, j int) bool {
		return crediter[i].Balance > crediter[j].Balance
	})

	sort.Slice(debiter, func(i, j int) bool {
		return debiter[i].Balance < debiter[j].Balance
	})

	minTransactions := make([]Transaction, 0)
	i, j := 0, 0

	for i < len(crediter) && j < len(debiter) {
		credit := crediter[i]
		debit := debiter[j]

		amount := min(credit.Balance, -debit.Balance)

		minTransactions = append(minTransactions, Transaction{
			SrcUser: debit.UserID,
			DstUser: credit.UserID,
			Amount:  amount,
		})

		crediter[i].Balance -= amount
		debiter[j].Balance += amount

		if crediter[i].Balance == 0 {
			i++
		}

		if debiter[j].Balance == 0 {
			j++
		}
	}

	return SortTransactions(minTransactions), nil

}

func SortTransactions(transactions []Transaction) []Transaction {
	sort.Slice(transactions, func(i, j int) bool {
		if transactions[i].SrcUser != transactions[j].SrcUser {
			return transactions[i].SrcUser < transactions[j].SrcUser
		}
		if transactions[i].DstUser != transactions[j].DstUser {
			return transactions[i].DstUser < transactions[j].DstUser
		}
		return transactions[i].Amount < transactions[j].Amount
	})

	return transactions
}
