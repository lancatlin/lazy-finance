package ledger

type Command struct {
	Command string
	Dir     string
	Files   []string
}

func (c Command) genArgs() []string {
	args := make([]string, 0)
	for _, file := range c.Files {
		args = append(args, "-f", file)
	}
	args = append(args, c.Command)
	return args
}

// func (c LedgerCommand) Execute() (string, error) {
// 	cmd := exec.Command("hledger", c.Command)
// 	cmd.Dir = c.Dir
// 	out, err := cmd.CombinedOutput()
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(out), nil
// }
