package main

import (
	"log"

	"github.com/martindrlik/goSDL/v2/events"
	"github.com/martindrlik/goSDL/v2/sdl"
	"github.com/martindrlik/goSDL/v2/video"
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
		log.Fatalf("demo: unable to initialize: %v", err)
	}
	if window, renderer, err = video.CreateWindowAndRenderer(
		640,
		480,
		video.Resizable); err != nil {
		log.Fatalf("demo: unable to create window and renderer: %v", err)
	}
	defer renderer.Destroy()
	defer window.Destroy()
	for {
		events.PollEvent(&ev)
		if ev.Type == events.Quit {
			break
		}
		if err = renderer.SetRenderDrawColor(0, 0, 0, 255); err != nil {
			log.Fatalf("demo: unable to set color for clear: %v", err)
		}
		if err = renderer.RenderClear(); err != nil {
			log.Fatalf("demo: unable to clear renderer: %v", err)
		}
		if err = renderer.SetRenderDrawColor(255, 255, 255, 255); err != nil {
			log.Fatalf("demo: unable to set color for drawing point: %v", err)
		}
		if err = renderer.RenderDrawPoint(320, 320); err != nil {
			log.Fatalf("demo: unable to draw center point: %v", err)
		}
		if err = renderer.SetRenderDrawColor(255, 0, 0, 255); err != nil {
			log.Fatalf("demo: unable to set red color for drawing 3 points")
		}
		if err = renderer.RenderDrawPoints([]struct{ X, Y int32 }{
			{10, 20},
			{30, 40},
			{50, 60},
		}); err != nil {
			log.Fatalf("demo: unable to draw 3 points: %v", err)
		}
		if err = renderer.SetRenderDrawColor(0, 255, 0, 255); err != nil {
			log.Fatalf("demo: unable to set green color for drawing a line: %v", err)
		}
		if err = renderer.RenderDrawLine(350, 350, 400, 400); err != nil {
			log.Fatalf("demo: unable to draw green line: %v", err)
		}
		renderer.RenderPresent()
	}
}
