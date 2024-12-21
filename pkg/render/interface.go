package render

import "time"

type TerminalRender interface {
	RenderContent(duration time.Duration)
	RenderFlashing()
	CanRender(duration time.Duration) error
	RenderHeight() int
}
