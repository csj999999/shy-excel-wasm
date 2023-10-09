package main

import (
	"bytes"
	"errors"
	shyexcel "github.com/buzzxu/shy-excel"
	"github.com/xuri/excelize/v2"
	"net/http"
	"reflect"
	"syscall/js"
)

type argsRule struct {
	opts  bool
	types []js.Type
}

var (
	// goBaseTypes defines Go's basic data types.
	goBaseTypes = map[reflect.Kind]bool{
		reflect.Bool:    true,
		reflect.Int:     true,
		reflect.Int8:    true,
		reflect.Int16:   true,
		reflect.Int32:   true,
		reflect.Int64:   true,
		reflect.Uint:    true,
		reflect.Uint8:   true,
		reflect.Uint16:  true,
		reflect.Uint32:  true,
		reflect.Uint64:  true,
		reflect.Uintptr: true,
		reflect.Float32: true,
		reflect.Float64: true,
		reflect.Map:     true,
		reflect.String:  true,
	}
	// jsToBaseGoTypeFuncs defined functions mapping for JavaScript to Go basic
	// data types convention.
	jsToBaseGoTypeFuncs = map[reflect.Kind]func(jsVal js.Value, kind reflect.Kind) (reflect.Value, error){
		reflect.Bool: func(jsVal js.Value, kind reflect.Kind) (reflect.Value, error) {
			if jsVal.Type() != js.TypeBoolean {
				return reflect.ValueOf(nil), errArgType
			}
			return reflect.ValueOf(jsVal.Bool()), nil
		},
		reflect.Uint: func(jsVal js.Value, kind reflect.Kind) (reflect.Value, error) {
			if jsVal.Type() != js.TypeNumber {
				return reflect.ValueOf(nil), errArgType
			}
			return reflect.ValueOf(uint(jsVal.Float())), nil
		},
		reflect.Uint8: func(jsVal js.Value, kind reflect.Kind) (reflect.Value, error) {
			if jsVal.Type() != js.TypeNumber {
				return reflect.ValueOf(nil), errArgType
			}
			return reflect.ValueOf(uint8(jsVal.Float())), nil
		},
		reflect.Uint64: func(jsVal js.Value, kind reflect.Kind) (reflect.Value, error) {
			if jsVal.Type() != js.TypeNumber {
				return reflect.ValueOf(nil), errArgType
			}
			return reflect.ValueOf(uint64(jsVal.Float())), nil
		},
		reflect.Int: func(jsVal js.Value, kind reflect.Kind) (reflect.Value, error) {
			if jsVal.Type() != js.TypeNumber {
				return reflect.ValueOf(nil), errArgType
			}
			return reflect.ValueOf(int(jsVal.Float())), nil
		},
		reflect.Int64: func(jsVal js.Value, kind reflect.Kind) (reflect.Value, error) {
			if jsVal.Type() != js.TypeNumber {
				return reflect.ValueOf(nil), errArgType
			}
			return reflect.ValueOf(int64(jsVal.Float())), nil
		},
		reflect.Float64: func(jsVal js.Value, kind reflect.Kind) (reflect.Value, error) {
			if jsVal.Type() != js.TypeNumber {
				return reflect.ValueOf(nil), errArgType
			}
			return reflect.ValueOf(jsVal.Float()), nil
		},
		reflect.String: func(jsVal js.Value, kind reflect.Kind) (reflect.Value, error) {
			if jsVal.Type() != js.TypeString {
				return reflect.ValueOf(nil), errArgType
			}
			return reflect.ValueOf(jsVal.String()), nil
		},
	}
	// goBaseValueToJSFuncs defined functions mapping for Go basic data types
	// value to JavaScript convention.
	goBaseValueToJSFuncs = map[reflect.Kind]func(goVal reflect.Value, kind reflect.Kind) (interface{}, error){
		reflect.Bool: func(goVal reflect.Value, kind reflect.Kind) (interface{}, error) {
			if kind != goVal.Kind() {
				return nil, errArgType
			}
			return goVal.Bool(), nil
		},
		reflect.Uint: func(goVal reflect.Value, kind reflect.Kind) (interface{}, error) {
			if kind != goVal.Kind() {
				return nil, errArgType
			}
			return int(goVal.Uint()), nil
		},
		reflect.Uint8: func(goVal reflect.Value, kind reflect.Kind) (interface{}, error) {
			if kind != goVal.Kind() {
				return nil, errArgType
			}
			return uint8(goVal.Uint()), nil
		},
		reflect.Uint64: func(goVal reflect.Value, kind reflect.Kind) (interface{}, error) {
			if kind != goVal.Kind() {
				return nil, errArgType
			}
			return int(goVal.Uint()), nil
		},
		reflect.Int: func(goVal reflect.Value, kind reflect.Kind) (interface{}, error) {
			if kind != goVal.Kind() {
				return nil, errArgType
			}
			return int(goVal.Int()), nil
		},
		reflect.Int64: func(goVal reflect.Value, kind reflect.Kind) (interface{}, error) {
			if kind != goVal.Kind() {
				return nil, errArgType
			}
			return goVal.Int(), nil
		},
		reflect.Float64: func(goVal reflect.Value, kind reflect.Kind) (interface{}, error) {
			if kind != goVal.Kind() {
				return nil, errArgType
			}
			return goVal.Float(), nil
		},
		reflect.String: func(goVal reflect.Value, kind reflect.Kind) (interface{}, error) {
			if kind != goVal.Kind() {
				return nil, errArgType
			}
			return goVal.String(), nil
		},
	}
	errArgNum  = errors.New("invalid arguments in call")
	errArgType = errors.New("invalid argument data type")
)

