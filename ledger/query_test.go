package ledger

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestQueryGetArgs(t *testing.T) {
	q := Query{
		Command: "reg",
		Keyword: "assets",
		Begin:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		End:     time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	expected := []string{"-f-", "reg", "assets", "--begin", "2021-01-01", "--end", "2021-12-31", "-O", "csv"}
	actual := q.GetArgs()
	assert.Equal(t, expected, actual)
}
