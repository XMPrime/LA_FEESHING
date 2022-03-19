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
	go printMousePos()
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

			// if fishing, real in fish. if not fishing, cast line
			if isFeeshing {
				// check for "!"
				exPixelCount := 0
				var pixelColors [9]string

				// grabs pixel colors in a 3x3 pixel area
				for i := 0; i <= 3; i++ {
					for j := 0; j <= 3; j++ {
						pixelColors[i] = robotgo.GetPixelColor(959+i, 499+j)
					}
				}
				// count pixels that are "yellow"
				for _, pixelColor := range pixelColors {
					totalFs := strings.Count(pixelColor, "f")
					if string(pixelColor[0]) == "f" && totalFs <= 4 {
						exPixelCount++
					}
				}

				isEx := exPixelCount >= 4
				// reel in feesh
				if isEx {
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
			// mouseX, mouseY := robotgo.GetMousePos();
			robotgo.KeyPress("d")
			time.Sleep(time.Second * 2)
			isBaiting = false
			time.Sleep(time.Second * 901)
		}
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func printMousePos() {
	for {
		time.Sleep(time.Millisecond * 500)
		mouseX, mouseY := robotgo.GetMousePos()
		fmt.Println(mouseX, mouseY)
	}

}
