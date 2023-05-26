package webhook

import (
	"github.com/jodydadescott/shelly-manager/shelly/plus/types"
)

type Webhook = types.WebhookHook
type Webhooks = types.Webhooks
type WebhookParams = types.WebhookParams
type Params = types.WebhookParams
type WebhookEventInputToggleOn = types.WebhookEventInputToggleOn
type WebhookEventInputToggleOff = types.WebhookEventInputToggleOff
type WebhookEventInputButtonPush = types.WebhookEventInputButtonPush
type WebhookEventInputButtonLongpush = types.WebhookEventInputButtonLongpush
type WebhookEventInputButtonDoublepush = types.WebhookEventInputButtonDoublepush
type WebhookEventSwitchOff = types.WebhookEventSwitchOff
type WebhookEventSwitchOn = types.WebhookEventSwitchOn
type WebhookEventTemperatureChange = types.WebhookEventTemperatureChange
type WebhookAttrs = types.WebhookAttrs
type WebhookTypes = types.WebhookTypes

type SetReport = types.SetReport
type Error = types.Error
type Request = types.Request
type Response = types.Response
type MessageHandlerFactory = types.MessageHandlerFactory
type MessageHandler = types.MessageHandler