func main() {
	c := make(chan struct{})
	regFuncs()
	<-c
}

func regFuncs() {

	for name, impl := range map[string]func(this js.Value, args []js.Value) interface{}{
		"NewTable": NewTable,
		"NewHTTP":  NewHTTP,
	} {
		js.Global().Get("shyexcel").Set(name, js.FuncOf(impl))
	}
}

func NewTable(this js.Value, args []js.Value) interface{} {
	fn := map[string]interface{}{"error": nil}
	fn["error"] = nil
	if err := prepareArgs(args, []argsRule{
		{types: []js.Type{js.TypeObject}, opts: true},
	}); err != nil {
		fn["error"] = err.Error()
		return js.ValueOf(fn)
	}
	if len(args) == 1 {
		goVal, err := jsValueToGo(args[0], reflect.TypeOf(shyexcel.Table{}))
		if err != nil {
			fn["error"] = err.Error()
			return js.ValueOf(fn)
		}
		table := goVal.Elem().Interface().(shyexcel.Table)
		return regInteropFunc(shyexcel.NewTable(&table), fn)
	}
	fn["error"] = ""
	return js.ValueOf(fn)
}

func NewHTTP(this js.Value, args []js.Value) interface{} {
	fn := map[string]interface{}{"error": nil}
	fn["error"] = nil
	if err := prepareArgs(args, []argsRule{
		{types: []js.Type{js.TypeObject}, opts: true},
	}); err != nil {
		fn["error"] = err.Error()
		return js.ValueOf(fn)
	}
	if len(args) == 1 {
		table, err := shyexcel.NewHTTP(args[0].String(), "GET", nil)
		if err != nil {
			fn["error"] = err.Error()
			return js.ValueOf(fn)
		}
		return regInteropFunc(table, fn)
	} else if len(args) == 2 {
		table, err := shyexcel.NewHTTP(args[0].String(), args[1].String(), nil)
		if err != nil {
			fn["error"] = err.Error()
			return js.ValueOf(fn)
		}
		return regInteropFunc(table, fn)
	} else if len(args) == 3 {
		goVal, err := jsValueToGo(args[1], reflect.TypeOf(map[string]string{}))
		headers := goVal.Elem().Interface().(map[string]string)
		table, err := shyexcel.NewHTTP(args[0].String(), args[1].String(), func(header http.Header) {
			for k, v := range headers {
				header.Set(k, v)
			}
		})
		if err != nil {
			fn["error"] = err.Error()
			return js.ValueOf(fn)
		}
		return regInteropFunc(table, fn)
	}
	fn["error"] = "参数错误"
	return js.ValueOf(fn)
}

