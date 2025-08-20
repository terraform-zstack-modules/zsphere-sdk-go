// Copyright (c) ZStack.io, Inc.

package jsonutils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/golog"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/errors"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/util/gotypes"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/util/reflectutils"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/util/sortedmap"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/util/timeutils"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/util/tristate"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/util/utils"
)

func (th *JSONValue) Unmarshal(obj interface{}, keys ...string) error {
	return jsonUnmarshal(th, obj, keys)
}

func (th *JSONArray) Unmarshal(obj interface{}, keys ...string) error {
	return jsonUnmarshal(th, obj, keys)
}

func (th *JSONDict) Unmarshal(obj interface{}, keys ...string) error {
	return jsonUnmarshal(th, obj, keys)
}

func jsonUnmarshal(jo JSONObject, o interface{}, keys []string) error {
	if len(keys) > 0 {
		var err error = nil
		jo, err = jo.Get(keys...)
		if err != nil {
			return errors.Wrap(err, "Get")
		}
	}
	value := reflect.ValueOf(o)
	err := jo.unmarshalValue(reflect.Indirect(value))
	if err != nil {
		return errors.Wrap(err, "jo.unmarshalValue")
	}
	return nil
}

func (th *JSONValue) unmarshalValue(val reflect.Value) error {
	if val.CanSet() {
		zeroVal := reflect.New(val.Type()).Elem()
		val.Set(zeroVal)
	}
	return nil
}

func (th *JSONInt) unmarshalValue(val reflect.Value) error {
	switch val.Type() {
	case JSONIntType:
		json := val.Interface().(JSONInt)
		json.data = th.data
		return nil
	case JSONIntPtrType, JSONObjectType:
		val.Set(reflect.ValueOf(th))
		return nil
	case JSONStringType:
		json := val.Interface().(JSONString)
		json.data = fmt.Sprintf("%d", th.data)
		return nil
	case JSONStringPtrType:
		json := val.Interface().(*JSONString)
		data := fmt.Sprintf("%d", th.data)
		if json == nil {
			json = NewString(data)
			val.Set(reflect.ValueOf(json))
		} else {
			json.data = data
		}
		return nil
	case JSONBoolType, JSONFloatType, JSONArrayType, JSONDictType, JSONBoolPtrType, JSONFloatPtrType, JSONArrayPtrType, JSONDictPtrType:
		return ErrTypeMismatch // fmt.Errorf("JSONInt type mismatch %s", val.Type())
	case tristate.TriStateType:
		if th.data == 0 {
			val.Set(tristate.TriStateFalseValue)
		} else {
			val.Set(tristate.TriStateTrueValue)
		}
	}
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		val.SetInt(th.data)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		val.SetUint(uint64(th.data))
	case reflect.Float32, reflect.Float64:
		val.SetFloat(float64(th.data))
	case reflect.Bool:
		if th.data == 0 {
			val.SetBool(false)
		} else {
			val.SetBool(true)
		}
	case reflect.String:
		val.SetString(fmt.Sprintf("%d", th.data))
	case reflect.Ptr:
		if val.IsNil() {
			val.Set(reflect.New(val.Type().Elem()))
		}
		return th.unmarshalValue(val.Elem())
	case reflect.Interface:
		val.Set(reflect.ValueOf(th.data))
	default:
		return errors.Wrapf(ErrTypeMismatch, "JSONInt vs. %s", val.Type())
	}
	return nil
}

