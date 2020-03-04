package utils

import (
	"os"
	"strconv"
)

//GetENV ...
func GetENV(experimentDetails *ExperimentDetails, expName string) {
	experimentDetails.ExperimentName = expName
	experimentDetails.ChaosDuration, _ = strconv.Atoi(os.Getenv("CHAOS_DURATION"))
	experimentDetails.ChaosInterval, _ = strconv.Atoi(os.Getenv("CHAOS_INTERVAL"))
	experimentDetails.RampTime, _ = strconv.Atoi(os.Getenv("RAMP_TIME"))
	experimentDetails.Force, _ = strconv.ParseBool(os.Getenv("FORCE"))
	experimentDetails.ChaosLib = "litmus"
	experimentDetails.ChaosServiceAccount = os.Getenv("CHAOS_SERVICE_ACCOUNT")
	experimentDetails.AppNS = "shubham"
	experimentDetails.AppLabel = "run=nginx"
	experimentDetails.AppKind = os.Getenv("APP_KIND")
	experimentDetails.KillCount, _ = strconv.Atoi(os.Getenv("KILL_COUNT"))
	experimentDetails.ChaosUID = os.Getenv("CHAOS_UID")
	experimentDetails.AuxiliaryAppInfo = "shubham:run=nginx,shubham:run=xyz"
}
