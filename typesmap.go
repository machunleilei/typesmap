package typesmap

import (
        "errors"
        "fmt"
        "reflect"
        "strings"
        "sync"
)

const (
        OptIgnore    = "-"
        OptOmitempty = "omitempty"
)

var (
        notFoundErr = errors.New("not found")
)

const (
        flagIgnore = 1 << iota
        flagOmiEmpty
)

type KvExtractor interface {
        GetString(string) (string, error)
        GetBool(string) (bool, error)
        GetInt64(string) (int64, error)
        GetUint64(string) (uint64, error)
        GetFloat64(string) (float64, error)
        GetComplex128(string) (complex128, error)
        GetInterface(string) (interface{}, error)
        GetBytes(string) ([]byte, error)
        GetStrings(string) ([]string, error)
        GetBools(string) ([]bool, error)
        GetInt64s(string) ([]int64, error)
        GetUint64s(string) ([]uint64, error)
        GetFloat64s(string) ([]float64, error)
        GetComplex128s(string) ([]complex128, error)
        GetKvExtractors(string) ([]KvExtractor, error)
}

func NewSimpleKvExtractor(tagName string) *SimpleKvExtractor {
        return &SimpleKvExtractor{
                tagName:   tagName,
                keyValues: make(map[string]interface{}),
        }
}

var _ KvExtractor = (*SimpleKvExtractor)(nil)

type SimpleKvExtractor struct {
        lock      sync.RWMutex
        keyValues map[string]interface{}
        tagName   string
}

func (s *SimpleKvExtractor) Put(key string, data interface{}) error {
        v := reflect.ValueOf(data)
        if v.Kind() == reflect.Ptr && v.IsNil() {
                return fmt.Errorf("%s is a nil pointer", v.Kind().String())
        }
        for v.Kind() == reflect.Ptr {
                v = v.Elem()
        }
        s.lock.Lock()
        defer s.lock.Unlock()
        s.keyValues[key] = data
        return s.convert2Maps(key, data)
}

func (s *SimpleKvExtractor) GetString(k string) (string, error) {
        k = strings.TrimSpace(k)
        s.lock.RLock()
        defer s.lock.RUnlock()
        if v, ok := s.keyValues[k]; ok {
                return v.(string), nil
        }
        return "", notFoundErr
}

func (s *SimpleKvExtractor) GetInt64(k string) (int64, error) {
        k = strings.TrimSpace(k)
        s.lock.RLock()
        defer s.lock.RUnlock()
        if v, ok := s.keyValues[k]; ok {
                return v.(int64), nil
        }
        return 0, notFoundErr
}

func (s *SimpleKvExtractor) GetUint64(k string) (uint64, error) {
        k = strings.TrimSpace(k)
        s.lock.RLock()
        defer s.lock.RUnlock()
        if v, ok := s.keyValues[k]; ok {
                return v.(uint64), nil
        }
        return 0, notFoundErr
}
func (s *SimpleKvExtractor) GetFloat64(k string) (float64, error) {
        k = strings.TrimSpace(k)
        s.lock.RLock()
        defer s.lock.RUnlock()
        if v, ok := s.keyValues[k]; ok {
                return v.(float64), nil
        }
        return 0, notFoundErr
}

func (s *SimpleKvExtractor) GetComplex128(k string) (complex128, error) {
        k = strings.TrimSpace(k)
        s.lock.RLock()
        defer s.lock.RUnlock()
        if v, ok := s.keyValues[k]; ok {
                return v.(complex128), nil
        }
        return 0, notFoundErr
}

func (s *SimpleKvExtractor) GetInterface(k string) (interface{}, error) {
        k = strings.TrimSpace(k)
        s.lock.RLock()
        defer s.lock.RUnlock()
        if v, ok := s.keyValues[k]; ok {
                return v, nil
        }
        return nil, notFoundErr
}

func (s *SimpleKvExtractor) GetBool(k string) (bool, error) {
        k = strings.TrimSpace(k)
        s.lock.RLock()
        defer s.lock.RUnlock()
        if v, ok := s.keyValues[k]; ok {
                return v.(bool), nil
        }
        return false, notFoundErr
}

