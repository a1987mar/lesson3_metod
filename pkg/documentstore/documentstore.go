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

var store = make(map[string]Document)

func (d Document) Put(doc Document) {
	for k, v := range doc.Fields {
		if _, ok := store[k]; !ok {
			if v.Type == DocumentFieldTypeString && len(v.Value.(string)) > 0 {
				store[k] = Document{
					Fields: make(map[string]DocumentField)}
				store[k].Fields[k] = doc.Fields[k]
			}

		}
	}

}

func (d Document) Get(key string) (bool, *Document) {
	data := d.List()
	doc := Document{
		Fields: make(map[string]DocumentField),
	}
	for _, i := range data {
		if dRes, ok := i.Fields[key]; ok {
			doc.Fields[key] = dRes
			return true, &doc
		}
	}
	return false, nil
}

func (d Document) Delete(key string) bool {
	l := d.List()              // беру список store по функції List()
	fmt.Println("befor \n", l) // виведення результат до DELETE
	for _, i := range l {
		if _, ok := i.Fields[key]; ok {
			delete(i.Fields, key)
			fmt.Println("after \n", l) // виведення результат після DELETE
			return true
		}
	}
	return false
}

func (d Document) List() []Document {
	f := store["Name"]
	f.Fields = map[string]DocumentField{}
	f.Fields["Name5"] = DocumentField{
		Type:  DocumentFieldTypeString,
		Value: "event.go",
	}
	f.Fields["Name8"] = DocumentField{
		Type:  DocumentFieldTypeNumber,
		Value: 123456789,
	}
	f.Fields["Name7"] = DocumentField{
		Type:  DocumentFieldTypeBool,
		Value: "event.go",
	}

	f2 := store["Name2"]
	f2.Fields = map[string]DocumentField{}
	f2.Fields["Name10"] = DocumentField{
		Type:  DocumentFieldTypeObject,
		Value: []DocumentField{},
	}
	f2.Fields["Name11"] = DocumentField{
		Type:  DocumentFieldTypeArray,
		Value: []string{},
	}
	f2.Fields["Name12"] = DocumentField{
		Type:  DocumentFieldTypeBool,
		Value: true,
	}

	return []Document{f, f2}
}
