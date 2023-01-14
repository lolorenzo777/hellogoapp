package component

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Content struct {
	app.Compo
	data []string
}

func (c *Content) OnMount(ctx app.Context) {
	fmt.Println("content: OnMount")

	c.data = make([]string, 0)
	c.data = append(c.data, "john")
	c.data = append(c.data, "greta")
	c.data = append(c.data, "bob")
	c.data = append(c.data, "alice")
	c.data = append(c.data, "mike")
}

func (c *Content) Render() app.UI {
	fmt.Println("content: Render")

	return app.Div().Body(
		app.H2().Text("Content"),
		app.P().Text("This is my text"),
		app.If(len(c.data) > 0,
			app.Ul().Body(
				app.Range(c.data).Slice(func(i int) app.UI {
					return app.Li().Text(c.data[i])
				}),
			),
		).Else(
			app.P().Text("No data"),
		),
	)
}

func (c *Content) OnDismount() {
	fmt.Println("content: OnDismount")
}
