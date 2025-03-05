package implementations

import (
	"github.com/AngeleLindarteL/godim/pkg/enums"
)

type Filter struct {
	Field     string
	Value     *any
	Operation *enums.SupportedFiltersEnum
}

func (f *Filter) NewFilter(field string) *Filter {
	return &Filter{
		Field:     field,
		Operation: nil,
	}
}

func (f *Filter) Equals(value any) *Filter {
	operation := enums.SupportedFiltersEnumEq

	f.Operation = &operation
	f.Value = &value

	return f
}
