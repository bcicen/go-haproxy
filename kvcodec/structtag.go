package kvcodec

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
	"regexp"
	"strings"
	"sync"
)

const (
	tagLabel = "mapstruct"
)

type structFields map[string]fieldMeta

type fieldMeta struct {
	Key        string
	Name       string
	OmitAlways bool
	OmitEmpty  bool
}

func readField(field reflect.StructField) (meta fieldMeta) {
	meta = fieldMeta{
		Name: field.Name,
	}
	fieldTags := strings.Split(field.Tag.Get(tagLabel), ",")
	for _, tag := range fieldTags {
		if tag == "-" {
			meta.OmitAlways = true
			return meta
		}

		if tag == "omitempty" {
			meta.OmitEmpty = true
		} else if tag != "" {
			meta.Key = tag
		} else {
			meta.Key = field.Name
		}
	}

	return meta
}

var structMap = make(map[reflect.Type]structFields)
var structMapMutex sync.RWMutex

func getstructFields(rType reflect.Type) (stInfo structFields) {
	structMapMutex.RLock()
	stInfo, ok := structMap[rType]
	if !ok {
		stInfo = getFieldInfos(rType)
		structMap[rType] = stInfo
	}
	structMapMutex.RUnlock()
	return stInfo
}

func getFieldInfos(rType reflect.Type) structFields {
	fieldsCount := rType.NumField()
	fieldMap := make(structFields)

	for i := 0; i < fieldsCount; i++ {
		field := rType.Field(i)
		meta := readField(field)

		if field.PkgPath != "" {
			continue
		}

		// if the field is an embedded struct, create a fieldInfo for each of its fields
		if field.Anonymous && field.Type.Kind() == reflect.Struct {
			panic(fmt.Errorf("embedded structs not supported"))
		}

		if !meta.OmitAlways {
			fieldMap[meta.Key] = meta
		}
	}

	return fieldMap
}

func Unmarshal(in io.Reader, out interface{}) error {
	outValue := reflect.ValueOf(out)
	if outValue.Kind() == reflect.Ptr {
		outValue = outValue.Elem()
	}
	outType := outValue.Type()

	fields := getstructFields(outType)
	//	for _, i := range fields {
	//		fmt.Printf("%s: %s\n", i.Name, i.Key)
	//	}

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), ":") {
			s := strings.Split(scanner.Text(), ":")
			k, v := s[0], s[1]
			if meta, ok := fields[k]; ok {
				field := outValue.FieldByName(meta.Name)
				setField(field, v)
			}
		}
	}
	return nil
}
