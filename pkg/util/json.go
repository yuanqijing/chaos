package util

import "encoding/json"

func ParseJson2Map(raw []byte) (map[string]interface{}, error) {
	ret := make(map[string]interface{})
	err := json.Unmarshal(raw, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
