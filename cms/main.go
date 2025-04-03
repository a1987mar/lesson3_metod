package main

import (
	"les3/pkg/documentstore"
)

func main() {
	t := documentstore.Document{}
	//fmt.Println(t.Get("Name11"))

	allDoc := documentstore.Document{make(map[string]documentstore.DocumentField)}
	allDoc.Fields["name1"] = documentstore.DocumentField{
		Type:  documentstore.DocumentFieldTypeString,
		Value: "setup.exe",
	}
	allDoc.Fields["name4"] = documentstore.DocumentField{
		Type:  documentstore.DocumentFieldTypeString,
		Value: "event.go",
	}
	allDoc.Fields["name3"] = documentstore.DocumentField{
		Type:  documentstore.DocumentFieldTypeString,
		Value: "go.txt",
	}
	allDoc.Fields["name2"] = documentstore.DocumentField{
		Type:  documentstore.DocumentFieldTypeBool,
		Value: true,
	}

	t.Put(allDoc)

	///коментар

	//fmt.Println(t.Delete("Name11"))

	//fmt.Println(t.List())

	
}
