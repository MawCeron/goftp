<div align="center">

# goftp

An FTP client for Go, forked from [secsy/goftp](https://github.com/secsy/goftp) and maintained
for [lazyftp](https://github.com/MawCeron/lazyftp).

[![Go](https://img.shields.io/badge/Go-1.18+-00ADD8?style=for-the-badge&logo=go)](https://go.dev)
[![Build](https://img.shields.io/github/actions/workflow/status/MawCeron/goftp/ci.yml?style=for-the-badge)](https://github.com/MawCeron/goftp/actions)
[![License](https://img.shields.io/github/license/MawCeron/goftp?style=for-the-badge)](LICENSE)
[![Reference](https://img.shields.io/badge/pkg.go.dev-reference-007d9c?style=for-the-badge&logo=go)](https://pkg.go.dev/github.com/MawCeron/goftp)
[![Stars](https://img.shields.io/github/stars/MawCeron/goftp?style=for-the-badge)](https://github.com/MawCeron/goftp/stargazers)

</div>

---

## About

goftp aims to be a high-level FTP client that takes advantage of useful FTP features when the
server supports them. It was written by Muir Manders, and its last commit upstream is from June
2020.

This fork exists for lazyftp: changes land when lazyftp needs them, with no support promised
beyond that.

---

## Features

- Connection pooling for parallel transfers and traversal
- Automatic resumption of interrupted file transfers
- Explicit and implicit FTPS (TLS only, no SSL)
- IPv6 support
- Reasonably good automated tests that run against pure-ftpd and proftpd

---

## How to use

```
go get github.com/MawCeron/goftp
```

```go
config := goftp.Config{
    User:           "user",
    Password:       "password",
    Timeout:        10 * time.Second,
    IgnorePASVHost: true, // dial the control connection's host for PASV, see below
}

client, err := goftp.DialConfig(config, "ftp.example.com")
if err != nil {
    log.Fatal(err)
}

// Download a file into any io.Writer
var buf bytes.Buffer
if err := client.Retrieve("pub/readme.txt", &buf); err != nil {
    log.Fatal(err)
}

// Upload from any io.Reader
f, _ := os.Open("report.pdf")
if err := client.Store("report.pdf", f); err != nil {
    log.Fatal(err)
}
```

The [package documentation](https://pkg.go.dev/github.com/MawCeron/goftp) has the rest: listing and
walking directories, FTPS, resuming transfers and sending raw commands.

---

## Differences from upstream

- The module has a `go.mod` and its own path, so it can be used without a `replace` directive.
- `Config.IgnorePASVHost` sends PASV data connections to the control connection's host instead of
  the host in the server's reply, as EPSV already does. A server behind NAT announces a private
  address, and a hostile one can point the connection anywhere. It is off by default. Also proposed
  upstream in [secsy/goftp#84](https://github.com/secsy/goftp/pull/84).
- Listing a directory whose path contains a percent sign no longer corrupts the command.
- `go vet` passes and runs in CI.

---

## Tests

How to run the tests (Windows is not supported):

- `./build_test_server.sh` from the repository root, which downloads and compiles pure-ftpd and
  proftpd
- `go test`

Tests that start their own fake server need neither. With `GOFTP_SKIP_SERVERS=1` the real ones are
not started, so you can run just those tests:

```
GOFTP_SKIP_SERVERS=1 go test -run 'PASVHost|PercentSigns' .
```

CI does the same. The older tests still need the real servers.

---

## Contributing

Pull requests are welcome if they come with tests.

---

## License

Distributed under the MIT License. See [LICENSE](LICENSE) for more information.
