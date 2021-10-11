[![go tests](https://github.com/henjue/notebook_server/actions/workflows/test.yml/badge.svg)](https://github.com/henjue/notebook_server/actions/workflows/test.yml)

## 1. Install golang 1.17 or Higher

### On linux

```bash
go build -ldflags '-w -s' -o server
```

### On windows

```cmd

go build -ldflags '-w -s' -o server.exe
```