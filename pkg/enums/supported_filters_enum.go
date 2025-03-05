package enums

type SupportedFiltersEnum string

const (
	SupportedFiltersEnumEq    SupportedFiltersEnum = "eq"
	SupportedFiltersEnumNotEq SupportedFiltersEnum = "notEq"
	SupportedFiltersEnumIn    SupportedFiltersEnum = "in"
	SupportedFiltersEnumNotIn SupportedFiltersEnum = "notIn"
	SupportedFiltersEnumGt    SupportedFiltersEnum = "gt"
	SupportedFiltersEnumGte   SupportedFiltersEnum = "gte"
	SupportedFiltersEnumLt    SupportedFiltersEnum = "lt"
	SupportedFiltersEnumLte   SupportedFiltersEnum = "lte"
)