func (th *JSONBool) unmarshalValue(val reflect.Value) error {
	switch val.Type() {
	case JSONBoolType:
		json := val.Interface().(JSONBool)
		json.data = th.data
		return nil
	case JSONBoolPtrType, JSONObjectType:
		val.Set(reflect.ValueOf(th))
		return nil
	case JSONStringType:
		json := val.Interface().(JSONString)
		json.data = strconv.FormatBool(th.data)
		return nil
	case JSONStringPtrType:
		json := val.Interface().(*JSONString)
		data := strconv.FormatBool(th.data)
		if json == nil {
			json = NewString(data)
			val.Set(reflect.ValueOf(json))
		} else {
			json.data = data
		}
		return nil
	case JSONIntType, JSONFloatType, JSONArrayType, JSONDictType, JSONIntPtrType, JSONFloatPtrType, JSONArrayPtrType, JSONDictPtrType:
		return ErrTypeMismatch // fmt.Errorf("JSONBool type mismatch %s", val.Type())
	case tristate.TriStateType:
		if th.data {
			val.Set(tristate.TriStateTrueValue)
		} else {
			val.Set(tristate.TriStateFalseValue)
		}
	}
	switch val.Kind() {
	case reflect.Int, reflect.Uint, reflect.Int8, reflect.Uint8,
		reflect.Int16, reflect.Uint16, reflect.Int32, reflect.Uint32, reflect.Int64, reflect.Uint64:
		if th.data {
			val.SetInt(1)
		} else {
			val.SetInt(0)
		}
	case reflect.Float32, reflect.Float64:
		if th.data {
			val.SetFloat(1.0)
		} else {
			val.SetFloat(0.0)
		}
	case reflect.Bool:
		val.SetBool(th.data)
	case reflect.String:
		if th.data {
			val.SetString("true")
		} else {
			val.SetString("false")
		}
	case reflect.Ptr:
		if val.IsNil() {
			val.Set(reflect.New(val.Type().Elem()))
		}
		return th.unmarshalValue(val.Elem())
	case reflect.Interface:
		val.Set(reflect.ValueOf(th.data))
	default:
		return errors.Wrapf(ErrTypeMismatch, "JSONBool vs. %s", val.Type())
	}
	return nil
}

func (th *JSONFloat) unmarshalValue(val reflect.Value) error {
	switch val.Type() {
	case JSONFloatType:
		json := val.Interface().(JSONFloat)
		json.data = th.data
		return nil
	case JSONFloatPtrType, JSONObjectType:
		val.Set(reflect.ValueOf(th))
		return nil
	case JSONStringType:
		json := val.Interface().(JSONString)
		json.data = fmt.Sprintf("%f", th.data)
		return nil
	case JSONStringPtrType:
		json := val.Interface().(*JSONString)
		data := fmt.Sprintf("%f", th.data)
		if json == nil {
			json = NewString(data)
			val.Set(reflect.ValueOf(json))
		} else {
			json.data = data
		}
		return nil
	case JSONIntType:
		json := val.Interface().(JSONInt)
		json.data = int64(th.data)
		return nil
	case JSONIntPtrType:
		json := val.Interface().(*JSONInt)
		if json == nil {
			json = NewInt(int64(th.data))
			val.Set(reflect.ValueOf(json))
		} else {
			json.data = int64(th.data)
		}
		return nil
	case JSONBoolType:
		json := val.Interface().(JSONBool)
		json.data = (int(th.data) != 0)
		return nil
	case JSONArrayType, JSONDictType, JSONBoolPtrType, JSONArrayPtrType, JSONDictPtrType:
		return ErrTypeMismatch // fmt.Errorf("JSONFloat type mismatch %s", val.Type())
	case tristate.TriStateType:
		if int(th.data) == 0 {
			val.Set(tristate.TriStateFalseValue)
		} else {
			val.Set(tristate.TriStateTrueValue)
		}
	}
	switch val.Kind() {
	case reflect.Int, reflect.Uint, reflect.Int8, reflect.Uint8,
		reflect.Int16, reflect.Uint16, reflect.Int32, reflect.Uint32, reflect.Int64, reflect.Uint64:
		val.SetInt(int64(th.data))
	case reflect.Float32, reflect.Float64:
		val.SetFloat(th.data)
	case reflect.Bool:
		if th.data == 0 {
			val.SetBool(false)
		} else {
			val.SetBool(true)
		}
	case reflect.String:
		val.SetString(fmt.Sprintf("%f", th.data))
	case reflect.Ptr:
		if val.IsNil() {
			val.Set(reflect.New(val.Type().Elem()))
		}
		return th.unmarshalValue(val.Elem())
	case reflect.Interface:
		val.Set(reflect.ValueOf(th.data))
	default:
		return errors.Wrapf(ErrTypeMismatch, "JSONFloat vs. %s", val.Type())
	}
	return nil
}

