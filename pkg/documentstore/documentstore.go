package documentstore

import "fmt"

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
	Value interface{}
}

type Document struct {
	Fields map[string]DocumentField
}

func (d Document) Put(doc Document) {
	for _, v := range doc.Fields {
		if v.Type == DocumentFieldTypeString {
			fmt.Println("string", v.Value)
		}
	}
}

func (d Document) Get(key string) (bool, *Document) {
	doc := Document{
		Fields: make(map[string]DocumentField),
	}
	for _, i := range d.List() {
		if dRes, ok := i.Fields[key]; ok {
			doc.Fields[key] = dRes
			return true, &doc
		}
	}
	return false, nil
}

func (d Document) Delete(key string) bool {
	for _, i := range d.List() {
		fmt.Println(i.Fields)
		if _, ok := i.Fields[key]; ok {
			delete(i.Fields, key)
			fmt.Println(i.Fields)
			return true
		}
	}
	return false
}

func (d Document) List() []Document {
	// Повертаємо список усіх документів
	listFields := Document{
		Fields: make(map[string]DocumentField),
	}
	listFields.Fields["Name_1"] = DocumentField{
		Type:  DocumentFieldTypeString,
		Value: "setup.exe",
	}
	listFields.Fields["Name_2"] = DocumentField{
		Type:  DocumentFieldTypeNumber,
		Value: 213143,
	}
	listFields.Fields["Name_3"] = DocumentField{
		Type:  DocumentFieldTypeBool,
		Value: true,
	}
	listFields.Fields["Name_4"] = DocumentField{
		Type:  DocumentFieldTypeArray,
		Value: []string{},
	}
	listFields.Fields["Name_5"] = DocumentField{
		Type:  DocumentFieldTypeObject,
		Value: DocumentField{},
	}
	aList := []Document{}
	aList = append(aList, listFields)

	return aList
}
