package light

import (
	"github.com/jodydadescott/shelly-manager/shelly/plus/types"
)

type Status = types.LightStatus
type Config = types.LightConfig
type Params = types.LightParams

type LightStatus = types.SwitchStatus
type SwitchConfig = types.LightConfig
type LightParams = types.LightParams

type SetReport = types.SetReport
type LightReport = types.LightReport
type Error = types.Error
type ShellyStatus = types.ShellyStatus
type Request = types.Request
type Response = types.Response
type MessageHandlerFactory = types.MessageHandlerFactory
type MessageHandler = types.MessageHandler