func (s *SimpleKvExtractor) GetBytes(k string) ([]byte, error) {
        k = strings.TrimSpace(k)
        s.lock.RLock()
        defer s.lock.RUnlock()
        if v, ok := s.keyValues[k]; ok {
                return v.([]byte), nil
        }
        return nil, notFoundErr
}

func (s *SimpleKvExtractor) GetStrings(k string) ([]string, error) {
        k = strings.TrimSpace(k)
        s.lock.RLock()
        defer s.lock.RUnlock()
        if v, ok := s.keyValues[k]; ok {
                return v.([]string), nil
        }
        return nil, notFoundErr
}

func (s *SimpleKvExtractor) GetBools(k string) ([]bool, error) {
        k = strings.TrimSpace(k)
        s.lock.RLock()
        defer s.lock.RUnlock()
        if v, ok := s.keyValues[k]; ok {
                return v.([]bool), nil
        }
        return nil, notFoundErr
}

func (s *SimpleKvExtractor) GetInt64s(k string) ([]int64, error) {
        k = strings.TrimSpace(k)
        s.lock.RLock()
        defer s.lock.RUnlock()
        if v, ok := s.keyValues[k]; ok {
                return v.([]int64), nil
        }
        return nil, notFoundErr
}

func (s *SimpleKvExtractor) GetUint64s(k string) ([]uint64, error) {
        k = strings.TrimSpace(k)
        s.lock.RLock()
        defer s.lock.RUnlock()
        if v, ok := s.keyValues[k]; ok {
                return v.([]uint64), nil
        }
        return nil, notFoundErr
}

func (s *SimpleKvExtractor) GetFloat64s(k string) ([]float64, error) {
        k = strings.TrimSpace(k)
        s.lock.RLock()
        defer s.lock.RUnlock()
        if v, ok := s.keyValues[k]; ok {
                return v.([]float64), nil
        }
        return nil, notFoundErr
}

func (s *SimpleKvExtractor) GetComplex128s(k string) ([]complex128, error) {
        k = strings.TrimSpace(k)
        s.lock.RLock()
        defer s.lock.RUnlock()
        if v, ok := s.keyValues[k]; ok {
                return v.([]complex128), nil
        }
        return nil, notFoundErr
}

func (s *SimpleKvExtractor) GetKvExtractors(k string) ([]KvExtractor, error) {
        k = strings.TrimSpace(k)
        s.lock.RLock()
        defer s.lock.RUnlock()
        if v, ok := s.keyValues[k]; ok {
                return v.([]KvExtractor), nil
        }
        return nil, notFoundErr
}

func (s *SimpleKvExtractor) convert2Maps(key string, data interface{}) (err error) {
        v := reflect.ValueOf(data)
        for v.Kind() == reflect.Ptr {
                v = v.Elem()
        }
        optKey := func(v string) string {
                if len(key) == 0 {
                        return v
                }
                if len(v) == 0 {
                        return key
                }
                return key + "." + v
        }
        value := v
        switch v.Kind() {
        case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
                s.keyValues[key] = value.Int()
        case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
                s.keyValues[key] = value.Uint()
        case reflect.Float32, reflect.Float64:
                s.keyValues[key] = value.Float()
        case reflect.String:
                s.keyValues[key] = value.String()
        case reflect.Bool:
                s.keyValues[key] = value.Bool()
        case reflect.Complex64, reflect.Complex128:
                s.keyValues[key] = value.Complex()
        case reflect.Interface:
                s.keyValues[key] = value.Interface()
        case reflect.Chan:
                s.keyValues[key] = value
        case reflect.Map:
                mapKeys := value.MapKeys()
                for _, mapKey := range mapKeys {
                        newKey := optKey(mapKey.String())
                        if err = s.convert2Maps(newKey, value.MapIndex(mapKey).Interface()); err != nil {
                                return err
                        }
                }
        case reflect.Struct:
                for i := 0; i < value.NumField(); i++ {
                        fieldType := value.Type().Field(i)
                        // ignore unexported field
                        if fieldType.PkgPath != "" {
                                continue
                        }
                        var flag int
                        var tagVal string
                        if tagVal, flag = readTag(fieldType, s.tagName); len(tagVal) == 0 {
                                tagVal = fieldType.Name
                        }
                        if flag&flagIgnore != 0 {
                                continue
                        }
                        fieldValue := v.Field(i)
                        if flag&flagOmiEmpty != 0 && fieldValue.IsZero() {
                                continue
                        }
                        if fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil() {
                                continue
                        }
                        if err = s.convert2Maps(optKey(tagVal), fieldValue.Interface()); err != nil {
                                return err
                        }
                }
        case reflect.Slice, reflect.Array:
                if value.Len() == 0 {
                        delete(s.keyValues, key)
                        break
                }
                var arr interface{}
                if arr, err = convertArray(s.tagName, key, data); err != nil {
                        return err
                }
                s.keyValues[key] = arr
        }
        return nil
}

