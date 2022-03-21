package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

var isFeeshing bool = false
var isBaiting bool = false

func main() {
	go throwBait()
	go goFeesh()
	fmt.Scanln()
}

func goFeesh() {
	// curMousePosX, curMousePosy := robotgo.GetMousePos();
	// resolution 1920 x 1080
	// need to randomly move the mouse around the ocean
	// press E when pixel turns 'yellow'
	// detect yellow by sampling pixels
	// exit when "x" is pressed

	// randomMouseMove

	// move mouse to ocean
	// Press "e"

	for {
		if !isBaiting {

			// if fishing, reel in fish. if not fishing, cast line
			if isFeeshing {
				// check for "!"
				exPixelCount := 0
				var pixelColors []string

				// grabs pixel colors in a 3x3 pixel area
				for i := 0; i < 3; i++ {
					for j := 0; j < 3; j++ {
						pixelColors = append(pixelColors, robotgo.GetPixelColor(959+i, 499+j))
					}
				}
				// count pixels that are "yellow"
				for _, pixelColor := range pixelColors {
					totalFs := strings.Count(pixelColor, "f")
					if string(pixelColor[0]) == "f" && totalFs <= 4 {
						exPixelCount++
					}
				}

				isEx := exPixelCount >= 5
				// reel in feesh
				if isEx {
					fmt.Println(pixelColors)
					rand.Seed(time.Now().UnixNano())
					randomReactionTime := time.Duration(rand.Intn(240-180+1) + 180)
					time.Sleep(randomReactionTime * time.Millisecond) // random "human" reaction time
					robotgo.KeyPress("e")                             // reel in fish
					time.Sleep(5 * time.Second)                       // wait for fish to enter inventory and give throwBait a chance to run
					isFeeshing = false
				}
			} else {
				// move mouse to ocean
				robotgo.KeyPress("e")       // cast fish line
				time.Sleep(2 * time.Second) // wait for screen to center on bobber
				isFeeshing = true
			}

		}

	}
}

func throwBait() {
	for {
		if isFeeshing {
			time.Sleep(time.Second * 10)
		} else {
			isBaiting = true
			robotgo.KeyPress("d")
			time.Sleep(time.Second * 1)
			isBaiting = false
			time.Sleep(time.Second * 901)
		}
	}
}