func (th *JSONString) unmarshalValue(val reflect.Value) error {
	switch val.Type() {
	case JSONStringType:
		json := val.Interface().(JSONString)
		json.data = th.data
		return nil
	case JSONStringPtrType, JSONObjectType:
		val.Set(reflect.ValueOf(th))
		return nil
	case gotypes.TimeType:
		var tm time.Time
		var err error
		if len(th.data) > 0 {
			tm, err = timeutils.ParseTimeStr(th.data)
			if err != nil {
				golog.Warnf("timeutils.ParseTimeStr %s %s", th.data, err)
			}
		} else {
			tm = time.Time{}
		}
		val.Set(reflect.ValueOf(tm))
		return nil
	case JSONBoolType:
		json := val.Interface().(JSONBool)
		switch strings.ToLower(th.data) {
		case "true", "yes", "on", "1":
			json.data = true
		default:
			json.data = false
		}
		return nil
	case JSONBoolPtrType:
		json := val.Interface().(*JSONBool)
		var data bool
		switch strings.ToLower(th.data) {
		case "true", "yes", "on", "1":
			data = true
		default:
			data = false
		}
		if json == nil {
			json = &JSONBool{data: data}
		} else {
			json.data = data
		}
		return nil
	case JSONIntType, JSONFloatType, JSONArrayType, JSONDictType,
		JSONBoolPtrType, JSONIntPtrType, JSONFloatPtrType, JSONArrayPtrType, JSONDictPtrType:
		return ErrTypeMismatch // fmt.Errorf("JSONString type mismatch %s", val.Type())
	case tristate.TriStateType:
		switch strings.ToLower(th.data) {
		case "true", "yes", "on", "1":
			val.Set(tristate.TriStateTrueValue)
		case "false", "no", "off", "0":
			val.Set(tristate.TriStateFalseValue)
		default:
			val.Set(tristate.TriStateNoneValue)
		}
	}
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if len(th.data) > 0 {
			intVal, err := strconv.ParseInt(normalizeCurrencyString(th.data), 10, 64)
			if err != nil {
				return err
			}
			val.SetInt(intVal)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if len(th.data) > 0 {
			intVal, err := strconv.ParseUint(normalizeCurrencyString(th.data), 10, 64)
			if err != nil {
				return err
			}
			val.SetUint(intVal)
		}
	case reflect.Float32, reflect.Float64:
		if len(th.data) > 0 {
			floatVal, err := strconv.ParseFloat(normalizeCurrencyString(th.data), 64)
			if err != nil {
				return err
			}
			val.SetFloat(floatVal)
		}
	case reflect.Bool:
		val.SetBool(utils.ToBool(th.data))
	case reflect.String:
		val.SetString(th.data)
	case reflect.Ptr:
		if val.IsNil() {
			val.Set(reflect.New(val.Type().Elem()))
		}
		return th.unmarshalValue(val.Elem())
	case reflect.Interface:
		val.Set(reflect.ValueOf(th.data))
	case reflect.Slice:
		dataLen := 1
		if val.Cap() < dataLen {
			newVal := reflect.MakeSlice(val.Type(), dataLen, dataLen)
			val.Set(newVal)
		} else if val.Len() != dataLen {
			val.SetLen(dataLen)
		}
		return th.unmarshalValue(val.Index(0))
	default:
		return errors.Wrapf(ErrTypeMismatch, "JSONString vs. %s", val.Type())
	}
	return nil
}

