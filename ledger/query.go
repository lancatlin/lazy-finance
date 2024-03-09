package ledger

import "time"

type Query struct {
	Command string
	Keyword string    `form:"keyword"`
	Begin   time.Time `form:"begin" time_format:"2006-01-02"`
	End     time.Time `form:"end" time_format:"2006-01-02"`
}

func (q Query) GetArgs() []string {
	args := []string{"-f-", q.Command}

	if q.Keyword != "" {
		args = append(args, q.Keyword)
	}

	if !q.Begin.IsZero() {
		args = append(args, "--begin", q.Begin.Format("2006-01-02"))
	}

	if !q.End.IsZero() {
		args = append(args, "--end", q.End.Format("2006-01-02"))
	}

	args = append(args, "-O", "csv")
	return args
}
