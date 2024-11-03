package helpers

import (
    "encoding/json"
)

// ToJSON converts a struct to a JSON string
func ToJSON(data interface{}) (string, error) {
    jsonBytes, err := json.Marshal(data)
    if err != nil {
        return "", err
    }
    return string(jsonBytes), nil
}

// FromJSON converts a JSON string to a struct
func FromJSON(jsonStr string, result interface{}) error {
    return json.Unmarshal([]byte(jsonStr), result)
}