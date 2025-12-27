package humanizer

import (
	"math"
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func RandomSleep(min, max int) {
	n := rand.Intn(max-min) + min
	time.Sleep(time.Duration(n) * time.Millisecond)
}

// TypeLikeHuman now has a 'safeMode' to avoid typos in passwords/emails
func TypeLikeHuman(el *rod.Element, text string, safeMode bool) {
	page := el.Page()

	// Ensure element has focus
	el.MustClick()

	for _, char := range text {
		// Optional typo simulation (disabled in safeMode)
		if !safeMode && rand.Float64() < 0.04 {
			_ = proto.InputInsertText{
				Text: "q",
			}.Call(page)
			RandomSleep(100, 250)

			_ = proto.InputDispatchKeyEvent{
				Type: proto.InputDispatchKeyEventTypeKeyDown,
				Key:  "Backspace",
			}.Call(page)
			RandomSleep(50, 150)
		}

		// Insert actual character (handles @ . _ - correctly)
		_ = proto.InputInsertText{
			Text: string(char),
		}.Call(page)

		RandomSleep(80, 250)
	}
}

func MoveMouseBezier(page *rod.Page, startX, startY, endX, endY float64) {
	cp1x := startX + (endX-startX)*rand.Float64()
	cp1y := startY + (endY-startY)*rand.Float64()
	cp2x := startX + (endX-startX)*rand.Float64()
	cp2y := startY + (endY-startY)*rand.Float64()

	steps := 15 + rand.Intn(10)
	for i := 0; i <= steps; i++ {
		t := float64(i) / float64(steps)
		x := math.Pow(1-t, 3)*startX + 3*math.Pow(1-t, 2)*t*cp1x + 3*(1-t)*math.Pow(t, 2)*cp2x + math.Pow(t, 3)*endX
		y := math.Pow(1-t, 3)*startY + 3*math.Pow(1-t, 2)*t*cp1y + 3*(1-t)*math.Pow(t, 2)*cp2y + math.Pow(t, 3)*endY

		_ = proto.InputDispatchMouseEvent{
			Type: proto.InputDispatchMouseEventTypeMouseMoved,
			X:    x,
			Y:    y,
		}.Call(page)
		time.Sleep(time.Duration(10+rand.Intn(15)) * time.Millisecond)
	}
}

func HumanScroll(page *rod.Page) {
	dist := float64(400 + rand.Intn(300))
	page.Mouse.MustScroll(0, dist)
	RandomSleep(1000, 2000)
}

func HoverElement(page *rod.Page, el *rod.Element) {
	shape, err := el.Shape()
	if err != nil {
		return
	}
	box := shape.Box()
	MoveMouseBezier(page, 0, 0, box.X+(box.Width/2), box.Y+(box.Height/2))
	RandomSleep(400, 800)
}

func IsBusinessHours() bool {
	h := time.Now().Hour()
	return h >= 9 && h <= 18
}
