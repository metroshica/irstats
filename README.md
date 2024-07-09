# irstats

[![Linters](https://github.com/metroshica/irstats/actions/workflows/lint.yml/badge.svg?branch=master)](https://github.com/metroshica/irstats/actions/workflows/lint.yml)
[![Test](https://github.com/metroshica/irstats/actions/workflows/test.yml/badge.svg)](https://github.com/metroshica/irstats/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/metroshica/irstats.svg)](https://pkg.go.dev/github.com/metroshica/irstats)
[![Go Report Card](https://goreportcard.com/badge/github.com/metroshica/irstats)](https://goreportcard.com/report/github.com/metroshica/irstats)
[![codecov](https://codecov.io/gh/metroshica/irstats/branch/master/graph/badge.svg?token=8S1FT6QP50)](https://codecov.io/gh/metroshica/irstats)

This package is an API "wrapper" for retrieving data from iRacing. We use the term "wrapper" loosely as iRacing does not yet have an officially documented API; However, we've done our best to build something that might resemble one.

The goal of this project is to provide access to iRacing stats in a manner that is convienent, flexible, and efficient. In using this package, if you find something in its design that goes against these goals, we want to know.

## Usage

```go
import "github.com/metroshica/irstats"
```

Construct a new irstats client, then make a request to fetch data from iRacing.
```go
irUsername := "<email address>"
irPassword := "<password>"

client, err := irstats.NewClient(irUsername, irPassword)
if err != nil {
  log.Fatalf("Failed to create client: %v", err)
}

subSessionID := "1000"
customerID := "1234"
data, _, err := client.SubSessionData(&subSessionID, &customerID)
```
