// We will use text/template package to generate dynamic/custom content output for the user
// html/template package is used for HTML APIs
package main

import (
	"os"
	"text/template"
)

func main() {
	//Create a new template(a mix of static text and actions) and parse its body from a string
	firstTemplate := template.New("First Template")
	firstTemplate, err := firstTemplate.Parse("The value is {{.}}\n")
	if err !=nil {
		panic(err)
	}

	//For global scoped templates, use template.Must to return panic in case a parse error is returned
	firstTemplate = template.Must(firstTemplate.Parse("The value is {{.}}\n"))

	//Replace {{.}} action with the value passed a parameter to Execute
	//Text is generated with specific actions for the template
	firstTemplate.Execute(os.Stdout, "Here is the text")
	firstTemplate.Execute(os.Stdout, 5)
	firstTemplate.Execute(os.Stdout, []string {"Go", "Rust", "C++", "C#",})

	//Create a helper function
	Create := func(name, secondTemplate string) *template.Template {
		return template.Must(template.New(name).Parse(secondTemplate))
	}

	//Export fields to be accessible when the template is executing.
	//Use {{.FieldName}} action to access its fields

	secondTemplate := Create("Second Template", "Name: {{.Name}}\n")
	secondTemplate.Execute(os.Stdout, struct {
		Name string
	}{"Mister Doe"})

	//Use maps to execute the above
	secondTemplate.Execute(os.Stdout, map[string]string{
		"Name": "Michael Brown",
	})

	//Use if/else to provide conditional execution for templates. 
	//Default type values are considered false e.g. empty string, 0, nil pointer 
	//Use - flag to trim whitespace
	thirdTemplate := Create("Third Template", "{{if . -}} yes {{else -}} no {{end}}\n")
	thirdTemplate.Execute(os.Stdout, "String is not empty")
	thirdTemplate.Execute(os.Stdout, "")

	//Use range to loop through slices, arrays, maps or channels
	//Set {{.}} to the current item of being iterated
	fourthTemplate := Create("Fourth Template", "Range: {{range .}}{{.}} {{end}}\n")
	fourthTemplate.Execute(os.Stdout,[]string {
		"Go", "Rust","C++","C#",
	})

}