package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	engine := html.New("./views", ".html")

	
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	
	app.Get("/test_printer", TestAllPrinter)

	
	log.Fatal(app.Listen(":8080"))
}

func TestAllPrinter(c *fiber.Ctx) error {
	printer_1 := "http://10.4.7.201"
	printer_2 := "http://10.4.7.203"

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp_1, err1 := client.Get(printer_1)
	if err1 != nil {
		log.Printf("Error fetching %s: %v", printer_1, err1)
	}
	defer closeResponse(resp_1)

	resp_2, err2 := client.Get(printer_2)
	if err2 != nil {
		log.Printf("Error fetching %s: %v", printer_2, err2)
	}
	defer closeResponse(resp_2)

	
	return c.Render("index", fiber.Map{
		"status_printer_1": getStatus(resp_1),
		"status_printer_2": getStatus(resp_2),
	})
}

func getStatus(resp *http.Response) int {
	if resp == nil {
		return 0 // 0 หมายถึง ไม่สามารถเชื่อมต่อได้
	}
	return resp.StatusCode
}

func closeResponse(resp *http.Response) {
	if resp != nil {
		resp.Body.Close()
	}
}
