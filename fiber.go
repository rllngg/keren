package keren

import (
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var pages = make(map[string]*Root)

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

			// Update the root object with the parsed values
			obj := map[string]string{}
			for k, v := range values {
				if len(v) > 0 {
					obj[k] = v[0]
					root.UpdateValue(k, v[0])
				}
			}

			// Set the response content type to HTML
			c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

			// Trigger the event on the root object and retrieve the event output
			eventOutput := root.TriggerEvent(elementID, event, c.Context())
			if len(root.PendingEvent) > 0 {
				c.Set("HX-Trigger", strings.Join(root.PendingEvent, ","))
				root.PendingEvent = []string{}
			}
			c.Set("HX-Retarget", "body")
			c.Set("HX-Reswap", "outerHTML")
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
