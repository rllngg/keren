package keren

import (
	"net/url"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/google/uuid"
)

type FiberKerenAdapter struct {
	storeSession *session.Store
	pages        map[string]*Root
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
func UseKerenWithFiber() {
	// handle css, js
	/// handle func template
}
func Response(c *fiber.Ctx, elem *Element) error {
	if elem.Root.RedirectURL != "" {
		c.Set("HX-Redirect", elem.Root.RedirectURL)
		elem.Root.RedirectURL = ""
	}
	c.Set("HX-Retarget", "#"+elem.ID)
	c.Set("HX-Reswap", "outerHTML")
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	c.SendString(HTMLTag(NewNode(elem), false))
	return nil
}
func NewFiberKerenAdapter(store *session.Store) *FiberKerenAdapter {
	return &FiberKerenAdapter{
		storeSession: store,
		pages:        map[string]*Root{},
	}
}
func (ctx *FiberKerenAdapter) Handle(handler func(*Root, *fiber.Ctx) error) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sess, err := ctx.storeSession.Get(c)
		defer sess.Save()
		if err != nil {
			return err
		}
		if c.Method() == "GET" {
			idInterface := sess.Get("page_id")
			id, ok := idInterface.(string)
			if !ok || id == "" {
				id = uuid.New().String()
				sess.Set("page_id", id)
			}
			ctx.pages[id] = NewRoot(DetectDevice(string(c.Request().Header.UserAgent())))
			ctx.pages[id].CurrentURL = c.OriginalURL()
			handler(ctx.pages[id], c)
			output := BuildHTML(ctx.pages[id])

			return c.Render("index", fiber.Map{
				"Content": output,
				"PageID":  id,
			})
		} else {
			// Retrieve the page ID from the request headers
			idInterface := sess.Get("page_id")
			pageID, ok := idInterface.(string)
			if !ok || pageID == "" {
				c.Set("HX-Refresh", "true")
				return c.SendString("Refresh")
			}

			// Retrieve the root object associated with the page ID
			root := ctx.pages[pageID]
			if root == nil {
				c.Set("HX-Refresh", "true")
				return c.SendString("Refresh")
			}

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

			// Update the root object with the parsed values
			obj := map[string]Data{}
			totalError := 0
			for k, v := range values {
				if len(v) > 0 {
					elem := root.UpdateValue(k, v[0])
					if elem == nil {
						continue
					}
					if elem.Validation != "" {
						errs := validate.Var(v[0], elem.Validation)
						// reset
						elem.RemoveClass("is-valid").RemoveClass("is-invalid").Parent.RemoveChildrenWithTag("div").RemoveClass("has-validation")
						if errs != nil {
							message := errs.Error()
							if elem.ErrorMessage != "" {
								message = elem.ErrorMessage
							}
							elem.AddClass("is-invalid").Parent.RemoveChildrenWithTag("div").Append(root.InvalidFeedback(message)).AddClass("has-validation")
							totalError++
						} else {
							elem.AddClass("is-valid").Parent.RemoveChildrenWithTag("div")
						}
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

				return Response(c, root.GetElementById(elementID))
			}

			// Set the response content type to HTML

			// Trigger the event on the root object and retrieve the event output
			eventOutput := root.TriggerEvent(elementID, event, c.Context(), obj)
			if len(root.PendingEvent) > 0 {
				c.Set("HX-Trigger", strings.Join(root.PendingEvent, ","))
				root.PendingEvent = []string{}
			}
			if eventOutput != nil {
				return Response(c, eventOutput)
			}
			c.Set("HX-Retarget", "body")
			c.Set("HX-Reswap", "outerHTML")
			c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

			// Build the HTML output from the root object
			output := BuildHTML(root)

			// Return the HTML output as the response
			return c.SendString(output)
		}
	}
}
