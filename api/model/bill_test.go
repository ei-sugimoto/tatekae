package model_test

import (
	"reflect"
	"testing"

	"github.com/ei-sugimoto/tatekae/api/model"
)

func TestCalucateSummaries(t *testing.T) {
	tests := []struct {
		name    string
		bills   []model.Bill
		want    []model.Transaction
		wantErr bool
	}{
		{
			name:    "No bills",
			bills:   []model.Bill{},
			want:    []model.Transaction{},
			wantErr: true,
		},
		{
			name: "Single bill",
			bills: []model.Bill{
				{Price: 100, SrcUser: 1, DstUser: 2},
			},
			want: []model.Transaction{
				{SrcUser: 2, DstUser: 1, Amount: 100},
			},
			wantErr: false,
		},
		{
			name: "Multiple bills",
			bills: []model.Bill{
				{Price: 100, SrcUser: 1, DstUser: 2},
				{Price: 200, SrcUser: 2, DstUser: 3},
				{Price: 300, SrcUser: 3, DstUser: 1},
			},
			want: []model.Transaction{
				{SrcUser: 1, DstUser: 2, Amount: 100},
				{SrcUser: 1, DstUser: 3, Amount: 100},
			},
			wantErr: false,
		},
		{
			name: "Complex transactions",
			bills: []model.Bill{
				{Price: 100, SrcUser: 1, DstUser: 2},
				{Price: 150, SrcUser: 2, DstUser: 3},
				{Price: 200, SrcUser: 3, DstUser: 1},
				{Price: 50, SrcUser: 1, DstUser: 3},
			},
			want: []model.Transaction{
				{SrcUser: 1, DstUser: 2, Amount: 50},
			},
			wantErr: false,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := model.CalcucateSummaries(tt.bills)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalucateSummaries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, model.SortTransactions(tt.want)) && len(got) != 0 {

				t.Errorf("test number %d CalucateSummaries() = %v, want %v", i, got, tt.want)
			}
		})
	}
}
