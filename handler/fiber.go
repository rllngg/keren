package handler

import (
	"net/url"

	"github.com/erlanggatampan/keren"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var pages = make(map[string]*keren.Root)

func FiberHandler(c *fiber.Ctx, handler func(*keren.Root, *fiber.Ctx) *keren.Root) error {
	if c.Method() == "GET" {

		id := uuid.New().String()
		pages[id] = keren.NewRoot()
		handler(pages[id], c)
		output := keren.BuildHTML(pages[id])

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
		eventOutput := root.TriggerEvent(elementID, event)
		c.Set("HX-Retarget", "body")
		c.Set("HX-Reswap", "outerHTML")
		if eventOutput != nil {
			c.Set("HX-Retarget", "#"+eventOutput.ID)
			return c.SendString(keren.HTMLTag(keren.NewNode(eventOutput)))
		}

		// Build the HTML output from the root object
		output := keren.BuildHTML(root)

		// Return the HTML output as the response
		return c.SendString(output)
	}
}
