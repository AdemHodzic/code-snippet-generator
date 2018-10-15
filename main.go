package main

import (
	"code-snippet-generator/models/color"
	"code-snippet-generator/models/file"
	"code-snippet-generator/models/line"

	"log"
	"strconv"
	"strings"
	"time"

	"github.com/fogleman/gg"
)

var (
	width  float64  = 1024
	height float64  = 768
	files  []string = reader.GetFiles()
	colors          = color.GetColorConfig()
)

func main() {
	log.Println("Processing Started...")
	for _, file := range files {
		log.Printf("Processing %v...\n", file)
		makeImage(file)
	}
}

func makeImage(file string) {

	text := reader.ReadFile(file)
	context := gg.NewContext(int(width), int(height))
	drawWindow(context)
	drawButtons(context)
	writeToWindow(text, context)
	saveImage(context)
}

func drawWindow(context *gg.Context) {
	context.SetHexColor(colors.BackgroundColor)
	context.Clear()
	context.DrawRectangle(width*0.1, height*0.05, width*0.8, height*0.05)
	context.SetHexColor(colors.WindowTitle)
	context.Fill()
	context.DrawRectangle(width*0.1, height*0.1, width*0.8, height*0.8)
	context.SetHexColor(colors.WindowBackground)
	context.Fill()
}

func drawButtons(context *gg.Context) {

	context.SetRGB255(242, 31, 31)
	context.DrawCircle(width*0.87, height*0.07, 7)
	context.Fill()

	context.SetRGB255(162, 255, 109)
	context.DrawCircle(width*0.82, height*0.07, 7)
	context.Fill()

	context.SetRGB255(249, 255, 84)
	context.DrawCircle(width*0.77, height*0.07, 7)
	context.Fill()
}

func writeToWindow(text []string, context *gg.Context) {

	context.SetHexColor(colors.TextColor)
	context.LoadFontFace("Roboto-Black.ttf", 18)
	for index, text := range text {
		x := width*0.2 + countTabs(text)*0.02*width
		y := height*0.2 + (height*0.05)*float64(index)
		text = strings.Replace(text, "\t", "", -1)
		var code line.Line
		code.Print(context, text, x, y)
	}
}

func countTabs(line string) float64 {

	return float64(strings.Count(line, "\t"))
}

func saveImage(context *gg.Context) {
	time.Sleep(1000 * time.Millisecond)
	now := strconv.FormatInt(time.Now().Unix(), 10)
	name := "dist/" + now + ".png"
	context.SavePNG(name)
}
