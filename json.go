package rr

import (
    "bytes"
    "encoding/json"
)

var (
    JsonMarshalAdapter   = json.Marshal
    JsonUnmarshalAdapter = json.Unmarshal
)

func JsonUnSerialize(d []byte, v interface{}) error {
    return JsonUnmarshal(d, v)
}
func JsonUnmarshal(d []byte, v interface{}) error {
    return JsonUnmarshalAdapter(d, v)
}
func JsonSerialize(v interface{}) string {
    return JsonMarshal(v)
}
func JsonSerializeAsBytes(v interface{}) []byte {
    return JsonMarshalAsBytes(v)
}
func JsonMarshalAsBytes(v interface{}) []byte {
    adapter, _ := JsonMarshalAdapter(v)
    return adapter
}
func JsonMarshal(v interface{}) string {
    return string(JsonMarshalAsBytes(v))
}

func JsonMarshalAsReader(v interface{}) *bytes.Reader {
    return bytes.NewReader(JsonMarshalAsBytes(v))
}
func JsonSerializeAsReader(v interface{}) *bytes.Reader {
    return JsonMarshalAsReader(v)
}