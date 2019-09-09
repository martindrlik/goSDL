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
	w, re := mustWR(sdl.CreateWindowAndRenderer(640, 480, sdl.Resizable))
	defer re.Destroy()
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
		check(re.SetDrawColor(0, 0, 0, 255), "unable to set color for clear")
		check(re.Clear(), "unable to clear")
		check(re.SetDrawColor(255, 255, 255, 255), "unable to set color for drawing point")
		check(re.DrawPoint(320, 320), "unable to draw center point")
		check(re.SetDrawColor(255, 0, 0, 255), "unable to set red color for drawing 3 points")
		check(re.DrawPoints(points), "unable to draw 3 points")
		check(re.SetDrawColor(0, 255, 0, 255), "unable to set green color for drawing a line")
		check(re.DrawLine(350, 350, 400, 400), "unable to draw green line")
		re.Present()
	}
}

func mustWR(w *sdl.Window, re *sdl.Renderer, err error) (*sdl.Window, *sdl.Renderer) {
	if err != nil {
		log.Fatalf("could not create window and renderer: %v", err)
	}
	return w, re
}

func check(err error, em string) {
	if err != nil {
		log.Fatalf("%s: %v", em, err)
	}
}
