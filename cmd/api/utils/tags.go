package utils

import (
	"errors"
	"fmt"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/logs"
	"gorm.io/gorm/utils"
	"reflect"
	"sort"
	"strings"
)

type Tags map[string]string

type dataSlice []string

// Len is part of sort.Interface.
func (d dataSlice) Len() int {
	return len(d)
}

// Swap is part of sort.Interface.
func (d dataSlice) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Less is part of sort.Interface. We use count as the value to sort by.
func (d dataSlice) Less(i, j int) bool {
	return d[i] < d[j]
}

func BuildTags(tags Tags) []string {
	taglist := []string{}

	for k, v := range tags {
		taglist = append(taglist, fmt.Sprintf("%s:%s", k, v))
	}

	sortList := make(dataSlice, 0, len(taglist))
	sortList = append(sortList, taglist...)
	sort.Sort(sortList)

	return sortList
}

func Merge(tags1 []string, tags2 ...string) []string {
	return append(tags1, tags2...)
}

// AppendTags ...
func AppendTags(builtTags *[]string, newTags Tags) {
	newTaglist := BuildTags(newTags)

	*builtTags = append(*builtTags, newTaglist...)
}

func TagsToFields(tags Tags) []logs.Field {
	fields := []logs.Field{}

	for k, v := range tags {
		fields = append(fields, logs.Any(k, v))
	}

	return fields
}

func MergeFields(fields1 []logs.Field, fields2 ...logs.Field) []logs.Field {
	return append(fields1, fields2...)
}

func BuildedTagsToFields(tags []string) []logs.Field {
	fields := []logs.Field{}

	for _, tag := range tags {
		k, v := SplitTag(tag)
		fields = append(fields, logs.Any(k, v))
	}

	return fields
}

func SplitTag(tag string) (string, string) {
	separatorIndex := strings.Index(tag, ":")
	return tag[0:separatorIndex], tag[separatorIndex+1:]
}

// GetParams will extract the values from the fields of the struct v to be used as parameters.
// The field should be considered as a parameter if it has the tag "param" or is exported in which case
// the field name will be used as the parameter name.
// The field will be ignored if it has the tag "param" with the value "-".
// The field values will be converted to string using the function toString.
func GetParams(value any) map[string]string {
	if value == nil {
		panic("value is nil")
	}

	rv, err := reflectValue(value)
	if err != nil {
		panic(fmt.Errorf("failed to obtain reflect value: %v", err))
	}

	params := make(map[string]string)
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Type().Field(i)
		if !field.IsExported() {
			continue
		}

		tag := field.Tag.Get("param")
		if tag == "-" {
			continue
		}

		if tag == "" {
			tag = field.Name
		}

		v := rv.Field(i).Interface()
		params[tag] = utils.ToString(v)
	}

	return params
}

// reflectValue will obtain the [reflect.Value] of v only if it is a struct or a pointer to a struct.
// If it is a pointer to a struct, it will dereference it and return the [reflect.Value] of the struct.
// Otherwise, it will return an error.
func reflectValue(v any) (reflect.Value, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return reflect.Value{}, errors.New("value is not a struct or a pointer to a struct")
	}

	return rv, nil
}
