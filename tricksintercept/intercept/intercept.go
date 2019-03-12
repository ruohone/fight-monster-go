package intercept

import (
	"errors"
	"fmt"
	"reflect"
)

type Intercept struct {
	
}

var FF = func() {}

var cacheMap = make(map[string]interface{})
var DataEmpty = errors.New("数据为空")

func (i Intercept) Cache(key string,fi interface{},params ...interface{}) (interface{},error) {
	f := reflect.ValueOf(fi)
	result,ok := cacheMap[key]
	if ok {
		fmt.Println("from cache get")
		return result,nil
	}

	in := make([]reflect.Value, len(params))
	for k,p := range params {
		in[k] = reflect.ValueOf(p)
	}

	vs := f.Call(in)

	if vs[1].Interface() != nil {
		fmt.Println("from db get")
		return nil,vs[1].Interface().(error)
	}

	result = vs[0].Interface()
	if result==nil {
		return nil,DataEmpty
	}
	fmt.Println("save to db")
	cacheMap[key] = result
	return result,nil
}