func regInteropFunc(f *excelize.File, fn map[string]interface{}) interface{} {
	for name, impl := range map[string]func(this js.Value, args []js.Value) interface{}{
		"WriteToBuffer": WriteToBuffer(f),
	} {
		fn[name] = js.FuncOf(impl)
	}
	return js.ValueOf(fn)
}

func prepareArgs(args []js.Value, types []argsRule) error {
	rules, arguments, opts := len(types), len(args), false
	if rules > 0 && types[rules-1].opts {
		opts = true
	}
	if (!opts && arguments != rules) || (opts && (arguments != rules && arguments+1 != rules)) {
		return errArgNum
	}
	for i := 0; i < len(types); i++ {
		if opts && i == arguments {
			return nil
		}
		excepted, received := types[i], args[i]
		if inTypeSlice(excepted.types, received.Type()) == -1 {
			return errArgType
		}
	}
	return nil
}

func inTypeSlice(a []js.Type, x js.Type) int {
	for idx, n := range a {
		if x == n {
			return idx
		}
	}
	return -1
}
func jsToGoBaseType(jsVal js.Value, kind reflect.Kind) (reflect.Value, error) {
	fn, ok := jsToBaseGoTypeFuncs[kind]
	if !ok {
		return reflect.ValueOf(nil), errArgType
	}
	return fn(jsVal, kind)
}

