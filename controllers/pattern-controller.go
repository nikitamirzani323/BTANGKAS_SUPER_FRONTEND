package controllers

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

type Responsepattern struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Record      interface{} `json:"record"`
	Perpage     int         `json:"perpage"`
	Totalrecord int         `json:"totalrecord"`
	Totallose   int         `json:"totallose"`
	Totalwin    int         `json:"totalwin"`
	Time        string      `json:"time"`
}

func Patternhome(c *fiber.Ctx) error {
	type payload_patternhome struct {
		Pattern_search string `json:"pattern_search"`
		Pattern_page   int    `json:"pattern_page"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_patternhome)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(Responsepattern{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"pattern_search":  client.Pattern_search,
			"pattern_page":    client.Pattern_page,
		}).
		Post(PATH + "api/pattern")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*Responsepattern)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":      result.Status,
			"perpage":     result.Perpage,
			"totalrecord": result.Totalrecord,
			"totallose":   result.Totallose,
			"totalwin":    result.Totalwin,
			"message":     result.Message,
			"record":      result.Record,
			"time":        time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func PatternSave(c *fiber.Ctx) error {
	type payload_patternsave struct {
		Page                  string          `json:"page"`
		Sdata                 string          `json:"sdata" `
		Pattern_search        string          `json:"pattern_search" `
		Pattern_page          int             `json:"pattern_page" `
		Pattern_id            string          `json:"pattern_id" `
		Pattern_resultcardwin string          `json:"pattern_resultcardwin" `
		Data                  json.RawMessage `json:"data"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_patternsave)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	log.Println("Array: ", string(client.Data))
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":       hostname,
			"page":                  client.Page,
			"sdata":                 client.Sdata,
			"pattern_search":        client.Pattern_search,
			"pattern_page":          client.Pattern_page,
			"pattern_id":            client.Pattern_id,
			"pattern_resultcardwin": client.Pattern_resultcardwin,
			"pattern_list":          string(client.Data),
		}).
		Post(PATH + "api/patternsave")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
