package api

import (
	"encoding/json"
	"fiber-starter/calculator"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Greet(c *fiber.Ctx) error {
	return c.SendString("Hello, world!")
}

func AddHandler(c *fiber.Ctx) error {
	var body Payload

	err := json.Unmarshal(c.Body(), &body)

	if err != nil {
		log.Println("Error: ", err.Error())
		return err
	}

	calc := calculator.NewTheBestCalculator("MyC")
	result := calc.Add(body.X, body.Y)
	return c.JSON(ResponseBody{
		Input:  body,
		Output: result,
	})
}

func SubtractHandler(c *fiber.Ctx) error {
	var body Payload

	err := json.Unmarshal(c.Body(), &body)

	if err != nil {
		log.Println("Error: ", err.Error())
		return err
	}

	calc := calculator.NewTheBestCalculator("MyC")
	result := calc.Subtract(body.X, body.Y)
	return c.JSON(ResponseBody{
		Input:  body,
		Output: result,
	})
}

func MultiplyHandler(c *fiber.Ctx) error {
	var body Payload

	err := json.Unmarshal(c.Body(), &body)

	if err != nil {
		log.Println("Error: ", err.Error())
		return err
	}

	calc := calculator.NewTheBestCalculator("MyC")
	result := calc.Multiply(body.X, body.Y)
	return c.JSON(ResponseBody{
		Input:  body,
		Output: result,
	})
}

func DivideHandler(c *fiber.Ctx) error {
	var body Payload

	err := json.Unmarshal(c.Body(), &body)

	if err != nil {
		log.Println("Error: ", err.Error())
		return err
	}

	calc := calculator.NewTheBestCalculator("MyC")
	result := calc.Divide(body.X, body.Y)
	return c.JSON(ResponseBody{
		Input:  body,
		Output: result,
	})
}

func IsOddHandler(c *fiber.Ctx) error {
	var body Payload

	err := json.Unmarshal(c.Body(), &body)

	if err != nil {
		log.Println("Error: ", err.Error())
		return err
	}

	calc := calculator.NewTheBestCalculator("MyC")
	result := calc.IsOdd(body.X)
	return c.JSON(ResponseBodyIsOdd{
		Output: result,
	})
}
