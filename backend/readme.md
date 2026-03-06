# Golang Architecture

## Command Line

- Run server

running service

```bash
go run ./cmd/_cli
```
or
```bash
make run
```

- Create template

command for create template

```bash
go run ./cmd/_cli make:template <file_name> <type>
```
or
```bash
make template name=<file_name> type=<type>
```

- Build project

command for build the project

```bash
go build ./cmd/server
```
or
```bash
make build
```