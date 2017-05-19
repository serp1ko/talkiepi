package talkiepi

import (
	"crypto/tls"

	"github.com/dchote/gpio"
	"github.com/dchote/gumble/gumble"
	"github.com/dchote/gumble/gumbleopenal"
)

// Raspberry Pi GPIO pin assignments (CPU pin definitions)
const (
	OnlineLEDPin        uint = 5
	ParticipantsLEDPin  uint = 6
	TransmitLEDPin      uint = 12
	PushToTalkButtonPin uint = 26
	SelectButtonPin     uint = 16
)

type Talkiepi struct {
	Config *gumble.Config
	Client *gumble.Client

	Address   string
	TLSConfig tls.Config

	ConnectAttempts uint

	Stream *gumbleopenal.Stream

	ChannelName    string
	IsConnected    bool
	IsTransmitting bool

	GPIOEnabled           bool
	OnlineLED             gpio.Pin
	ParticipantsLED       gpio.Pin
	TransmitLED           gpio.Pin
	PushToTalkButton      gpio.Pin
	PushToTalkButtonState uint
	SelectButton          gpio.Pin
	SelectButtonState     uint
}
