package system

import (
	"github.com/jodydadescott/shelly-manager/shelly/plus/types"
)

type Config = types.SystemConfig
type Status = types.SystemStatus

type SysConfig = types.SystemConfig
type SysStatus = types.SystemStatus
type DeviceConfig = types.SystemDevice
type LocationConfig = types.SystemLocation
type DebugConfig = types.SystemDebug
type UIDataConfig = types.SystemUIData
type RPCUDPConfig = types.SystemRPCUDP
type SntpConfig = types.SystemSntp
type MqttDebug = types.SystemMqtt
type WebsocketDebug = types.SystemWebsocket
type UDP = types.SystemUDP

type SetReport = types.SetReport
type Error = types.Error
type Request = types.Request
type Response = types.Response
type MessageHandlerFactory = types.MessageHandlerFactory
type MessageHandler = types.MessageHandler
