package main

import (
	"github.com/Tor-Nom/second"
)

func main() {
	second.Metrics(policyGetUserInsureListService, "error", "GetPlicyInfo.res.gpi", second.LEVEL_ERROR, "GetPlicyInfo.res.gpi err")
}
