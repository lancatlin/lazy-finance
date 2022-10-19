# Ledger Quick Note

![screenshot](screenshots/home.png)

**Add ledger transactions on the fly!**

## Feature

### Transaction Template

add your transaction template in `tx/` (in Go's template syntax), and create transaction from them on the fly. 

Examples:

Take some cash
```
{{ .Date }} * cash
    expenses:cash    ${{ .Amount }}
    assets:cash

```

Cash expenses
```
{{ .Date }} {{ with .Name }}{{ . }}{{ else }}{{ .Account }}{{ end }}
    {{ .Account }}      ${{ .Amount }}
    expenses:cash
    
```

Checkout `tx/` folder for more examples.

### Ledger Scripts

Run some commonly used ledger commands.

Define your commands in config.go

```go
var SCRIPTS = map[string][]string{
	"balance assets":     {"b", "assets", "-X", "$"},
	"register":           {"r", "--tail", "10"},
	"balance this month": {"b", "-b", "this month"},
}
```

Execute them and see the result in the browser.

![execute result](screenshots/exec.png)

## Install

Requirements:
* go
* ledger (Only required when you use scripts)

```
git clone https://github.com/lancatlin/ledger-quicknote.git
```

```
go build
```

```
./ledger-quicknote
```

Checkout `deployment/` for Nginx & Systemd example configuration.

