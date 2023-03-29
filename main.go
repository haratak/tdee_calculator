package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Gender int

const (
	Male Gender = iota
	Female
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	height, err := getInteger(reader, "Enter your height (in cm): ")
	if err != nil {
		fmt.Println(err)
		return
	}

	weight, err := getInteger(reader, "Enter your weight (in kg): ")
	if err != nil {
		fmt.Println(err)
		return
	}

	age, err := getInteger(reader, "Enter your age: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	gender, err := getGender(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	activityLevel, err := getActivityLevel(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	bmr := calculateBMR(height, weight, age, gender)
	tdee := calculateTDEE(bmr, activityLevel)

	fmt.Println("BMR:", bmr)
	fmt.Println("TDEE:", tdee)
}

func getInteger(reader *bufio.Reader, message string) (int, error) {
	fmt.Print(message)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalid input: %s", input)
	}

	return value, nil
}

func getGender(reader *bufio.Reader) (Gender, error) {
	fmt.Print("Enter your gender (M/F): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return Male, err
	}

	input = strings.TrimSpace(input)
	switch strings.ToLower(input) {
	case "m":
		return Male, nil
	case "f":
		return Female, nil
	default:
		return Male, fmt.Errorf("invalid gender:%s", input)
	}
}

func calculateTDEE(bmr float64, activityLevel float64) float64 {
	return bmr * activityLevel
}
func calculateBMR(height int, weight int, age int, gender Gender) float64 {
	var bmr float64

	if gender == Male {
		bmr = 88.362 + (13.397 * float64(weight)) + (4.799 * float64(height)) - (5.677 * float64(age))
	} else {
		bmr = 447.593 + (9.247 * float64(weight)) + (3.098 * float64(height)) - (4.33 * float64(age))
	}

	return bmr
}

func getActivityLevel(reader *bufio.Reader) (float64, error) {
	fmt.Println("Choose your activity level:")
	fmt.Println("(1) Sedentary (little or no exercise)")
	fmt.Println("(2) Lightly active (light exercise/sports 1-3 days per week)")
	fmt.Println("(3) Moderately active (moderate exercise/sports 3-5 days per week)")
	fmt.Println("(4) Very active (hard exercise/sports 6-7 days per week)")
	fmt.Println("(5) Super active (very hard exercise/sports or physical job)")

	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	switch input {
	case "1":
		return 1.2, nil
	case "2":
		return 1.375, nil
	case "3":
		return 1.55, nil
	case "4":
		return 1.725, nil
	case "5":
		return 1.9, nil
	default:
		return 0, fmt.Errorf("invalid activity level: %s", input)
	}
}
