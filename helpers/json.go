package helpers

import (
    "encoding/json"
)

func ToJSON(data interface{}) (string, error) {
    jsonBytes, err := json.Marshal(data)
    if err != nil {
        return "", err
    }
    return string(jsonBytes), nil
}

func FromJSON(jsonStr string, result interface{}) error {
    return json.Unmarshal([]byte(jsonStr), result)
}