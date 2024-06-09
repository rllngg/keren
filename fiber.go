package keren

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/url"
	"strings"
	"time"
)

type FiberKerenAdapter struct {
	Sessions       map[string]*App
	SessionTimeout time.Duration
}

var validate *validator.Validate

func init() {
	validate = validator.New()

}

func DetectDevice(agent string) string {
	if strings.Contains(agent, "Mobile") {
		return "mobile"
	}
	return "desktop"

}

func Response(c *fiber.Ctx, elem *Element) error {
	response := elem.App.BuildHTML(elem)
	if response.RedirectURL != "" {
		c.Set("HX-Redirect", response.RedirectURL)
	}

	c.Set("HX-Retarget", "body")
	c.Set("HX-Reswap", "outerHTML")
	if elem != nil {
		c.Set("HX-Retarget", "#"+elem.ID)
	}
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

	return c.SendString(response.HTML)
}
func NewFiberKerenAdapter(timeout int) *FiberKerenAdapter {
	adapter := &FiberKerenAdapter{
		Sessions:       map[string]*App{},
		SessionTimeout: time.Duration(timeout) * time.Second,
	}
	// Remove Sessions everytimeout
	fmt.Println("[KERENFIBER] Waiting for sessions collector...")
	go func() {
		for {
			fmt.Println("[KERENFIBER] Waiting for sessions collector...")
			time.Sleep(adapter.SessionTimeout * time.Second)
			adapter.CleanSessions()
		}
	}()
	return adapter
}
func (adapter *FiberKerenAdapter) CleanSessions() {
	totalDeleted := 0
	fmt.Println("[KERENFIBER] Sessions Size ", Of(adapter.Sessions)/10000, "mb")
	fmt.Println("[KERENFIBER] Clearing Sessions")
	for id, session := range adapter.Sessions {
		// check if session has timed out
		if time.Now().After(session.LastUpdate.Add(adapter.SessionTimeout)) {
			// if it has, remove it from your session manager
			delete(adapter.Sessions, id)
			totalDeleted++
		}
	}
	fmt.Println("[KERENFIBER] Cleared Sessions Size ", Of(adapter.Sessions)/10000, "mb")
	fmt.Println("[KERENFIBER] Total Cleared Sessions")
}
func (adapter *FiberKerenAdapter) Handle(handler func(*App, *fiber.Ctx) error) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if c.Method() == "GET" {
			initialize := false
			appID := c.Get("Hx-App-ID")
			var app *App
			if appID == "" {
				appID = uuid.New().String()
				initialize = true
			}
			app = NewApp(DetectDevice(string(c.Request().Header.UserAgent())))
			app.Lock()
			defer app.Unlock()
			adapter.Sessions[appID] = app
			app.CurrentURL = c.OriginalURL()
			handler(app, c)

			println(appID, " : ", Of(app)/1000, "kb")
			println("Session Holder : ", Of(adapter.Sessions)/1000, "kb")

			response := app.BuildHTML(nil)
			if !initialize {
				c.Set("HX-Retarget", "body")
				c.Set("HX-Reswap", "outerHTML")
			}
			return c.Render("index", fiber.Map{
				"Content": response.HTML,
				"AppID":   appID,
			})
		} else {
			// Retrieve the page ID from the request headers
			pageID := c.Get("Hx-App-ID")
			c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
			app := adapter.Sessions[pageID]

			if app == nil || pageID == "" {
				c.Set("HX-Refresh", "true")
				return c.SendString("Refresh")
			}
			app.Lock()
			defer app.Unlock()

			// Retrieve the element ID and event type from the request headers
			elementID := string(c.Request().Header.Peek("Hx-Trigger"))
			event := strings.Trim(string(c.Request().Header.Peek("Hx-Event")), " ")
			if event == "" {
				event = "default"
			}

			// Parse the values from the request body
			values, err := url.ParseQuery(string(c.Body()))
			if err != nil {
				return err
			}

			// Update the app object with the parsed values
			obj := map[string]Data{}
			totalError := 0
			for k, v := range values {
				if len(v) > 0 {
					elem, err := app.UpdateValue(k, v[0])
					if err != nil {
						totalError = totalError + 1
						continue
					}
					obj[elem.Name] = Data{
						Value: v[0],
						File:  nil,
					}
					if elem.Tag == "input" {
						if elem.HasAttribute("type") && elem.GetAttribute("type") == "file" {
							file, err := c.FormFile(k)
							if err == nil {
								obj[elem.Name] = Data{
									Value: v[0],
									File:  file,
								}
							}
						}

					}
				}
			}
			if totalError > 0 {
				// form

				return Response(c, app.GetElementById(elementID))
			}

			// Set the response content type to HTML

			// Trigger the event on the app object and retrieve the event output
			eventOutput := app.TriggerEvent(elementID, event, c.Context(), obj)
			if eventOutput != nil {
				return Response(c, eventOutput)
			}
			c.Set("HX-Retarget", "body")
			c.Set("HX-Reswap", "outerHTML")

			// Build the HTML output from the app object
			output := BuildHTML(app)

			// Return the HTML output as the response
			return c.SendString(output)
		}
	}
}
