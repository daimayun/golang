package conv

import "encoding/json"

// InterfaceToString interface转string
func InterfaceToString(i interface{}) (str string, err error) {
	switch i.(type) {
	case string:
		str = i.(string)
	case []byte:
		str = string(i.([]byte))
	case int:
		str = IntToString(i.(int))
	case int8:
		str = Int64ToString(Int8ToInt64(i.(int8)))
	case int16:
		str = Int64ToString(Int16ToInt64(i.(int16)))
	case int32:
		str = Int64ToString(Int32ToInt64(i.(int32)))
	case int64:
		str = Int64ToString(i.(int64))
	case uint:
		str = IntToString(UintToInt(i.(uint)))
	case uint8:
		str = Int64ToString(Uint8ToInt64(i.(uint8)))
	case uint16:
		str = Int64ToString(Uint16ToInt64(i.(uint16)))
	case uint32:
		str = Int64ToString(Uint32ToInt64(i.(uint32)))
	case uint64:
		str = Int64ToString(Uint64ToInt64(i.(uint64)))
	default:
		var jsonByte []byte
		jsonByte, err = json.Marshal(i)
		if err != nil {
			return
		}
		str = string(jsonByte)
	}
	return
}
