package main

import (
	"fmt"
	"math"
	"strconv"

	g "github.com/AllenDang/giu"
)

var (
	maxPointsStr string = "100"
	pointsStr    string = "100"
	gradeStr     string = "1.00"
)

func gradeFromPoints(points int, maxPoints int) float32 {
	return 6 - 5*(float32(points)/float32(maxPoints))
}

func roundGrade(grade float32) float32 {
	var ratio float32 = 4
	return float32(math.Round(float64(grade*ratio))) / ratio
}

func pointsFromText(text string) int {
	if text == "" {
		return 0
	}

	points, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}

	return points
}

func calculateGrade() {
	points := pointsFromText(pointsStr)
	maxPoints := pointsFromText(maxPointsStr)

	preciseGrade := gradeFromPoints(points, maxPoints)
	roundedGrade := roundGrade(preciseGrade)

	gradeStr = fmt.Sprintf("%4.2f", roundedGrade)
	if preciseGrade != roundedGrade {
		gradeStr = fmt.Sprintf("%s (Genau: %4.2f)", gradeStr, preciseGrade)
	}
}

func loop() {
	g.SingleWindow().Layout(
		g.Label("Notenrechner"),
		g.Spacing(),
		g.InputText(&pointsStr).Size(100).
			Label("Erreichte Puktzahl").OnChange(calculateGrade),
		g.InputText(&maxPointsStr).Size(100).
			Label("Maximal erreichbare Puktzahl").OnChange(calculateGrade),
		g.Spacing(),
		g.Row(
			g.Button("Note berechnen").OnClick(calculateGrade),
			g.Label(gradeStr),
		),
	)
}

func main() {
	wnd := g.NewMasterWindow("Notenrechner", 250, 100, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}
