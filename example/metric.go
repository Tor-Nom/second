package main

import (
	"github.com/Tor-Nom/second"
)

var (
	NameSpace = "service_ebao"
	SubSystem = "open_api"

	policyGetUserInsureListService = second.InitPrometheusCounterVec(NameSpace, SubSystem, "policy_get_user_insure_list")
)
