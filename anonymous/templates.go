package anonymous

import (
	"os"
	"text/template"
)

func TemplatesDemo() {
	// creating a new template with name t1.
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// must will panic automatically in case of error from parse.
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// the execute function will replace the {{.}} with the text.
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// function to create a new template.
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// using the template with struct and map.
	t2 := Create("t2", "Name: {{.Name}}\n")
	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})
}