func jsValueToGo(jsVal js.Value, goType reflect.Type) (reflect.Value, error) {
	result := reflect.New(goType)
	s := result.Elem()

	for resultFieldIdx := 0; resultFieldIdx < s.NumField(); resultFieldIdx++ {
		field := goType.Field(resultFieldIdx)
		if goBaseTypes[field.Type.Kind()] {
			jsBaseVal := jsVal.Get(field.Name)
			if jsBaseVal.Type() != js.TypeUndefined {
				goBaseVal, err := jsToGoBaseType(jsBaseVal, field.Type.Kind())
				if err != nil {
					return result, err
				}
				s.Field(resultFieldIdx).Set(goBaseVal.Convert(s.Field(resultFieldIdx).Type()))
			}
			continue
		}
		switch field.Type.Kind() {
		case reflect.Ptr:
			// Pointer of the Go data type, for example: *excelize.Options or *string
			ptrType := field.Type.Elem()
			if !goBaseTypes[ptrType.Kind()] {
				// Pointer of the Go struct, for example: *excelize.Options
				jsObjVal := jsVal.Get(field.Name)
				if jsObjVal.Type() != js.TypeUndefined {
					if jsObjVal.Type() != js.TypeObject {
						return result, errArgType
					}
					v, err := jsValueToGo(jsObjVal, ptrType)
					if err != nil {
						return result, err
					}
					s.Field(resultFieldIdx).Set(v)
				}
			}
			if goBaseTypes[ptrType.Kind()] {
				// Pointer of the Go basic data type, for example: *string
				jsBaseVal := jsVal.Get(field.Name)
				if jsBaseVal.Type() != js.TypeUndefined {
					v, err := jsToGoBaseType(jsBaseVal, ptrType.Kind())
					if err != nil {
						return result, err
					}
					x := reflect.New(ptrType)
					x.Elem().Set(v)
					s.Field(resultFieldIdx).Set(x.Elem().Addr())
				}
			}
		case reflect.Struct:
			// The Go struct, for example: excelize.Options, convert sub fields recursively
			structType := field.Type
			jsObjVal := jsVal.Get(field.Name)
			if jsObjVal.Type() != js.TypeUndefined {
				if jsObjVal.Type() != js.TypeObject {
					return result, errArgType
				}
				v, err := jsValueToGo(jsObjVal, structType)
				if err != nil {
					return result, err
				}
				s.Field(resultFieldIdx).Set(v.Elem())
			}
		case reflect.Slice:
			// The Go data type array, for example:
			// []*excelize.Options, []excelize.Options, []string, []*string
			ele := field.Type.Elem()
			jsArray := jsVal.Get(field.Name)
			if jsArray.Type() != js.TypeUndefined {
				if jsArray.Type() != js.TypeObject {
					return result, errArgType
				}
				if ele.Kind() == reflect.Ptr {
					// Pointer array of the Go data type, for example: []*excelize.Options or []*string
					subEle := ele.Elem()
					for i := 0; i < jsArray.Length(); i++ {
						if goBaseTypes[subEle.Kind()] {
							// Pointer array of the Go basic data type, for example: []*string
							v, err := jsToGoBaseType(jsArray.Index(i), subEle.Kind())
							if err != nil {
								return result, err
							}
							x := reflect.New(subEle)
							x.Elem().Set(v)
							s.Field(resultFieldIdx).Set(reflect.Append(s.Field(resultFieldIdx), x.Elem().Addr()))
						} else {
							// Pointer array of the Go struct, for example: []*excelize.Options
							v, err := jsValueToGo(jsArray.Index(i), subEle)
							if err != nil {
								return result, err
							}
							x := reflect.New(subEle)
							x.Elem().Set(v.Elem())
							s.Field(resultFieldIdx).Set(reflect.Append(s.Field(resultFieldIdx), x.Elem().Addr()))
						}
					}
				} else {
					// The Go data type array, for example: []excelize.Options or []string
					subEle := ele
					for i := 0; i < jsArray.Length(); i++ {
						if subEle.Kind() == reflect.Uint8 { // []byte
							buf := make([]byte, jsArray.Length())
							js.CopyBytesToGo(buf, jsArray)
							s.Field(resultFieldIdx).Set(reflect.ValueOf(buf))
							break
						}
						if goBaseTypes[subEle.Kind()] {
							// The Go basic data type array, for example: []string
							v, err := jsToGoBaseType(jsArray.Index(i), subEle.Kind())
							if err != nil {
								return result, err
							}
							s.Field(resultFieldIdx).Set(reflect.Append(s.Field(resultFieldIdx), v))
						} else {
							// The Go struct array, for example: []excelize.Options
							v, err := jsValueToGo(jsArray.Index(i), subEle)
							if err != nil {
								return result, err
							}
							s.Field(resultFieldIdx).Set(reflect.Append(s.Field(resultFieldIdx), v.Elem()))
						}
					}
				}
			}
		}
	}
	return result, nil
}

func WriteToBuffer(f *excelize.File) func(this js.Value, args []js.Value) interface{} {
	return func(this js.Value, args []js.Value) interface{} {
		ret := map[string]interface{}{"buffer": js.ValueOf([]interface{}{}), "error": nil}
		err := prepareArgs(args, []argsRule{
			{types: []js.Type{js.TypeObject}, opts: true},
		})
		if err != nil {
			ret["error"] = err.Error()
			return js.ValueOf(ret)
		}
		var opts excelize.Options
		if len(args) == 1 {
			goVal, err := jsValueToGo(args[0], reflect.TypeOf(excelize.Options{}))
			if err != nil {
				ret["error"] = err.Error()
				return js.ValueOf(ret)
			}
			opts = goVal.Elem().Interface().(excelize.Options)
		}
		buf := new(bytes.Buffer)
		if err := f.Write(buf, opts); err != nil {
			ret["error"] = err.Error()
			return js.ValueOf(ret)
		}
		src := buf.Bytes()
		dst := js.Global().Get("Uint8Array").New(len(src))
		js.CopyBytesToJS(dst, src)
		ret["buffer"] = dst
		return js.ValueOf(ret)
	}
}
