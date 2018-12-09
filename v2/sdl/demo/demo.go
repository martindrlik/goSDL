package main

import (
	"image/color"
	"log"

	"github.com/martindrlik/v2/sdl"
	"github.com/martindrlik/v2/sdl/events"
	"github.com/martindrlik/v2/sdl/video"
)

func main() {
	defer sdl.Quit()
	var (
		err      error
		window   *video.Window
		renderer *video.Renderer
		ev       events.Event
	)
	if err = sdl.Init(); err != nil {
		log.Fatalf("demo: unable initialize: %v", err)
	}
	if window, renderer, err = video.CreateWindowAndRenderer(
		640,
		640,
		video.Resizable); err != nil {
		log.Fatalf("demo: unable create window and renderer: %v", err)
	}
	defer renderer.Destroy()
	defer window.Destroy()
	for {
		events.PollEvent(&ev)
		if ev.Type == events.Quit {
			break
		}
		if err = renderer.SetRenderDrawColor(color.Black); err != nil {
			log.Fatalf("demo: unable to set color for clear: %v", err)
		}
		if err = renderer.RenderClear(); err != nil {
			log.Fatalf("demo: unable clear renderer: %v", err)
		}
		if err = renderer.SetRenderDrawColor(color.White); err != nil {
			log.Fatalf("demo: unable to set color for drawing point: %v", err)
		}
		if err = renderer.RenderDrawPoint(320, 320); err != nil {
			log.Fatalf("demo: unable to draw center point: %v", err)
		}
		renderer.RenderPresent()
	}
}
