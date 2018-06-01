[![Mozilla Public License](https://img.shields.io/badge/license-MPL-blue.svg)](https://www.mozilla.org/MPL)
[![Go Report Card](https://goreportcard.com/badge/github.com/myENA/hcheck)](https://goreportcard.com/report/github.com/myENA/hcheck)

# hcheck

## Summary

This repo povides a very simple tool called `hcheck` that will check the HTTP return code of a given URL and verify it matches your expected code.  This tool is useful to ensure expected results when ran in parrallel with load testing activity.  The tool can operate in a one-and-done mode or consistently polling in the foreground and logging any unexpected result.

In addition to checking, we use [nathanejohnson/intransport](https://github.com/nathanejohnson/intransport) when fetching secure resources to
automatically fetch intermediate certificates and do full chain verification including stapled OCSP responses.

## Installing

Users with a proper Go environment (1.8+ required) ...

```
go get -u github.com/myENA/hcheck
```

Developers that wish to take advantage of vendoring and other options ...

```
git clone https://github.com/myENA/hcheck.git
cd hcheck
make
```

## Usage

### Summary

```
ahurt$ ./hcheck
Usage: hcheck [--version] [--help] <command> [<args>]

Available commands are:
    check    Validate host return code and exit.
    watch    Poll target host and validate return code.
```

### Common Options

| Option     | Description                       | Default |
|------------|-----------------------------------|---------|
| `url`      | The fully qualified URL to check  | `empty`
| `code`     | The expected return code          | `200`
| `timeout`  | Time to wait for a response       | `5s`
| `insecure` | Skip TLS validation               | `false`

### Watch Specific Options

| Option     | Description         | Default |
|------------|---------------------|---------|
| `interval` | Watch poll interval | `5s`

### Example

```
ahurt$ ./hcheck check -url www.google.com -code 400
2018/06/01 14:44:13 [Error] Check failed: expected 400, got 200
ahurt$ echo $?
1
```
