package types

// This service allows Shelly devices to send HTTP requests triggered by events. Events usually occur
// when functional device parts change their state (switch toggling, button pushes, sensor readings,
// etc.). There is a limit of 20 hook instances per device (10 for battery-operated devices).
// The same event may trigger any number of webhooks associated with it. A revision number is maintained
// which is incremented on every update of the schedules. It is returned in the result of RPC calls and
// is also included in the Status object of the Sys component

// Each webhook is associated with an event type and component instance. The association is determined as follows:

// The event attribute of a hook is formatted as <component type>.<event type>. It identifies the type of component,
// and the specific event for this component type.
// The cid attribute of a hook is the id of the component instance.
// <component type> together with <cid> specify the component instance, identified as <component type>:<cid>.

// Conditional execution
// Some events support attributes - additional pieces of information about the event which occurred.

// Hooks support conditions: logical statements which yield a boolean result whether to trigger the webhook or not.
// The conditions can use information from the system config, status and device info available as config, status and
// info objects with the same structure as returned by Shelly.GetConfig, Shelly.GetStatus and Shelly.GetDeviceInfo.
// Additionally, if the event supports attributes, their information is available in an ev or event object.
// Condition statements must be valid script expressions.

// Example: for a webhook associated with a temperature component, Webhook.Create
// {"cid":0, "enable":true, "event":"temperature.change", "urls":["http://example.com"], "condition":"event.tC > 20"}
// will only invoke the url if attribute tC of event temperature.change is above 20. In this case the attribute tC of
// the ev object is used in the condition string.

// Repeatability
// Hooks support also repeat period: number specifying the minimum interval in seconds for two consecutive invocations
// of the hook. If events which trigger the hook during this interval occur they will be suppressed. A negative repeat_period
// is interperted as "never repeat". Such webhooks will only be executed once, as their condition transitions from false to true.
// The default for repeat_period is 0, which means the webhook will be executed every time the event occurs.

// URL token replacement
// Hooks events also support url token replacement. Before the url is invoked, it is parsed for notations in format ${token}.
// token can be a valid javascript expression. ${token} will be replaced with the result of the evaluation of this expression.
// If evaluation fails the contents of token are copied verbatim. During evaluation the objects config, status, info and ev or
// event for events with attributes are available as for conditions:

// status is an object which contains the entire device status as returned by Shelly.GetStatus
// config is an object which contains the entire device configuration as returned by Shelly.GetConfig
// info is an object which contains device info as returned by Shelly.GetDeviceInfo
// Interpolated token values are urlencoded.

// Examples
// http://example.com/endpoint?tC=${ev.tC} for webhooks associated with temperature.change events will replace ${ev.tC}
// with the actual temperature measured by the sensor
// http://example.com/endpoint?hum=${ev.rh}&temp=${status["temperature:0"].tC}&batt=${status["devicepower:0"].battery.V}
// will include temperature, battery voltage, along with humidity measured by sensor on a PlusHT device, triggered by a
// humidity.change event.
// http://example.com/endpoint?uptime=${status.sys.uptime} will include the uptime in seconds as a query string parameter
// http://example.com/endpoint?mac=${config.sys.device.mac} will include the MAC address of the device as a query string parameter
// http://example.com/endpoint?switch=${status["switch:0"].output will include output state of Switch component with id 0.
// NOTE
// For example Webhook.Create {"cid":0, "enable":true, "event":"temperature.change", "urls":["http://example.com/tC=${ev.tC}"]}
// will replace ${ev.tC} with the actual temperature, for instance http://example.com/tC=20.00.

// NOTE
// To insert a literal ${ in the url without triggering a replacement, $${ should be used as an escape sequence.
// That is $${ will get swapped with ${ and no token interpolation will take place. For example, http://example.com/tC=$${ev.tC}
// will be invoked as http://example.com/tC=${ev.tC}.

// NOTE
// When upgrading from a firmware version which does not support url token replacement to one which does,
// existing urls will be migrated: if they contain the ${ literal, it will be swapped with $${.

// WebhookHook WebHook
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Webhook/#webhooklist
type WebhookHook struct {
	// ID of the webhook
	ID *int `json:"id" yaml:"id"`
	// Cid Id of the component
	Cid int `json:"cid" yaml:"cid"`
	// Enable true to be enabled, false otherwise
	Enable bool `json:"enable" yaml:"enable"`
	// Event which will trigger the execution of the webhook. Valid events are listed by Webhook.ListSupported.
	// Example values: switch.on, input.toggle_off.
	Event string `json:"event" yaml:"event"`
	// Name sser-defined name for the webhook instance
	Name *string `json:"name" yaml:"name"`
	// SslCa type of the TCP sockets:
	// null : Plain TCP connection
	// user_ca.pem : TLS connection verified by the user-provided CA
	// ca.pem : TLS connection verified by the built-in CA bundle
	SslCa *string `json:"ssl_ca" yaml:"ssl_ca"`
	// URLs Containing url addresses that will be called when the webhook event occurs
	URLs []string `json:"urls" yaml:"urls"`
	// Condition hook trigger condition associated with event.
	Condition any `json:"condition" yaml:"condition"`
	// RepeatPeriod minimum interval for invocations of the hook.
	RepeatPeriod int `json:"repeat_period" yaml:"repeat_period"`
}

