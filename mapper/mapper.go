package mapper

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

func Json2Map(s string) (m map[string]interface{}, err error) {
	m = make(map[string]interface{})
	err = json.Unmarshal([]byte(s), &m)
	return m, err
}

func Map2Json(m interface{}) (string, error) {
	v := reflect.ValueOf(m)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Map {
		return "", errors.New("the param is not a map")
	}
	marshal, err := json.Marshal(v.Interface())
	if err != nil {
		return "", err
	}
	return string(marshal), err
}

func Map2Struct(m interface{}, v interface{}) error {
	s, err := Map2Json(m)
	if err != nil {
		return err
	}
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return errors.New(fmt.Sprintf("the v(%+v) is not a pointer", v))
	}
	err = json.Unmarshal([]byte(s), v)
	return err
}

func Struct2Map(v interface{}) (map[string]interface{}, error) {
	value := reflect.ValueOf(v)
	for value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return nil, errors.New("the param is not a struct")
	}
	t := reflect.TypeOf(value.Interface())
	var m = make(map[string]interface{})
	for i := 0; i < value.NumField(); i++ {
		m[t.Field(i).Name] = value.Field(i).Interface()
	}
	return m, nil
}