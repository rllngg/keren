package keren

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var pages = make(map[string]*Root)
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
func UseKerenWithFiber() {
	// handle css, js
	/// handle func template
}
func FiberHandler(handler func(*Root, *fiber.Ctx) error) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if c.Method() == "GET" {

			id := uuid.New().String()
			pages[id] = NewRoot(DetectDevice(string(c.Request().Header.UserAgent())))
			handler(pages[id], c)
			output := BuildHTML(pages[id])

			return c.Render("index", fiber.Map{
				"Content": output,
				"PageID":  id,
			})
		} else {
			// Retrieve the page ID from the request headers
			pageID := c.Get("Hx-Page-Id")

			if pageID == "" {
				c.Set("HX-Refresh", "true")
				return c.SendString("Refresh")
			}

			// Retrieve the root object associated with the page ID
			root := pages[pageID]
			if root == nil {
				c.Set("HX-Refresh", "true")
				return c.SendString("Refresh")
			}

			// Retrieve the element ID and event type from the request headers
			elementID := string(c.Request().Header.Peek("Hx-Trigger"))
			event := string(c.Request().Header.Peek("Hx-Event"))
			if event == "" {
				event = "load"
			}

			// Parse the values from the request body
			values, err := url.ParseQuery(string(c.Body()))
			if err != nil {
				return err
			}
			c.Set("HX-Retarget", "body")
			c.Set("HX-Reswap", "outerHTML")
			c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

			// Update the root object with the parsed values
			obj := map[string]Data{}
			for k, v := range values {
				if len(v) > 0 {
					elem := root.UpdateValue(k, v[0])
					if elem.Validation != "" {
						errs := validate.Var(v[0], elem.Validation)
						if errs != nil {
							fmt.Println("Validation Error : ", errs)
							c.Set("HX-Retarget", "#"+elem.ID)
							return c.SendString(HTMLTag(NewNode(elem.RemoveChildren().Append(root.AlertMessage(errs.Error(), "danger")))))
						}
					}
					obj[elem.Name] = Data{
						Value: v[0],
						File:  nil,
					}
					fmt.Println("Element Name : ", elem.Name, "Value : ", v[0])
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

			// Set the response content type to HTML

			// Trigger the event on the root object and retrieve the event output
			eventOutput := root.TriggerEvent(elementID, event, c.Context(), obj)
			if len(root.PendingEvent) > 0 {
				c.Set("HX-Trigger", strings.Join(root.PendingEvent, ","))
				root.PendingEvent = []string{}
			}

			if eventOutput != nil {
				c.Set("HX-Retarget", "#"+eventOutput.ID)
				return c.SendString(HTMLTag(NewNode(eventOutput)))
			}

			// Build the HTML output from the root object
			output := BuildHTML(root)

			// Return the HTML output as the response
			return c.SendString(output)
		}
	}
}
