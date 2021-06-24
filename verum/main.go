// Copyright (c) 2021 Satvik Reddy
package main

import (
	"github.com/SatvikR/liveassist/verum/config"
	"github.com/SatvikR/liveassist/verum/delivery/http"
)

func main() {
	config.LoadConfig(8081)

	http.StartServer(config.Port, []string{"http://localhost:3000"})
}
