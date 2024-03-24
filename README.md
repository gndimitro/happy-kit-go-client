# HappyKit Go Client (Inactive)

![Tests](https://github.com/gndimitro/happykit-go-client/workflows/Tests/badge.svg)
[![GoDoc](https://godoc.org/github.com/gndimitro/happykit-go-client?status.svg)](https://pkg.go.dev/github.com/gndimitro/happykit-go-client?tab=doc)
[![codecov](https://codecov.io/gh/gndimitro/happykit-go-client/branch/main/graph/badge.svg?token=QXUUJ67CCP)](https://codecov.io/gh/gndimitro/happykit-go-client)

A simple client library that interaces with the HappyKit feature flagging service

## Install

```
$ go get github.com/gndimitro/happykit-go-client
```

## Quick Start
```go
// 1. Import the library
// 2. Initialize the client
// 3. Call isEnabled where needed

import (
	"fmt"

	happyKitClient "github.com/gndimitro/happykit-go-client"
)

func main() {
	happyKitClient.Initialize("flags_pub_development_XXXXXXX")

	if happyKitClient.IsEnabled("testFlag") {
		fmt.Println("Flag is enabled")
	}
}
```