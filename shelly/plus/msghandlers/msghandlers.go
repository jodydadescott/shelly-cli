package msghandlers

import (
	"time"

	"github.com/jodydadescott/shelly-manager/shelly/plus/types"

	"github.com/jodydadescott/shelly-manager/shelly/plus/msghandlers/ws"
)

type Request = types.Request
type MessageHandlerFactory = types.MessageHandlerFactory
type MessageHandler = types.MessageHandler

type Config struct {
	Hostname     string
	Password     string
	Username     string
	SendTimeout  time.Duration
	DebugEnabled bool
}

func (t *Config) GetHostname() string {
	return t.Hostname
}

func (t *Config) GetUsername() string {
	return t.Username
}

func (t *Config) GetPassword() string {
	return t.Password
}

func (t *Config) GetSendTimeout() time.Duration {
	return t.SendTimeout
}

func (t *Config) IsDebugEnabled() bool {
	return t.DebugEnabled
}

func NewWS(config *Config) (MessageHandlerFactory, error) {
	return ws.New(config)
}
