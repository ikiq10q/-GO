package main

import (
	"fmt"
	"math"
)

func main() {
	//Zadanie 1
	bannerWidth := 12
	bannerHeight := 8
	bannerArea := bannerWidth * bannerHeight
	fmt.Println("Size:", bannerArea)

	halfBannerArea := bannerArea / 2
	fmt.Println("halfBannerArea:", halfBannerArea)

	bannerBorderLength := (bannerWidth + bannerHeight) * 2
	fmt.Println("bannerBorderLength:", bannerBorderLength)

	//Zadanie 2
	boxCount := 29
	leftoverBoxes := boxCount % 5

	fmt.Println("leftoverBoxes:", leftoverBoxes)

	//Zadanie 3
	tempMorning := 15
	tempAfternoon := 20
	tempEvening := 30
	totalTemp := tempMorning + tempAfternoon + tempEvening
	averageTemp := float64(totalTemp) / 3

	fmt.Printf("averageTemp: %.2f°C\n", averageTemp)

	//zadanie 4
	knownWords := 47
	wordsGoal := 120
	progressPercent := (float64(knownWords) / float64(wordsGoal)) * 100

	fmt.Printf("progressPercent: %.2f%%\n", progressPercent)

	//zadanie 5
	coins := 0
	fmt.Println("Coins Score:", coins)

	coins += 500
	fmt.Println("Coins Score:", coins)

	coins += 1200
	fmt.Println("Coins Score:", coins)

	coins /= 2
	fmt.Println("Coins Score:", coins)

	coins *= 2
	fmt.Println("Coins Score:", coins)

	coins -= 300
	fmt.Println("Coins Score:", coins)

	//zadanie 6
	participants := 42
	groupCount := 8
	participantsPerGroup := participants / groupCount

	fmt.Println("ParticipantsPerGroup", participantsPerGroup)

	//zadanie 7 расчитывается по математическому "Порядку действий"
	fmt.Println(float64(20 - 4*3))
	fmt.Println(float64(20-4) * 3)

	//zadanie 8
	squareValue := 81
	result := math.Sqrt(float64(squareValue))

	multiplier := 5
	exponent := 2
	result_2 := math.Pow(float64(multiplier), float64(exponent))

	fmt.Println(result, result_2)

}
