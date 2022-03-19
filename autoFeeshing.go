package main

import (
  "fmt"
  "math/rand"
  "strings"
  "time"
  "github.com/go-vgo/robotgo"
)

var isFeeshing bool = false;
var isBaiting bool = false;

func main() {
  go printMousePos();
  go feesh();
  fmt.Scanln();
}

func feesh() {
  // curMousePosX, curMousePosy := robotgo.GetMousePos();
  // resolution 1920 x 1080
  // need to randomly move the mouse around the ocean
  // press E when pixel turns 'yellow'
  // detect yellow by sampling pixels
  // exit when "x" is pressed

  // randomMouseMove
  isFeeshing := true;

  // move mouse to ocean
  // Press "e"

  for {
    rand.Seed(time.Now().UnixNano())

    exPixelCount := 0;
    pixelColors := [9]string;

    // grabs pixel colors in a 3x3 pixel area
    for i := 0; i <= 3; i++ {
      for j := 0; j <= 3; j++ {
        pixelColors[i] = robotgo.GetPixelColor(959 + i, 499 + j);
      }
    }
    // count pixels that are "yellow"
    for _, pixelColor := range pixelColors {
      if strings.Count(pixelColor, "f") <= 4 {
        exPixelCount++;
      }
    }

    isEx := exPixelCount >= 4;

    // cast line
    robotgo.KeyPress("e");

    // wait for feesh to bite



    // reel in feesh
    if isEx == true {
      // add random wait time
      fmt.Println("SUCCESS! " + pixelColor);
      randomReactionTime := time.Duration(rand.Intn(240 - 180 + 1) + 180);
      time.Sleep(randomReactionTime * time.Millisecond);
      robotgo.KeyPress("e");
      time.Sleep(5 * time.Second);
      isFeeshing = false;
    }
  }
}

func throwBait() {
  for {
    if isFeeshing == true {
      time.Sleep(time.Second * 10);
    } else {
      time.Sleep(time.Second * 903);
      mouseX, mouseY := robotgo.GetMousePos();
      robotgo.KeyPress("d");
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
    time.Sleep(time.Millisecond * 500);
    mouseX, mouseY := robotgo.GetMousePos();
    fmt.Println(mouseX, mouseY);
  }

}