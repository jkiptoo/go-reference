// Using XML with encoding/xml package
package main

import (
	"encoding/xml"
	"fmt"
)

//Create struct using field tags to create directives for encoding and decoding xml
type Crop struct {
	XMLName xml.Name `xml:"crop"`
	Id int `xml:"id, attr"`
	Name string `xml:"name"`
	Origin []string `xml:"origin"`
}

func (c Crop) String() string {
		return fmt.Sprintf("Crop id=%v, name=%v, origin=%v", c.Id, c.Name, c.Origin)
	}


func main() {
	tea := &Crop{Id: 30, Name: "Tea"}
	tea.Origin = []string{"Kenya", "Sri Lanka"}

	//Use MarshalIndent to output more human-readable content. This helps remove XML representing Crop
	out, _ := xml.MarshalIndent(tea, " ", " ")
	fmt.Println(string(out))

	//Append XML header to add it to the output
	fmt.Println(xml.Header + string(out))

	//Parse a stream of bytes into XML into a data structure by using Unmarshal
	var c Crop 
	if err := xml.Unmarshal(out, &c);
	err != nil {
		panic(err)
	} 
	fmt.Println(c)

	maize := &Crop{Id: 52, Name: "Maize"}
	maize.Origin = []string{"Tanzania", "Uganda"}

	//Nest crops under parent>child>crop construct
	type xmlNest struct {
		XMLName xml.Name `xml:"Nesting"`
		Crops []*Crop `xml:"parent>child>crop"`
	}

	nesting := &xmlNest{}
	nesting.Crops = []*Crop{tea, maize}

	out, _ = xml.MarshalIndent(nesting, " ", " ")
	fmt.Println(string(out))


}	
