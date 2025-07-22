package output

import "github.com/fatih/color"

/*
This file contains colour mappings that are used throughout the app.
*/

// SafetyColors maps some basic colour names to the specific RGB values used by TfL on safety signs and notices. Because
// these describe basic often-used colours, we use the same values throughout the app for consistency.
var SafetyColors = map[string]*color.Color{
	"blue":   color.RGB(0, 92, 185),
	"red":    color.RGB(220, 36, 31),
	"yellow": color.RGB(255, 200, 10),
	"green":  color.RGB(0, 125, 50),
	"black":  color.RGB(0, 0, 0),
	"white":  color.RGB(255, 255, 255),
}
