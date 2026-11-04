# goftp - an FTP client for Go

[![CI](https://github.com/MawCeron/goftp/actions/workflows/ci.yml/badge.svg)](https://github.com/MawCeron/goftp/actions/workflows/ci.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/MawCeron/goftp.svg)](https://pkg.go.dev/github.com/MawCeron/goftp)

This is a fork of [secsy/goftp](https://github.com/secsy/goftp) by Muir Manders, whose last commit is from June 2020. It exists for [lazyftp](https://github.com/MawCeron/lazyftp): changes land when lazyftp needs them, with no support promised beyond that. Pull requests are welcome if they come with tests.

Import it as `github.com/MawCeron/goftp`. The API is upstream's, with the changes below.

### Differences from upstream

* The module has a `go.mod` and its own path, so it can be used without a `replace` directive.
* `Config.IgnorePASVHost` sends PASV data connections to the control connection's host instead of the host in the server's reply, as EPSV already does. A server behind NAT announces a private address, and a hostile one can point the connection anywhere. It is off by default. Also proposed upstream in [secsy/goftp#84](https://github.com/secsy/goftp/pull/84).
* Listing a directory whose path contains a percent sign no longer corrupts the command.
* `go vet` passes and runs in CI.

### About ###
goftp aims to be a high-level FTP client that takes advantage of useful FTP features when supported by the server.

Here are some notable package highlights:

* Connection pooling for parallel transfers/traversal.
* Automatic resumption of interruped file transfers.
* Explicit and implicit FTPS support (TLS only, no SSL).
* IPv6 support.
* Reasonably good automated tests that run against pure-ftpd and proftpd.

Please see the godocs for details and examples.

### Tests ###

How to run tests (windows not supported):
* ```./build_test_server.sh``` from root goftp directory (this downloads and compiles pure-ftpd and proftpd)
* ```go test``` from the root goftp directory

Tests that start their own fake server need neither. With `GOFTP_SKIP_SERVERS=1` the real ones are not started, so you can run just those tests: `GOFTP_SKIP_SERVERS=1 go test -run 'PASVHost|PercentSigns' .`. CI does the same. The older tests still need the real servers.
