# kinano-go

Go implement of きなの

```plaintext
$ go run dev/main.go -h
I am kinano v2

Usage:
  @BOT_kinano_v2 [flags]
  @BOT_kinano_v2 [command]

Available Commands:
  call        call custom functions
  help        Help about any command

Flags:
  -h, --help   help for @BOT_kinano_v2

Use "@BOT_kinano_v2 [command] --help" for more information about a command.
```

## Environment Variables

See [./pkg/config/](pkg/config/) for details.

- `TRAQ_BOT_ACCESS_TOKEN`
- `TRAQ_BOT_OAUTH2_CLIENT_ID`
- `TRAQ_BOT_OAUTH2_REDIRECT_URL`

## Run on traQ

```bash
go run main.go
```

## Run on Command Line

```bash
go run dev/main.go
```

## Add a command

```bash
go run github.com/spf13/cobra-cli@latest add {{CMDNAME}} --author {{YOURNAME}} --license mit
```

Then, edit `cmd/{{CMDNAME}}.go`.

## Add a custom function (called by `call` command)

```bash
go run github.com/spf13/cobra-cli@latest add {{FUNCNAME}} --author {{YOURNAME}} --license mit
```

Then, move `cmd/{{FUNCNAME}}.go` to `cmd/callevent/{{FUNCNAME}}.go` and edit it.
