package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/apocelips/ascii-count-down/pkg/render"
	"github.com/apocelips/ascii-count-down/pkg/util"
)

func main() {
	durationArg := flag.String("duration", "", "set the count down duration, same format as time.ParseDuration")
	untilArg := flag.String("until", "", "wait until, the format is `HH:MM:SS`")
	typeArg := flag.String("t", "sharp", `"sharp" or "asciiart", default is "sharp"`)

	flag.Parse()

	if len(*durationArg) == 0 && len(*untilArg) == 0 {
		panic("need set --duration/--until")
	}

	if len(*durationArg) != 0 && len(*untilArg) != 0 {
		panic("only one option can be set")
	}

	var terminalRender render.TerminalRender
	switch *typeArg {
	case "sharp":
		terminalRender = render.NewSharpCharRender()
	case "asciiart":
		terminalRender = render.NewASCIIArtCharRender()
	default:
		panic("unknown type: " + *typeArg)
	}
	if err := terminalRender.CanRender(); err != nil {
		panic(err)
	}

	now := time.Now()
	var until time.Time
	if len(*untilArg) != 0 {
		var err error
		until, err = time.ParseInLocation("2006-01-02 15:04:05", util.GetToday()+" "+*untilArg, time.Local)
		if err != nil {
			panic(err)
		}
		if until.Before(now) {
			until = until.Add(24 * time.Hour)
		}
	}
	if len(*durationArg) != 0 {
		var err error
		var duration time.Duration
		duration, err = time.ParseDuration(*durationArg)
		if err != nil {
			panic(err)
		}
		if int(duration.Hours()) > 99 {
			panic("supports at most 99h59m59s")
		}
		until = now.Add(duration)
	}

	flashing := false
	started := false
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(500 * time.Millisecond)

	fmt.Printf("距离 %s 还有: \n\n", until.Format("2006-01-02 15:04:05"))
	for now.Before(until) {
		if started {
			util.CursorUp(terminalRender.RenderHeight())
		}

		started = true
		if flashing {
			terminalRender.RenderFlashing()
		} else {
			duration := until.Sub(now).Truncate(time.Second)
			terminalRender.RenderContent(duration)
		}

		flashing = !flashing
		now = time.Now()
		select {
		case <-sigs:
			util.CursorUp(terminalRender.RenderHeight())
			terminalRender.RenderContent(0)
			util.PrintlnRed("\ncount down canceled")
			os.Exit(0)
		case <-ticker.C:
		}
	}
}
