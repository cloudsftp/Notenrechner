package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	g "github.com/AllenDang/giu"
)

var (
	maxPointsStr string = "100"
	pointsStr    string = "100"
	gradeStr     string = "1"
)

func gradeFromPoints(points float32, maxPoints float32) float32 {
	return 6 - 5*(points/maxPoints)
}

func roundGrade(grade float32) float32 {
	var ratio float32 = 4
	return float32(math.Round(float64(grade*ratio))) / ratio
}

func pointsFromText(text string) (float32, error) {
	if text == "" {
		return 0, nil
	}

	text = strings.Replace(text, ",", ".", 1)
	points, err := strconv.ParseFloat(text, 32)
	if err != nil {
		return 0, err
	}

	return float32(points), nil
}

func calculateGrade() {
	points, err := pointsFromText(pointsStr)
	if err != nil {
		gradeStr = fmt.Sprintf("Ungültige Eingabe (%s)", err)
		return
	}

	maxPoints, err := pointsFromText(maxPointsStr)
	if err != nil {
		gradeStr = fmt.Sprintf("Ungültige Eingabe (%s)", err)
		return
	}

	preciseGrade := gradeFromPoints(points, maxPoints)
	roundedGrade := roundGrade(preciseGrade)

	gradeStr, err = stringifyGrade(roundedGrade)
	if err != nil {
		gradeStr = fmt.Sprintf("Ungültige Eingabe (%s)", err)
		return
	}

	gradeStr = fmt.Sprintf("%s (Genau: %5.3f)", gradeStr, preciseGrade)
}

func stringifyGrade(roundedGrade float32) (string, error) {
	if roundedGrade < 1.0 || roundedGrade > 6.0 {
		return "", fmt.Errorf("Außerhalb des gültigen Bereichs [0, %s]: %f", maxPointsStr, roundedGrade)
	}

	integerGrade := math.Round(float64(roundedGrade))
	gradeString := fmt.Sprintf("%d", int(integerGrade))

	switch roundedGrade - float32(integerGrade) {
	case 0.25:
		gradeString = fmt.Sprintf("%d-", int(integerGrade))
	case -0.25:
		gradeString = fmt.Sprintf("%d+", int(integerGrade))
	case 0.5:
		gradeString = fmt.Sprintf("%d-%d", int(integerGrade), int(integerGrade+1.0))
	case -0.5:
		gradeString = fmt.Sprintf("%d-%d", int(integerGrade-1.0), int(integerGrade))
	}

	return gradeString, nil
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
	wnd := g.NewMasterWindow("Notenrechner", 400, 150, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}
