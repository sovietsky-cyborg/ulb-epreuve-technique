package main

import (
	"ucl-epreuve-technique/app"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon app="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	a := app.App{}
	a.InitializeRoutes()
	a.Run()

}
