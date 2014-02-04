package main

import (
	"fmt"
	"math"
)

func main() {
	w1 := climateEdmonton.Warmest()
	fmt.Printf("Highest temperature was %v in %d.\n", w1.max, w1.year)

	w2 := climateEdmonton.Coldest()
	fmt.Printf("Lowest temperature was %v in %d.\n", w2.min, w2.year)
}

// Weather holds the maximum and minimum temperatures for a given day
type Weather struct {
	year     int
	max, min Celsius
}

// Climate holds the weather across many days
type Climate []Weather

// January 27th temperatures in Edmonton, Alberta
// data downloaded from http://edmonton.weatherstats.ca/download.html
var climateEdmonton = Climate{
	{1988, 0.2, -5.9},
	{1989, 10.7, -3.7},
	{1990, -11.7, -17.5},
	{1991, -0.7, -15},
	{1992, 4.7, -6.3},
	{1994, 0.5, -11.5},
	{1995, -2.6, -14.6},
	{1996, -25.1, -32.1},
	{1997, -22, -33.4},
	{1998, 3.6, -7.3},
	{1999, -6.6, -17.1},
	{2000, 0.8, -8.8},
	{2001, 4.7, -6},
	{2002, -19.8, -27.8},
	{2003, -7.8, -17.3},
	{2004, -30.6, -35.5},
	{2005, -1, -5.6},
	{2006, -2.1, -9},
	{2007, -3.5, -17.6},
	{2008, -4.3, -27.7},
	{2009, 2.3, -15.2},
	{2010, -10.4, -21.5},
	{2011, 8.7, -2.8},
	{2012, 0.2, -7.6},
	{2013, -2.2, -16.7},
}

// Warmest finds the warmest weather
func (climate Climate) Warmest() Weather {
	warmest := Weather{max: -math.MaxFloat64}

	for _, w := range climate {
		warmest = Warmer(warmest, w)
	}
	return warmest
}

// Warmer returns the warmer of two weather datums
func Warmer(weather1, weather2 Weather) Weather {
	if weather2.max > weather1.max {
		return weather2
	}
	return weather1
}

// Coldest finds the coldest weather
func (climate Climate) Coldest() Weather {
	coldest := Weather{min: math.MaxFloat64}

	for _, w := range climate {
		coldest = Colder(coldest, w)
	}
	return coldest
}

// Colder returns the colder of two weather datums
func Colder(weather1, weather2 Weather) Weather {
	if weather2.min < weather1.min {
		return weather2
	}
	return weather1
}

// Celsius represents a temperature
type Celsius float64

// String returns a pretty printed temperature
func (c Celsius) String() string {
	return fmt.Sprintf("%.2fºC", c)
}
