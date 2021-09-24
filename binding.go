package binding

import (
	"github.com/billcoding/reflectx"
	"net/http"
	"reflect"
	"strings"
)

// Binding struct
type Binding struct {
	structPtr interface{}
	typ       *Type
	fields    []*reflect.StructField
	items     []interface{}
	values    []*reflect.Value
	dataMap   map[string]interface{}
}

// New return new Binding
func New(structPtr interface{}) *Binding {
	return NewWithType(structPtr, Param)
}

// NewWithType return new Binding
func NewWithType(structPtr interface{}, typ *Type) *Binding {
	structFields, structFieldValues, items := reflectx.ParseTag(structPtr, new(Item), "alias", "binding", true)
	return &Binding{
		structPtr: structPtr,
		typ:       typ,
		fields:    structFields,
		items:     items,
		values:    structFieldValues,
		dataMap:   make(map[string]interface{}, 0),
	}
}

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

func (b *Binding) setVal() {
	for i, field := range b.fields {
		setVal(b.typ, b.items[i].(*Item), field, b.values[i], b.dataMap)
	}
}

func setVal(typ *Type, item *Item, field *reflect.StructField, value *reflect.Value, dataMap map[string]interface{}) {
	var innerInterface interface{}
	switch {
	case field.Type.Kind() == reflect.Struct:
		innerInterface = value.Addr().Interface()
	case field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct:
		innerInterface = value.Elem().Addr().Interface()
	default:
		item.Bind(field, *value, dataMap)
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
	fields, fieldValues, items := reflectx.ParseTag(innerInterface, new(Item), "alias", "binding", false)
	b := &Binding{
		structPtr: innerInterface,
		typ:       typ,
		fields:    fields,
		items:     items,
		values:    fieldValues,
		dataMap:   dataMap,
	}
	b.BindMap(innerDataMap)
}

func recursiveMap(dataMap map[string]interface{}) {
	for k := range dataMap {
		firstPos := strings.IndexByte(k, '.')
		if firstPos != -1 {
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
				m[sk] = v
			} else {
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
	if req != nil {
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
			if len(cts) > 0 && strings.EqualFold(strings.TrimSpace(cts[0]), "multipart/form-data") {
				err := req.ParseMultipartForm(0)
				if err != nil {
				} else {
					for k, v := range req.MultipartForm.Value {
						setMap(b.dataMap, k, v)
					}
				}
			}
		}
		recursiveMap(b.dataMap)
	}
}
