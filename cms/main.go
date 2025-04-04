package main

import (
	"les3/pkg/documentstore"
)

func main() {
	t := documentstore.Document{}

	// fmt.Println(t.Get("Name6"))

	// allDoc := documentstore.Document{Fields: make(map[string]documentstore.DocumentField)}
	// allDoc.Fields["Name1"] = documentstore.DocumentField{
	// 	Type:  documentstore.DocumentFieldTypeString,
	// 	Value: "setup.exe",
	// }
	// allDoc.Fields["Name4"] = documentstore.DocumentField{
	// 	Type:  documentstore.DocumentFieldTypeString,
	// 	Value: "event.go",
	// }
	// allDoc.Fields["Name6"] = documentstore.DocumentField{
	// 	Type:  documentstore.DocumentFieldTypeString,
	// 	Value: "go.txt",
	// }
	// allDoc.Fields["Name2"] = documentstore.DocumentField{
	// 	Type:  documentstore.DocumentFieldTypeBool,
	// 	Value: true,
	// }

	// t.Put(allDoc)

	//fmt.Println(t.Delete("Name4"))

	fmt.Println(t.List())
	
}
