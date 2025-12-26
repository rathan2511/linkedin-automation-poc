package humanizer

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// RandomSleep waits for a random time between min and max milliseconds
func RandomSleep(min, max int) {
	n := rand.Intn(max-min) + min
	time.Sleep(time.Duration(n) * time.Millisecond)
}

// MoveMouseSmoothly moves the mouse using the raw browser protocol
func MoveMouseSmoothly(page *rod.Page, toX, toY float64) {
	// 1. Move slightly off-target (Random jitter)
	jitterX := float64(rand.Intn(20) - 10)
	jitterY := float64(rand.Intn(20) - 10)

	// We use the "proto" command directly
	_ = proto.InputDispatchMouseEvent{
		Type: proto.InputDispatchMouseEventTypeMouseMoved,
		X:    toX + jitterX,
		Y:    toY + jitterY,
	}.Call(page)

	// 2. Short pause
	RandomSleep(50, 150)

	// 3. Move to exact target
	_ = proto.InputDispatchMouseEvent{
		Type: proto.InputDispatchMouseEventTypeMouseMoved,
		X:    toX,
		Y:    toY,
	}.Call(page)
}

// TypeLikeHuman types text one character at a time with delays
func TypeLikeHuman(el *rod.Element, text string) {
	for _, char := range text {
		el.MustInput(string(char))
		time.Sleep(100 * time.Millisecond)
	}
}
