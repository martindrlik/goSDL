package main

import (
	"log"

	"github.com/martindrlik/goSDL/sdl"
)

func main() {
	defer sdl.Quit()
	if err := sdl.Init(sdl.Video | sdl.Events); err != nil {
		log.Fatalf("could not initialize subsystems: %v", err)
	}
	w, r := mustWR(sdl.CreateWindowAndRenderer(640, 480, sdl.Resizable))
	defer r.Destroy()
	defer w.Destroy()
	b := w.Brightness()
	defer func() { w.SetBrightness(b) }()
	w.SetBrightness(1.5)
	var ev sdl.Event
	points := []struct{ X, Y int32 }{
		{10, 20},
		{30, 40},
		{50, 60},
	}
	for {
		sdl.PollEvent(&ev)
		if ev.Type == sdl.EvQuit {
			break
		}
		check(r.SetRenderDrawColor(0, 0, 0, 255), "unable to set color for clear")
		check(r.RenderClear(), "unable to clear")
		check(r.SetRenderDrawColor(255, 255, 255, 255), "unable to set color for drawing point")
		check(r.RenderDrawPoint(320, 320), "unable to draw center point")
		check(r.SetRenderDrawColor(255, 0, 0, 255), "unable to set red color for drawing 3 points")
		check(r.RenderDrawPoints(points), "unable to draw 3 points")
		check(r.SetRenderDrawColor(0, 255, 0, 255), "unable to set green color for drawing a line")
		check(r.RenderDrawLine(350, 350, 400, 400), "unable to draw green line")
		r.RenderPresent()
	}
}

func mustWR(window *sdl.Window, renderer *sdl.Renderer, err error) (*sdl.Window, *sdl.Renderer) {
	if err != nil {
		log.Fatalf("could not create window and renderer: %v", err)
	}
	return window, renderer
}

func check(err error, em string) {
	if err != nil {
		log.Fatalf("%s: %v", em, err)
	}
}
