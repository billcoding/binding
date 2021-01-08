package binding

import (
	"encoding/json"
	"encoding/xml"
	"github.com/billcoding/calls"
	"github.com/billcoding/reflectx"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
)

// Binding struct
type Binding struct {
	structPtr interface{}
	typ       *Type
	items     []*Item
	fields    []*reflect.StructField
	logger    *log.Logger
	dataMap   map[string]interface{}
}

// New return new Binding
func New(structPtr interface{}) *Binding {
	return NewWithType(structPtr, Body)
}

// NewWithType return new Binding
func NewWithType(structPtr interface{}, typ *Type) *Binding {
	items := make([]*Item, 0)
	fields := reflectx.CreateFromTag(structPtr, &items, "alias", "binding")
	if len(items) != len(fields) {
		panic("[New]invalid len both items and fields")
	}
	return &Binding{
		structPtr: structPtr,
		typ:       typ,
		items:     items,
		fields:    fields,
		logger:    log.New(os.Stdout, "[Binding]", log.LstdFlags),
		dataMap:   make(map[string]interface{}, 0),
	}
}

//setMap
func setMap(m map[string]interface{}, k string, v []string) {
	if len(v) > 0 {
		if len(v) == 1 {
			m[k] = v[0]
		} else {
			m[k] = v
		}
	}
}

// BindReq start bind from Request
func (b *Binding) BindReq(req *http.Request) {
	b.initMapFromReq(req)
	b.setVal()
}

// BindMap start bind from data map
func (b *Binding) BindMap(dataMap map[string]interface{}) {
	b.dataMap = dataMap
	b.setVal()
}

// BindJSON start bind from JSON
func (b *Binding) BindJSON(jsonData string) {
	m := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(jsonData), &m)
	calls.NNil(err, func() {
		b.logger.Printf("[BindJSON]%v\n", err)
	})
	calls.Nil(err, func() {
		b.dataMap = m
		b.setVal()
	})
}

// BindXML start bind from XML
func (b *Binding) BindXML(xmlData string) {
	m := make(map[string]interface{}, 0)
	err := xml.Unmarshal([]byte(xmlData), &m)
	calls.NNil(err, func() {
		b.logger.Printf("[BindXML]%v\n", err)
	})
	calls.Nil(err, func() {
		b.dataMap = m
		b.setVal()
	})
}

func (b *Binding) setVal() {
	for pos, item := range b.items {
		field := b.fields[pos]
		value := reflect.ValueOf(b.structPtr).Elem().FieldByName(field.Name)
		setVal(b.typ, item, field, value, b.dataMap)
	}
}

func setVal(typ *Type, item *Item, field *reflect.StructField, value reflect.Value, dataMap map[string]interface{}) {
	var innerInterface interface{}
	switch {
	case field.Type.Kind() == reflect.Struct:
		innerInterface = value.Addr().Interface()
	case field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct:
		innerInterface = value.Elem().Addr().Interface()
	default:
		item.Bind(field, value, dataMap)
		return
	}
	if innerInterface == nil {
		return
	}
	_innerDataMap, have := dataMap[item.Name]
	if !have {
		return
	}
	innerDataMap, ok := _innerDataMap.(map[string]interface{})
	if !ok {
		return
	}
	items := make([]*Item, 0)
	fields := reflectx.CreateFromTag(innerInterface, &items, "alias", "binding")
	if len(items) != len(fields) {
		panic("[New]invalid len both items and fields")
	}
	b := &Binding{
		structPtr: innerInterface,
		typ:       typ,
		items:     items,
		fields:    fields,
		logger:    log.New(os.Stdout, "[Binding]", log.LstdFlags),
	}
	b.BindMap(innerDataMap)
}

func recursiveMap(dataMap map[string]interface{}) {
	for k := range dataMap {
		// model2.model3.model4.name
		firstPos := strings.IndexByte(k, '.')
		if firstPos != -1 {
			// model2 => model3 => model4
			prefix := k[:firstPos]
			if dataMap[prefix] == nil {
				subMap := getSubMap(prefix, dataMap)
				dataMap[prefix] = subMap
			}
		}
	}
}

func getSubMap(prefix string, dataMap map[string]interface{}) map[string]interface{} {
	m := make(map[string]interface{}, 0)
	for k, v := range dataMap {
		if strings.HasPrefix(k, prefix+".") {
			sk := strings.TrimPrefix(k, prefix+".")
			endPos := strings.IndexByte(sk, '.')
			if endPos == -1 {
				// the end
				m[sk] = v
			} else {
				// next prefix
				midPrefix := sk[:endPos]
				nextPrefix := strings.Join([]string{prefix, midPrefix}, ".")
				if m[midPrefix] == nil {
					subMap := getSubMap(nextPrefix, dataMap)
					m[midPrefix] = subMap
				}
			}
		}
	}
	return m
}

func (b *Binding) initMapFromReq(req *http.Request) {
	if req == nil {
		return
	}
	switch b.typ {
	case Header:
		for k, v := range req.Header {
			setMap(b.dataMap, k, v)
		}
	case Param:
		_ = req.ParseForm()
		for k, v := range req.Form {
			setMap(b.dataMap, k, v)
		}
		cts := strings.Split(req.Header.Get("Content-Type"), ";")
		if len(cts) > 0 {
			switch strings.TrimSpace(cts[0]) {
			case "multipart/form-data":
				err := req.ParseMultipartForm(0)
				if err == nil {
					for k, v := range req.MultipartForm.Value {
						setMap(b.dataMap, k, v)
					}
				}
			}
		}
	case Body:
		cts := strings.Split(req.Header.Get("Content-Type"), ";")
		if len(cts) <= 0 {
			b.logger.Println("[Bind]Not found header 'Content-Type'")
			return
		}
		switch strings.TrimSpace(cts[0]) {
		case "application/json":
			bytes, err := ioutil.ReadAll(req.Body)
			if err == nil && json.Valid(bytes) {
				err := json.Unmarshal(bytes, &b.dataMap)
				if err != nil {
					b.logger.Printf("[Bind]%v\n", err)
				}
			}
		default:
			b.logger.Println("[Bind]Only support Content-Type of 'application/json'")
		}
	}

	if b.typ == Header || b.typ == Param {
		recursiveMap(b.dataMap)
	}
}
