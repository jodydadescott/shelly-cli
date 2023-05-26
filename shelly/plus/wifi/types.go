package wifi

import (
	"github.com/jodydadescott/shelly-manager/shelly/plus/types"
)

type Config = types.WifiConfig
type Status = types.WifiStatus

type WifiConfig = types.WifiConfig
type WifiStatus = types.WifiStatus
type WifiScanResults = types.WifiScanResults
type APClients = types.WifiAPClients
type WifiNet = types.WifiNet
type APClient = types.WifiAPClient
type APConfig = types.WifiAP
type StaConfig = types.WifiSTA
type RoamConfig = types.WifiRoam
type RangeExtenderConfig = types.WifiRangeExtender

type SetReport = types.SetReport
type Error = types.Error
type Request = types.Request
type Response = types.Response
type MessageHandlerFactory = types.MessageHandlerFactory
type MessageHandler = types.MessageHandler
