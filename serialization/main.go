package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

func toJson(i interface{}) string {
	byte ,_ := json.Marshal(i)
	return string(byte)
}

func formJson(jsonStr string, t interface{}) interface{} {
	var data []byte = []byte(jsonStr)
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil
	}
	return t
}

func toGob(i interface{}) string {
	buf := new(bytes.Buffer)
	gob.NewEncoder(buf).Encode(i)
	return string(buf.Bytes())
}



func main() {

}