func (th *JSONArray) unmarshalValue(val reflect.Value) error {
	switch val.Type() {
	case JSONArrayType:
		array := val.Interface().(JSONArray)
		if th.data != nil {
			array.Add(th.data...)
		}
		val.Set(reflect.ValueOf(array))
		return nil
	case JSONArrayPtrType, JSONObjectType:
		val.Set(reflect.ValueOf(th))
		return nil
	case JSONDictType, JSONIntType, JSONStringType, JSONBoolType, JSONFloatType,
		JSONDictPtrType, JSONIntPtrType, JSONStringPtrType, JSONBoolPtrType, JSONFloatPtrType:
		return ErrTypeMismatch //fmt.Errorf("JSONArray type mismatch %s", val.Type())
	}
	switch val.Kind() {
	case reflect.String:
		val.SetString(th.String())
		return nil
	case reflect.Ptr:
		kind := val.Type().Elem().Kind()
		if kind == reflect.Array || kind == reflect.Slice {
			if val.IsNil() {
				val.Set(reflect.New(val.Type().Elem()))
			}
			return th.unmarshalValue(val.Elem())
		}
		return ErrTypeMismatch // fmt.Errorf("JSONArray type mismatch %s", val.Type())
	case reflect.Interface:
		val.Set(reflect.ValueOf(th.data))
	case reflect.Slice, reflect.Array:
		if val.Kind() == reflect.Array {
			if val.Len() != len(th.data) {
				return ErrArrayLengthMismatch // fmt.Errorf("JSONArray length unmatch %s: %d != %d", val.Type(), val.Len(), len(th.data))
			}
		} else if val.Kind() == reflect.Slice {
			dataLen := len(th.data)
			if val.Cap() < dataLen {
				newVal := reflect.MakeSlice(val.Type(), dataLen, dataLen)
				val.Set(newVal)
			} else if val.Len() != dataLen {
				val.SetLen(dataLen)
			}
		}
		for i, json := range th.data {
			err := json.unmarshalValue(val.Index(i))
			if err != nil {
				return errors.Wrap(err, "unmarshalValue")
			}
		}
	default:
		return errors.Wrapf(ErrTypeMismatch, "JSONArray vs. %s", val.Type())
	}
	return nil
}

func (th *JSONDict) unmarshalValue(val reflect.Value) error {
	switch val.Type() {
	case JSONDictType:
		dict := val.Interface().(JSONDict)
		dict.Update(th)
		val.Set(reflect.ValueOf(dict))
		return nil
	case JSONDictPtrType, JSONObjectType:
		val.Set(reflect.ValueOf(th))
		return nil
	case JSONArrayType, JSONIntType, JSONBoolType, JSONFloatType, JSONStringType,
		JSONArrayPtrType, JSONIntPtrType, JSONBoolPtrType, JSONFloatPtrType, JSONStringPtrType:
		return ErrTypeMismatch // fmt.Errorf("JSONDict type mismatch %s", val.Type())
	}
	switch val.Kind() {
	case reflect.String:
		val.SetString(th.String())
		return nil
	case reflect.Map:
		return th.unmarshalMap(val)
	case reflect.Struct:
		return th.unmarshalStruct(val)
	case reflect.Interface:
		if val.Type().Implements(gotypes.ISerializableType) {
			objPtr, err := gotypes.NewSerializable(val.Type())
			if err != nil {
				return err
			}
			if objPtr == nil {
				val.Set(reflect.ValueOf(th.data)) // ???
				return nil
			}
			err = th.unmarshalValue(reflect.ValueOf(objPtr))
			if err != nil {
				return errors.Wrap(err, "unmarshalValue")
			}
			//
			// XXX
			//
			// cannot unmarshal nested anonymous interface
			// as nested anonymous interface is treated as a named field
			// please use jsonutils.Deserialize to descrialize such interface
			// ...
			// objPtr = gotypes.Transform(val.Type(), objPtr)
			//
			val.Set(reflect.ValueOf(objPtr).Convert(val.Type()))
		} else {
			return errors.Wrapf(ErrInterfaceUnsupported, "JSONDict.unmarshalValue: %s", val.Type())
		}
	case reflect.Ptr:
		kind := val.Type().Elem().Kind()
		if kind == reflect.Struct || kind == reflect.Map {
			if val.IsNil() {
				newVal := reflect.New(val.Type().Elem())
				val.Set(newVal)
			}
			return th.unmarshalValue(val.Elem())
		}
		fallthrough
	default:
		return errors.Wrapf(ErrTypeMismatch, "JSONDict.unmarshalValue: %s", val.Type())
	}
	return nil
}

