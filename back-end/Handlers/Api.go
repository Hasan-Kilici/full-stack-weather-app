package Handlers

import (
	"github.com/gofiber/fiber/v2"
	"weather/Utils"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
)

func HelloHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"hello": "world",
	})
}

func Query(c *fiber.Ctx) error {
	Key := Utils.GetConfig("api", "key")
	City := c.Params("City")
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?units=metric&appid="+Key+"&q="+City+"&lang=tr")
	if err != nil {
	   log.Fatalln(err)
	}
 
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	   log.Fatalln(err)
	}
	
	var Data map[string]interface{}
	json.Unmarshal(body, &Data)

	c.JSON(fiber.Map{
		"data":Data,
	})
	
	return nil
}