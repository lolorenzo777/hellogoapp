package component

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type Page2 struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *Page2) Render() app.UI {
	return app.Section().Class("section").
		Body(app.Div().Class("container").
			Body(app.H1().Class("title").Text("Page2"),
				app.A().Class("button").Text("home").Href("/"),
			),
		)
}
