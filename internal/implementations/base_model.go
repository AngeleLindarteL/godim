package implementations

import (
	"fmt"
	"reflect"
)

type BaseModel struct {
	modelMap       map[string]string
	childReference *BaseModel
}

type ChildModel struct {
	BaseModel `godim:"embedded"`

	Name   string  `godim:"name"`
	Age    int     `godim:"age"`
	Height float64 `godim:"height"`
}

func (b ChildModel) MapModel() string {
	return b.CreateModel(ChildModel{})
}

func (b *BaseModel) CreateModel(model any) string {
	fmt.Println("Creating model...")

	val := reflect.ValueOf(model)
	typ := reflect.TypeOf(model)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	modelMap := make(map[string]string)
	mapStructFields(val, typ, modelMap, true)

	return "Created"
}

func mapStructFields(val reflect.Value, typ reflect.Type, modelMap map[string]string, itsSubModel bool) {
	var subModelMap map[string]string

	if itsSubModel {
		subModelMap = make(map[string]string)
	}

	for i := range typ.NumField() {
		field := typ.Field(i)
		fieldValue := val.Field(i)

		goFieldName := field.Name
		goFieldType := field.Type
		goFieldValue := fieldValue.Interface()
		godimTag := field.Tag.Get("godim")

		if godimTag == "" || godimTag == "-" {
			fmt.Printf("Igonoring field: %s, Type: %v, Value: %v, Tag: %s\n",
				goFieldName, goFieldType, goFieldValue, godimTag)

			continue
		}

		if field.Anonymous && fieldValue.Kind() == reflect.Struct {
			if itsSubModel {
				mapStructFields(fieldValue, fieldValue.Type(), subModelMap, false)
			}
		}

		fmt.Printf("Field: %s, Type: %v, Value: %v, Tag: %s\n",
			goFieldName, goFieldType, goFieldValue, godimTag)

	}
}