func convertArray(tagName, key string, data interface{}) (ret interface{}, err error) {
        v := reflect.ValueOf(data)
        for v.Kind() == reflect.Ptr {
                v = v.Elem()
        }
        value := v
        if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
                return nil, fmt.Errorf("kind:%v not slice or array", value.Kind())
        }
        if value.Len() == 0 {
                return nil, nil
        }
        size := value.Len()
        v1 := value.Index(0)
        for v1.Kind() == reflect.Ptr {
                v1 = v1.Elem()
        }
        switch v1.Kind() {
        case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
                arr := make([]int64, 0, size)
                for i := 0; i < size; i++ {
                        arr = append(arr, value.Index(i).Int())
                }
                ret = arr
        case reflect.Uint8:
                arr := make([]byte, 0, size)
                arr = append(arr, value.Bytes()...)
                ret = arr
        case reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
                arr := make([]uint64, 0, size)
                for i := 0; i < size; i++ {
                        arr = append(arr, value.Index(i).Uint())
                }
                ret = arr
        case reflect.Float32, reflect.Float64:
                arr := make([]float64, 0, size)
                for i := 0; i < size; i++ {
                        arr = append(arr, value.Index(i).Float())
                }
                ret = arr
        case reflect.String:
                arr := make([]string, 0, size)
                for i := 0; i < size; i++ {
                        arr = append(arr, value.Index(i).String())
                }
                ret = arr
        case reflect.Bool:
                arr := make([]bool, 0, size)
                for i := 0; i < size; i++ {
                        arr = append(arr, value.Index(i).Bool())
                }
                ret = arr
        case reflect.Complex64, reflect.Complex128:
                arr := make([]complex128, 0, size)
                for i := 0; i < size; i++ {
                        arr = append(arr, value.Index(i).Complex())
                }
                ret = arr
        case reflect.Interface:
                arr := make([]interface{}, 0, size)
                for i := 0; i < size; i++ {
                        arr = append(arr, value.Index(i).Interface())
                }
                ret = arr
        case reflect.Chan:
                return nil, errors.New("has not supportted yet")
        case reflect.Map, reflect.Struct:
                arr := make([]KvExtractor, 0, value.Len())
                for i := 0; i < value.Len(); i++ {
                        ss := NewSimpleKvExtractor(tagName)
                        if err = ss.convert2Maps("", value.Index(i).Interface()); err != nil {
                                return nil, err
                        }
                        arr = append(arr, ss)
                }
                ret = arr
        }
        return
}

func readTag(f reflect.StructField, tag string) (string, int) {
        val, ok := f.Tag.Lookup(tag)
        fieldTag := ""
        flag := 0

        // no tag, skip this field
        if !ok {
                return "", flag
        }
        opts := strings.Split(val, ",")

        fieldTag = opts[0]
        for i := 0; i < len(opts); i++ {
                switch opts[i] {
                case OptIgnore:
                        flag |= flagIgnore
                case OptOmitempty:
                        flag |= flagOmiEmpty
                }
        }

        return fieldTag, flag
}
