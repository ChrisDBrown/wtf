package wtf

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const offscreen = -1000
const modalWidth = 80
const modalHeight = 22

// NewBillboardModal creates and returns a billboard modal
// These modals are suitable for displaying static text as a fixed overlay
func NewBillboardModal(text string, closeFunc func()) *tview.Frame {
	keyboardIntercept := func(event *tcell.EventKey) *tcell.EventKey {
		if string(event.Rune()) == "/" {
			closeFunc()
			return nil
		}

		switch event.Key() {
		case tcell.KeyEsc:
			closeFunc()
			return nil
		case tcell.KeyTab:
			return nil
		default:
			return event
		}
	}

	textView := tview.NewTextView()
	textView.SetDynamicColors(true)
	textView.SetInputCapture(keyboardIntercept)
	textView.SetText(text)
	textView.SetWrap(true)

	frame := tview.NewFrame(textView)
	frame.SetRect(offscreen, offscreen, modalWidth, modalHeight)

	drawFunc := func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		w, h := screen.Size()
		frame.SetRect((w/2)-(width/2), (h/2)-(height/2), width, height)
		return x, y, width, height
	}

	frame.SetBorder(true)
	frame.SetBorders(1, 1, 0, 0, 1, 1)
	frame.SetDrawFunc(drawFunc)

	return frame
}
