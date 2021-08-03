package keydb

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type RedisHash struct {
	fields      map[string]reflect.Value
	RedisArgs   []interface{}
	EmptyFields []string
}

func (s *RedisHash) appendArgs(name string, val reflect.Value) error {
	switch val.Kind() {
	case reflect.Ptr:
		if val.IsNil() {
			s.EmptyFields = append(s.EmptyFields, name)
		} else {
			s.appendArgs(name, val.Elem())
		}
	case reflect.String:
		if val.String() == "" {
			s.EmptyFields = append(s.EmptyFields, name)
		} else {
			s.RedisArgs = append(s.RedisArgs, name, val.String())

		}
	case reflect.Bool:
		if val.Bool() {
			s.RedisArgs = append(s.RedisArgs, name, "true")
		} else {
			s.RedisArgs = append(s.RedisArgs, name, "false")

		}
	case reflect.Struct, reflect.Array, reflect.Slice, reflect.Map:
		b, err := json.Marshal(val.Interface())
		if err != nil {
			return err
		} else {
			s.RedisArgs = append(s.RedisArgs, name, string(b))
		}
	default:
		s.RedisArgs = append(s.RedisArgs, name, val.Interface())
	}
	return nil
}

func newStructFields(dst interface{}) (*RedisHash, error) {
	sf := new(RedisHash)
	sfv := reflect.ValueOf(dst)
	if sfv.Kind() != reflect.Ptr || sfv.IsNil() {
		return nil, fmt.Errorf("non-pointer %T", dst)
	}
	sfv = sfv.Elem()
	if sfv.Kind() != reflect.Struct {
		return nil, fmt.Errorf("non-struct %T", dst)
	}
	sft := sfv.Type()
	sf.fields = make(map[string]reflect.Value)
	for i := 0; i < sft.NumField(); i++ {
		tf := sft.Field(i)
		vf := sfv.Field(i)
		if tf.PkgPath != "" {
			continue
		}
		sf.fields[tf.Name] = vf
		err := sf.appendArgs(tf.Name, vf)
		if err != nil {
			return nil, err
		}
	}
	return sf, nil
}

func (s *RedisHash) fromStringStringMap(mp map[string]string) error {
	for k, v := range mp {
		if f, ok := s.fields[k]; ok {
			var ptr bool
			kind := f.Kind()
			if kind == reflect.Ptr {
				kind = f.Type().Elem().Kind()
				ptr = true
			}
			switch kind {
			case reflect.Float32, reflect.Float64:
				fl, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				if ptr {
					f.Set(reflect.ValueOf(&fl))
				} else {
					f.SetFloat(fl)
				}
			case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
				i, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				if ptr {
					f.Set(reflect.ValueOf(&i))
				} else {
					f.SetInt(i)
				}
			case reflect.String:
				if ptr {
					var ptrStr string = v
					f.Set(reflect.ValueOf(&ptrStr))
				} else {
					f.SetString(v)
				}
			case reflect.Bool:
				if ptr {
					var ptrBool bool = v == "true"
					f.Set(reflect.ValueOf(&ptrBool))
				} else {
					f.SetBool(v == "true")
				}
			case reflect.Struct, reflect.Array, reflect.Slice, reflect.Map:
				json.Unmarshal([]byte(v), f.Interface())
			default:
				return fmt.Errorf("undifinie type %s", kind)
			}
		}
	}
	return nil
}