func (th *JSONDict) unmarshalMap(val reflect.Value) error {
	if val.IsNil() {
		mapVal := reflect.MakeMap(val.Type())
		val.Set(mapVal)
	}
	valType := val.Type()
	keyType := valType.Key()
	if keyType.Kind() != reflect.String {
		return ErrMapKeyMustString // fmt.Errorf("map key must be string")
	}
	for iter := sortedmap.NewIterator(th.data); iter.HasMore(); iter.Next() {
		k, vinf := iter.Get()
		v := vinf.(JSONObject)
		keyVal := reflect.ValueOf(k)
		if keyType != keyVal.Type() {
			keyVal = keyVal.Convert(keyType)
		}
		valVal := reflect.New(valType.Elem()).Elem()

		err := v.unmarshalValue(valVal)
		if err != nil {
			return errors.Wrap(err, "JSONDict.unmarshalMap")
		}
		val.SetMapIndex(keyVal, valVal)
	}
	return nil
}

func setStructFieldAt(key string, v JSONObject, fieldValues reflectutils.SStructFieldValueSet, keyIndexMap map[string][]int, visited map[string]bool) error {
	if visited == nil {
		visited = make(map[string]bool)
	}
	if _, ok := visited[key]; ok {
		// reference loop detected
		return nil
	}
	visited[key] = true
	indexes, ok := keyIndexMap[key]
	if !ok || len(indexes) == 0 {
		// try less strict match name
		indexes = fieldValues.GetStructFieldIndexes2(key, false)
		if len(indexes) == 0 {
			// no field match k, ignore
			return nil
		}
	}
	for _, index := range indexes {
		err := v.unmarshalValue(fieldValues[index].Value)
		if err != nil {
			return errors.Wrap(err, "JSONDict.unmarshalStruct")
		}
		depInfo, ok := fieldValues[index].Info.Tags[TagDeprecatedBy]
		if ok {
			err := setStructFieldAt(depInfo, v, fieldValues, keyIndexMap, visited)
			if err != nil {
				return errors.Wrap(err, "setStructFieldAt")
			}
		}
	}
	return nil
}

func (th *JSONDict) unmarshalStruct(val reflect.Value) error {
	fieldValues := reflectutils.FetchStructFieldValueSetForWrite(val)
	keyIndexMap := fieldValues.GetStructFieldIndexesMap()
	errs := make([]error, 0)
	for iter := sortedmap.NewIterator(th.data); iter.HasMore(); iter.Next() {
		k, vinf := iter.Get()
		v := vinf.(JSONObject)
		err := setStructFieldAt(k, v, fieldValues, keyIndexMap, nil)
		if err != nil {
			// store error, not interrupt the process
			errs = append(errs, errors.Wrapf(err, "setStructFieldAt %s: %s", k, v))
		}
	}
	callStructAfterUnmarshal(val)
	if len(errs) > 0 {
		return errors.NewAggregate(errs)
	} else {
		return nil
	}
}

func callStructAfterUnmarshal(val reflect.Value) {
	switch val.Kind() {
	case reflect.Struct:
		structType := val.Type()
		for i := 0; i < val.NumField(); i++ {
			fieldType := structType.Field(i)
			if fieldType.Anonymous {
				callStructAfterUnmarshal(val.Field(i))
			}
		}
		valPtr := val.Addr()
		afterMarshalFunc := valPtr.MethodByName("AfterUnmarshal")
		if afterMarshalFunc.IsValid() && !afterMarshalFunc.IsNil() {
			afterMarshalFunc.Call([]reflect.Value{})
		}
	case reflect.Ptr:
		callStructAfterUnmarshal(val.Elem())
	}
}
