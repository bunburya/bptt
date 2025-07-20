package tfl

import "github.com/fatih/color"

var lineColors = map[string]*color.Color{
	"bakerloo":         color.BgRGB(178, 99, 0),
	"central":          color.BgRGB(220, 36, 31),
	"circle":           color.BgRGB(255, 200, 10),
	"district":         color.BgRGB(0, 125, 50),
	"hammersmith-city": color.BgRGB(245, 137, 166),
	"jubilee":          color.BgRGB(131, 141, 147),
	"metropolitan":     color.BgRGB(155, 0, 88),
	"northern":         color.BgRGB(0, 0, 0),
	"piccadilly":       color.BgRGB(0, 25, 168),
	"victoria":         color.BgRGB(3, 155, 229),
	"waterloo-city":    color.BgRGB(118, 208, 189),
	"liberty":          color.BgRGB(93, 96, 97),
	"lioness":          color.BgRGB(250, 166, 26),
	"mildmay":          color.BgRGB(0, 119, 173),
	"suffragette":      color.BgRGB(91, 189, 114),
	"weaver":           color.BgRGB(130, 58, 98),
	"windrush":         color.BgRGB(237, 27, 0),
	"sl1":              color.BgRGB(228, 59, 23),
	"sl2":              color.BgRGB(187, 204, 0),
	"sl3":              color.BgRGB(129, 27, 109),
	"sl4":              color.BgRGB(91, 91, 90),
	"sl5":              color.BgRGB(55, 171, 221),
	"sl6":              color.BgRGB(225, 0, 122),
	"sl7":              color.BgRGB(190, 0, 94),
	"sl8":              color.BgRGB(17, 52, 131),
	"sl9":              color.BgRGB(5, 142, 156),
	"sl10":             color.BgRGB(242, 149, 0),
}

var modeColors = map[string]*color.Color{
	"tfl":            color.BgRGB(0, 25, 168),
	"dlr":            color.BgRGB(0, 175, 173),
	"elizabeth-line": color.BgRGB(96, 57, 158),
	"bus":            color.BgRGB(220, 36, 31),
	"cable-car":      color.BgRGB(115, 79, 160),
	"coach":          color.BgRGB(255, 166, 0),
	"overground":     color.BgRGB(250, 123, 5),
	"river-bus":      color.BgRGB(3, 155, 229),
	"tram":           color.BgRGB(95, 181, 38),
	"tube":           color.BgRGB(0, 25, 168),
	"cycle-hire":     color.BgRGB(236, 0, 0),
}

var safetyColors = map[string]*color.Color{
	"blue":   color.RGB(0, 92, 185),
	"red":    color.RGB(220, 36, 31),
	"yellow": color.RGB(255, 200, 10),
	"green":  color.RGB(0, 125, 50),
	"black":  color.RGB(0, 0, 0),
	"white":  color.RGB(255, 255, 255),
}
