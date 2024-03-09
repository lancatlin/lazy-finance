package ledger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadBalances(t *testing.T) {
	input := `"account","balance"
"assets","$-250, 1 USD"
"assets:cash","$-1100"
"expenses","$100, -1 USD"
"expenses:food","$940"`
	expected := []Balance{
		{
			Account: "assets",
			Balance: "$-250, 1 USD",
		},
		{
			Account: "assets:cash",
			Balance: "$-1100",
		},
		{
			Account: "expenses",
			Balance: "$100, -1 USD",
		},
		{
			Account: "expenses:food",
			Balance: "$940",
		},
	}
	balances, err := LoadBalances(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, balances)
}
