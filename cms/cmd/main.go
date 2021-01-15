package main 
import (
	"go_assessment\cms"
	os "os"
)

func main() {
	p := &cms.Page{
		Title: "Hello, world!",
		Content: "This is the body of our webpage",
	}

	cms.Tmpl.ExecuteTemplate(os.Stdout, "index", p)
}