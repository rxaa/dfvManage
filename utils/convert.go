package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"shenxun-server/app/utils/quicktag"
	"strconv"
)

func StructToStringMap(obj interface{}) map[string]string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		switch t.Field(i).Type.Name() {
		case "string":
			data[t.Field(i).Name] = v.Field(i).String()
		case "int":
			data[t.Field(i).Name] = strconv.FormatInt(v.Field(i).Int(), 10)
		case "int8":
			data[t.Field(i).Name] = strconv.FormatInt(v.Field(i).Int(), 10)
		case "int16":
			data[t.Field(i).Name] = strconv.FormatInt(v.Field(i).Int(), 10)
		case "int32":
			data[t.Field(i).Name] = strconv.FormatInt(v.Field(i).Int(), 10)
		case "int64":
			data[t.Field(i).Name] = strconv.FormatInt(v.Field(i).Int(), 10)
		case "uint":
			data[t.Field(i).Name] = strconv.FormatUint(v.Field(i).Uint(), 10)
		case "uint8":
			data[t.Field(i).Name] = strconv.FormatUint(v.Field(i).Uint(), 10)
		case "uint16":
			data[t.Field(i).Name] = strconv.FormatUint(v.Field(i).Uint(), 10)
		case "uint32":
			data[t.Field(i).Name] = strconv.FormatUint(v.Field(i).Uint(), 10)
		case "uint64":
			data[t.Field(i).Name] = strconv.FormatUint(v.Field(i).Uint(), 10)
		case "float32":
			data[t.Field(i).Name] = strconv.FormatFloat(v.Field(i).Float(), 'f', -1, 64)
		case "float64":
			data[t.Field(i).Name] = strconv.FormatFloat(v.Field(i).Float(), 'f', -1, 64)
		case "bool":
			data[t.Field(i).Name] = strconv.FormatBool(v.Field(i).Bool())
		}
	}
	return data
}

func StringMapToStruct(data map[string]string, result interface{}) error {
	t := reflect.ValueOf(result).Elem()
	for k, v := range data {
		val := t.FieldByName(k)
		if !val.IsValid() {
			fmt.Printf("StringMapToStruct No such field: %s in datamap \n", k)
			continue
		}
		if !val.CanSet() {
			fmt.Printf("StringMapToStruct Cannot set %s field value \n", k)
			continue
		}
		dval := reflect.ValueOf(v)
		switch val.Type().Name() {
		case "string":
			val.SetString(dval.String())
		case "int":
			pv, _ := strconv.ParseInt(dval.String(), 10, 64)
			val.SetInt(pv)
		case "int8":
			pv, _ := strconv.ParseInt(dval.String(), 10, 64)
			val.SetInt(pv)
		case "int16":
			pv, _ := strconv.ParseInt(dval.String(), 10, 64)
			val.SetInt(pv)
		case "int32":
			pv, _ := strconv.ParseInt(dval.String(), 10, 64)
			val.SetInt(pv)
		case "int64":
			pv, _ := strconv.ParseInt(dval.String(), 10, 64)
			val.SetInt(pv)
		case "uint":
			pv, _ := strconv.ParseUint(dval.String(), 10, 64)
			val.SetUint(pv)
		case "uint8":
			pv, _ := strconv.ParseUint(dval.String(), 10, 64)
			val.SetUint(pv)
		case "uint16":
			pv, _ := strconv.ParseUint(dval.String(), 10, 64)
			val.SetUint(pv)
		case "uint32":
			pv, _ := strconv.ParseUint(dval.String(), 10, 64)
			val.SetUint(pv)
		case "uint64":
			pv, _ := strconv.ParseUint(dval.String(), 10, 64)
			val.SetUint(pv)
		case "float32":
			pv, _ := strconv.ParseFloat(dval.String(), 64)
			val.SetFloat(pv)
		case "float64":
			pv, _ := strconv.ParseFloat(dval.String(), 64)
			val.SetFloat(pv)
		case "bool":
			pv, _ := strconv.ParseBool(dval.String())
			val.SetBool(pv)
		}
	}
	return nil
}

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func StructToDbMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	fmt.Println(v)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		// 获取tag中的内容
		data[t.Field(i).Tag.Get("db")] = v.Field(i).Interface()
	}
	return data
}

//myData := make(map[string]interface{})
//myData["Name"] = "Tony"
//myData["Age"] = 23
//result := &MyStruct{}
//MapToStruct(myData, result)
//fmt.Println(result.Name)
func MapToStruct(data map[string]interface{}, result interface{}) {
	t := reflect.ValueOf(result).Elem()
	for k, v := range data {
		val := t.FieldByName(k)
		val.Set(reflect.ValueOf(v))
	}
}

/**
 * json编码, 并自动将字段大小转下划线小写
 */
func JsonEncode(data interface{}) string {
	res, err := json.Marshal(quicktag.Q(data))
	if err != nil {
		panic(err)
	}
	return string(res)
}

//json解码, 并自动匹配json下划线小写
func JsonDecode(str string, data interface{}) {
	err := json.Unmarshal([]byte(str), quicktag.Q(data))
	if err != nil {
		panic(err)
	}
}
