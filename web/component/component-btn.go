package component

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type MyButton struct {
	app.Compo
	n        int
	NameInit string
}

func (btn *MyButton) OnMount(ctx app.Context) {
	fmt.Println("button mounted")
}

func (btn *MyButton) Render() app.UI {
	var txt string
	if btn.n == 0 {
		txt = btn.NameInit
	} else {
		txt = fmt.Sprintf("clicked %v", btn.n)
	}
	return app.Button().Class("button mr-3").Text(txt).OnClick(btn.onClick)
}

func (btn *MyButton) onClick(ctx app.Context, e app.Event) {
	fmt.Println("onClick is called")
	btn.n++
}
