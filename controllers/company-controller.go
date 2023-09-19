package controllers

import (
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/wl_super_backend_frontend/entities"
)

type responsecompany struct {
	Status   int         `json:"status"`
	Message  string      `json:"message"`
	Listcurr interface{} `json:"listcurr"`
	Record   interface{} `json:"record"`
}
type responsecompanyadminrule struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Listcompany interface{} `json:"listcompany"`
	Record      interface{} `json:"record"`
}

func Companyhome(c *fiber.Ctx) error {
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(entities.Home)
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
		SetResult(responsecompany{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"page":            client.Page,
		}).
		Post(PATH + "api/company")
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
	result := resp.Result().(*responsecompany)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":   result.Status,
			"message":  result.Message,
			"record":   result.Record,
			"listcurr": result.Listcurr,
			"time":     time.Since(render_page).String(),
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
func Companyadminrulehome(c *fiber.Ctx) error {
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(entities.Home)
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
		SetResult(responsecompanyadminrule{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"page":            client.Page,
		}).
		Post(PATH + "api/companyadminrule")
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
	result := resp.Result().(*responsecompanyadminrule)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":      result.Status,
			"message":     result.Message,
			"record":      result.Record,
			"listcompany": result.Listcompany,
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
func CompanySave(c *fiber.Ctx) error {
	type payload_companysave struct {
		Page               string `json:"page"`
		Sdata              string `json:"sdata" `
		Company_id         string `json:"company_id"`
		Company_name       string `json:"company_name" `
		Company_idcurr     string `json:"company_idcurr" `
		Company_nmowner    string `json:"company_nmowner" `
		Company_phoneowner string `json:"company_phoneowner" `
		Company_emailowner string `json:"company_emailowner" `
		Company_url        string `json:"company_url" `
		Company_status     string `json:"company_status" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_companysave)
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
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":    hostname,
			"page":               client.Page,
			"sdata":              client.Sdata,
			"company_id":         client.Company_id,
			"company_name":       client.Company_name,
			"company_idcurr":     client.Company_idcurr,
			"company_nmowner":    client.Company_nmowner,
			"company_phoneowner": client.Company_phoneowner,
			"company_emailowner": client.Company_emailowner,
			"company_url":        client.Company_url,
			"company_status":     client.Company_status,
		}).
		Post(PATH + "api/companysave")
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
func CompanyadminruleSave(c *fiber.Ctx) error {
	type payload_companyadminrulesave struct {
		Page                       string `json:"page"`
		Sdata                      string `json:"sdata" `
		Companyadminrule_id        string `json:"companyadminrule_id"`
		Companyadminrule_idcompany string `json:"companyadminrule_idcompany" `
		Companyadminrule_nmrule    string `json:"companyadminrule_nmrule" `
		Companyadminrule_rule      string `json:"companyadminrule_rule" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_companyadminrulesave)
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
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":            hostname,
			"page":                       client.Page,
			"sdata":                      client.Sdata,
			"companyadminrule_id":        client.Companyadminrule_id,
			"companyadminrule_idcompany": client.Companyadminrule_idcompany,
			"companyadminrule_nmrule":    client.Companyadminrule_nmrule,
			"companyadminrule_rule":      client.Companyadminrule_rule,
		}).
		Post(PATH + "api/companyadminrulesave")
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
