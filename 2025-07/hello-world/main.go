package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "Language to use for the application")
	flag.Parse()

	// Here you can use the lang variable to set up your application
	println("Application started with language:", lang)

	// Additional application logic goes here
	// For example, you might load language-specific resources or configurations
	// based on the value of lang.
	switch lang {
	case "en":
		println("Hello, World!")
	case "es":
		println("¡Hola, mundo!")
	case "fr":
		println("Bonjour, le monde!")
	case "zh":
		println("你好，世界！ (Nǐ hǎo, shì jiè!)")
	case "hi":
		println("नमस्ते, दुनिया! (Namaste, duniya!)")
	default:
		println("Hello, World!")
	}

	// Simple print statement
	fmt.Println("Hello, Go!")

	// Variable declaration
	var message string
	message = "This is a message."
	fmt.Println(message)

	// Short variable declaration
	count := 5
	fmt.Printf("The count is: %d\n", count)

	// Basic if statement
	if count > 0 {
		fmt.Println("Count is positive.")
	} else {
		fmt.Println("Count is not positive.")
	}

	// Basic for loop
	for i := 0; i < 3; i++ {
		fmt.Printf("Iteration: %d\n", i)
	}

	loop := true
	loopCount := 0
	// Basic while-like loop

	for loop {
		fmt.Println("This is a while-like loop.")
		if loopCount >= 5 {
			loop = false // Break the loop after 5 iterations
		}
		loopCount++
	}

	// loop with new range over int
	for i := range 3 {
		fmt.Printf("Iteration: %d\n", i)
	}

	basicMap := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	basicMap["key3"] = "value3" // Adding a new key-value pair

	// Basic map usage
	for key, value := range basicMap {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
		if key == "key2" {
			fmt.Println("Found key2, value:", value)
		}
	}
}