// WebhookParams
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Webhook/#webhookcreate &
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Webhook/#webhookupdate
type WebhookParams struct {
	// Event which will trigger the execution of the webhook. Valid events are listed by Webhook.ListSupported.
	// Example values: switch.on, input.toggle_off. Required
	Event string `json:"event" yaml:"event"`
	// Cid Id of the component Required
	Cid int `json:"cid" yaml:"cid"`
	// Enable true to be enabled, false otherwise. It is false by default. Optional
	Enable bool `json:"enable" yaml:"enable"`
	// Name user-defined name for the webhook instance. Optional
	Name *string `json:"name" yaml:"name"`
	// SslCa type of the TCP sockets:
	// null : Plain TCP connection
	// user_ca.pem : TLS connection verified by the user-provided CA
	// ca.pem : TLS connection verified by the built-in CA bundle. Optional
	SslCa *string `json:"ssl_ca" yaml:"ssl_ca"`
	// URLs containing url addresses that will be called when the webhook event occurs. Each url
	// address is limited to 300 characters and the total number of url addresses associate with one
	// webhook is 5. At least one url address is Required
	URLs []string `json:"urls" yaml:"urls"`
	// ActiveBetween the first element indicates the start of the period during which the webhook will be active,
	// the second indicates the end of that period. Both start and end are strings in the format HH:MM,
	// where HH and MM are hours and minutes with optional leading zeros. To clear active_between its
	// value should be set to empty array or null. When active_between is empty, this attribute is not
	// visible in Webhook.List and the webhook is active all the time. Optional
	ActiveBetween []string `json:"active_between" yaml:"active_between"`
	// Condition hook trigger condition associated with event. Optional
	Condition *string `json:"condition" yaml:"condition"`
	// RepeatPeriod minimum interval for invocations of the hook. If set to negative the hook will be invoked only
	// once when the condition changes from false to true. If set to 0 the hook will be invoked every time the
	// triggering event occurs. Default is 0.Optional
	RepeatPeriod int `json:"repeat_period" yaml:"repeat_period"`
}

// WebhookEventInputToggleOn
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Webhook/#webhooklistsupported
type WebhookEventInputToggleOn struct {
}

// WebhookEventInputToggleOff
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Webhook/#webhooklistsupported
type WebhookEventInputToggleOff struct {
}

// WebhookEventInputButtonPush
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Webhook/#webhooklistsupported
type WebhookEventInputButtonPush struct {
}

// WebhookEventInputButtonLongpush
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Webhook/#webhooklistsupported
type WebhookEventInputButtonLongpush struct {
}

// WebhookEventInputButtonDoublepush
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Webhook/#webhooklistsupported
type WebhookEventInputButtonDoublepush struct {
}

// WebhookEventSwitchOff
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Webhook/#webhooklistsupported
type WebhookEventSwitchOff struct {
}

// WebhookEventSwitchOn
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Webhook/#webhooklistsupported
type WebhookEventSwitchOn struct {
}

// WebhookEventTemperatureChange
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Webhook/#webhooklistsupported
type WebhookEventTemperatureChange struct {
	Attrs []WebhookAttrs `json:"attrs" yaml:"attrs"`
}

// WebhookAttrs Since version 0.11.0 Contains all events that can be used to trigger a webhook, extended with a
// declaration of supported event attributes. Events are listed as JSON objects with keys in format component
// type,event type. Supported event attributes, if such exist, are represented in attrs array
type WebhookAttrs struct {
	// Name attribute name
	Name string `json:"name" yaml:"name"`
	Type string `json:"type" yaml:"type"`
	Desc string `json:"desc" yaml:"desc"`
}

// WebhookTypes
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Webhook/#webhooklistsupported
type WebhookTypes struct {
	InputToggleOn         WebhookEventInputToggleOn         `json:"input.toggle_on,omitempty" yaml:"input.toggle_on,omitempty"`
	InputToggleOff        WebhookEventInputToggleOff        `json:"input.toggle_off,omitempty" yaml:"input.toggle_off,omitempty"`
	InputButtonPush       WebhookEventInputButtonPush       `json:"input.button_push,omitempty" yaml:"input.button_push,omitempty"`
	InputButtonLongpush   WebhookEventInputButtonLongpush   `json:"input.button_longpush,omitempty" yaml:"input.button_longpush,omitempty"`
	InputButtonDoublepush WebhookEventInputButtonDoublepush `json:"input.button_doublepush,omitempty" yaml:"input.button_doublepush,omitempty"`
	SwitchOff             WebhookEventSwitchOff             `json:"switch.off,omitempty" yaml:"switch.off,omitempty"`
	SwitchOn              WebhookEventSwitchOn              `json:"switch.on,omitempty" yaml:"switch.on,omitempty"`
	TemperatureChange     WebhookEventTemperatureChange     `json:"temperature.change,omitempty" yaml:"temperature.change,omitempty"`
}

// Webhooks
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Webhook/#webhooklist
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Webhook/#webhooklistsupported
type Webhooks struct {
	ID int `json:"id" yaml:"id"`
	// Hooks list of hooks
	Hooks []WebhookHook `json:"hooks,omitempty" yaml:"hooks,omitempty"`
	// Rev current revision number of the webhook instances
	Rev int `json:"rev" yaml:"rev"`
	// Types of events
	Types WebhookTypes `json:"types,omitempty" yaml:"types,omitempty"`
}
