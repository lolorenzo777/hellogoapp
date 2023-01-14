package component

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type Hello struct {
	app.Compo
	section2 app.HTMLSection
}

func (h *Hello) OnPreRender(ctx app.Context) {
	fmt.Println("hello: OnPreRender")
}

func (h *Hello) OnMount(ctx app.Context) {
	fmt.Println("hello: OnMount")
	//h.section2 = app.Section().Class("section").ID("section2")
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *Hello) Render() app.UI {
	fmt.Println("hello: Render")

	return app.Div().Body(
		app.Section().Class("section").ID("section1").
			Body(app.Div().Class("container").
				Body(app.H1().Class("title").Text("Hello World !"),
					app.Raw(`<p class="subtitle">My first website with <strong>Bulma</strong>!</p>`),
					app.Div().Class("hello").
						Body(&MyButton{NameInit: "click here"},
							app.Button().Class("button mr-3 mb-3").Text("Navigate to page1").OnClick(h.onClickPage1),
							app.A().Class("button mr-3 is-primary").Text("URL page2").Href("/page2"),
							app.Button().Class("button mr-3 mb-3").Text("Force server reload").OnClick(h.onClickReload),
							app.Button().Class("button mr-3 mb-3").Text("Go To home").OnClick(h.onClickGotohome),
							app.Button().Class("button mr-3 mb-3").Type("button").Text("Toggle content").OnClick(h.onClickContent),
						),
				),
			),
		//h.section2,
	)
}

func (h *Hello) OnDismount() {
	fmt.Println("hello: OnDismount")
}

func (h *Hello) onClickPage1(ctx app.Context, e app.Event) {
	fmt.Println("hello: onClick page 1 is called")
	ctx.Navigate("/page1")
}

func (h *Hello) onClickReload(ctx app.Context, e app.Event) {
	fmt.Println("hello: onClick reload is called")
	ctx.Reload()
}

func (h *Hello) onClickContent(ctx app.Context, e app.Event) {
	fmt.Println("hello: onClick content is called")

	//h.section2.Body(&Content{})
	//h.Update()
}

func (h *Hello) onClickGotohome(ctx app.Context, e app.Event) {
	fmt.Println("hello: onClick content is called")

	ctx.Defer(func(ctx app.Context) {
		app.Window().Get("location").Set("href", "/")
	})
}
