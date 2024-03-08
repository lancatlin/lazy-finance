package model

import "testing"

func TestConvertAmount(t *testing.T) {
	tests := []struct {
		money         string
		wantAmount    float64
		wantCommodity string
		wantErr       bool
	}{
		{"100 USD", 100, "USD", false},
		{"$100", 100, "$", false},
		{"100.5", 100.5, "", false},
		{"100.5 USD", 100.5, "USD", false},
		{"100.5 $", 100.5, "$", false},
		{"$-100.5", -100.5, "$", false},
		{"-100 USD", -100, "USD", false},
		{"", 0, "", true},
		{"100.5.5 USD", 0, "", true},
		{"100 USD EUR", 0, "", true},
		{"XYZ", 0, "", true},
	}

	for _, tt := range tests {
		amount, commodity, err := ConvertAmount(tt.money)
		if tt.wantErr {
			if err == nil {
				t.Errorf("Testcase %s: Expected an error, but got nil", tt.money)
			}
		} else {
			if err != nil {
				t.Errorf("Testcase %s: Did not expect an error, but got: %v", tt.money, err)
			}
			if tt.wantAmount != amount {
				t.Errorf("Testcase %s: Expected amount %v, got %v", tt.money, tt.wantAmount, amount)
			}
			if tt.wantCommodity != commodity {
				t.Errorf("Testcase %s: Expected commodity %s, got %s", tt.money, tt.wantCommodity, commodity)
			}
		}
	}
}

func TestConvertRegisterToAccount(t *testing.T) {
	reg := Register{
		Account: "assets",
		Amount:  "100 USD",
	}
	account, err := reg.ToAccount()
	assertNil(t, err)
	expectedAccount := Account{
		Name:      "assets",
		Amount:    100,
		Commodity: "USD",
	}
	assertEqual(t, expectedAccount, account)
}
