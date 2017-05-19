package talkiepi

import (
	"fmt"
	"time"

	"github.com/dchote/gpio"
	"github.com/stianeikeland/go-rpio"
)

func (b *Talkiepi) initGPIO() {
	// we need to pull in rpio to pullup our button pin
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		b.GPIOEnabled = false
		return
	} else {
		b.GPIOEnabled = true
	}

	PTTButtonPinPullUp := rpio.Pin(PushToTalkButtonPin)
	PTTButtonPinPullUp.PullUp()

	SelectButtonPinPullUp := rpio.Pin(SelectButtonPin)
	SelectButtonPinPullUp.PullUp()

	rpio.Close()

	// unfortunately the gpio watcher stuff doesnt work for me in this context, so we have to poll the button instead
	b.PushToTalkButton = gpio.NewInput(PushToTalkButtonPin)
	b.SelectButton = gpio.NewInput(SelectButtonPin)
	go func() {
		for {
			pttState, err := b.PushToTalkButton.Read()

			if pttState != b.PushToTalkButtonState && err == nil {
				b.PushToTalkButtonState = pttState

				if b.Stream != nil {
					if b.PushToTalkButtonState == 1 {
						fmt.Printf("PTT Button is released\n")
						b.TransmitStop()
					} else {
						fmt.Printf("PTT Button is pressed\n")
						b.TransmitStart()
					}
				}

			}

			selectState, err := b.SelectButton.Read()

			if selectState != b.SelectButtonState && err == nil {
				b.SelectButtonState = selectState

				if b.Stream != nil {
					if b.SelectButtonState == 1 {
						fmt.Printf("Select Button is released\n")
						// do nothing
					} else {
						fmt.Printf("Select Button is pressed\n")
						b.NextChannel()
					}
				}
			}

			time.Sleep(500 * time.Millisecond)
		}
	}()

	// then we can do our gpio stuff
	b.OnlineLED = gpio.NewOutput(OnlineLEDPin, false)
	b.ParticipantsLED = gpio.NewOutput(ParticipantsLEDPin, false)
	b.TransmitLED = gpio.NewOutput(TransmitLEDPin, false)
}

func (b *Talkiepi) LEDOn(LED gpio.Pin) {
	if b.GPIOEnabled == false {
		return
	}

	LED.High()
}

func (b *Talkiepi) LEDOff(LED gpio.Pin) {
	if b.GPIOEnabled == false {
		return
	}

	LED.Low()
}
