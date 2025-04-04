package documentstore

import (
	"fmt"
)

type DocumentFieldType string

const (
	DocumentFieldTypeString DocumentFieldType = "string"
	DocumentFieldTypeNumber DocumentFieldType = "number"
	DocumentFieldTypeBool   DocumentFieldType = "bool"
	DocumentFieldTypeArray  DocumentFieldType = "array"
	DocumentFieldTypeObject DocumentFieldType = "object"
)

type DocumentField struct {
	Type  DocumentFieldType
	Value any
}

type Document struct {
	Fields map[string]DocumentField
}

var store = map[string]Document{}

//// при добавленні перевіряю чи існує через Get в d.List() 
func (d Document) Put(doc Document) {
	data := d.List() 
	for k, v := range doc.Fields {
		b, _ := d.Get(k)
		if b {
			fmt.Println("document exist")
		} else {
			if v.Type == DocumentFieldTypeString && len(v.Value.(string)) > 0 {
				{
					store[k] = Document{
						Fields: make(map[string]DocumentField)}
					store[k].Fields[k] = doc.Fields[k]
					data = append(data, store[k])
				}
			}
		}
	}
}

func (d Document) Get(key string) (bool, *Document) {
	for _, v := range d.List() {
		if _, ok := v.Fields[key]; ok {
			return true, &v
		}
	}
	return false, nil
}

///при видаленні виводжу дані до і після
///для тесту роботи видалення беру дані із d.List()
func (d Document) Delete(key string) bool {
	data := d.List()
	fmt.Println("befor --- \n", d.List())
	for _, v := range data {
		if _, ok := v.Fields[key]; ok {
			delete(v.Fields, key)
			fmt.Println("\n after", data)
			return true
		}
	}
	return false
}


func (d Document) List() []Document {

	store = make(map[string]Document)

	f := store["Name1"]
	f.Fields = map[string]DocumentField{}
	f.Fields["Name1"] = DocumentField{
		Type:  DocumentFieldTypeString,
		Value: "event.go",
	}
	store["Name1"] = f

	f2 := store["Name2"]
	f2.Fields = map[string]DocumentField{}
	f2.Fields["Name2"] = DocumentField{
		Type:  DocumentFieldTypeString,
		Value: "setup.go",
	}
	store["Name2"] = f2

	f3 := store["Name3"]
	f3.Fields = map[string]DocumentField{}
	f3.Fields["Name3"] = DocumentField{
		Type:  DocumentFieldTypeString,
		Value: "event.go",
	}
	store["Name3"] = f3

	f4 := store["Name4"]
	f4.Fields = map[string]DocumentField{}
	f4.Fields["Name4"] = DocumentField{
		Type:  DocumentFieldTypeString,
		Value: "golang.go",
	}
	store["Name4"] = f4

	f5 := store["Name5"]
	f5.Fields = map[string]DocumentField{}
	f5.Fields["Name5"] = DocumentField{
		Type:  DocumentFieldTypeString,
		Value: "doc.go",
	}
	store["Name5"] = f5

	storeSlice := []Document{}
	for _, v := range store {
		storeSlice = append(storeSlice, v)
	}
	return storeSlice
}

