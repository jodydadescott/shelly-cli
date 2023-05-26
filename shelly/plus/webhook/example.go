package webhook

// type WebHookParams struct {
// 	// Event which will trigger the execution of the webhook. Valid events are listed by Webhook.ListSupported.
// 	// Example values: switch.on, input.toggle_off. Required
// 	Event string `json:"event" yaml:"event"`
// 	// Cid Id of the component Required
// 	Cid int `json:"cid" yaml:"cid"`
// 	// Enable true to be enabled, false otherwise. It is false by default. Optional
// 	Enable bool `json:"enable" yaml:"enable"`
// 	// Name user-defined name for the webhook instance. Optional
// 	Name *string `json:"name" yaml:"name"`
// 	// SslCa type of the TCP sockets:
// 	// null : Plain TCP connection
// 	// user_ca.pem : TLS connection verified by the user-provided CA
// 	// ca.pem : TLS connection verified by the built-in CA bundle. Optional
// 	SslCa *string `json:"ssl_ca" yaml:"ssl_ca"`
// 	// URLs containing url addresses that will be called when the webhook event occurs. Each url
// 	// address is limited to 300 characters and the total number of url addresses associate with one
// 	// webhook is 5. At least one url address is Required
// 	URLs []string `json:"urls" yaml:"urls"`
// 	// ActiveBetween the first element indicates the start of the period during which the webhook will be active,
// 	// the second indicates the end of that period. Both start and end are strings in the format HH:MM,
// 	// where HH and MM are hours and minutes with optional leading zeros. To clear active_between its
// 	// value should be set to empty array or null. When active_between is empty, this attribute is not
// 	// visible in Webhook.List and the webhook is active all the time. Optional
// 	ActiveBetween []string `json:"active_between" yaml:"active_between"`
// 	// Condition hook trigger condition associated with event. Optional
// 	Condition *string `json:"condition" yaml:"condition"`
// 	// RepeatPeriod minimum interval for invocations of the hook. If set to negative the hook will be invoked only
// 	// once when the condition changes from false to true. If set to 0 the hook will be invoked every time the
// 	// triggering event occurs. Default is 0.Optional
// 	RepeatPeriod int `json:"repeat_period" yaml:"repeat_period"`
// }

func ExampleConfig() *Params {

	name := "name is optional"
	condition := "condition is optional"

	var urls []string
	urls = append(urls, "http://example.com")

	var activeBetween []string
	activeBetween = append(activeBetween, "active between is optional")

	return &Params{
		Event:         "eg. 'switch.on'. Use list supported to find supported for device",
		Enable:        true,
		Cid:           3,
		Name:          &name,
		URLs:          urls,
		Condition:     &condition,
		RepeatPeriod:  0,
		ActiveBetween: activeBetween,
	}
}
