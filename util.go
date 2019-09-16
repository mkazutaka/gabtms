package btcwsc

import "encoding/json"

func UnmarshalDataString(data json.RawMessage, dest interface{}) error {
	var dataStr string
	err := json.Unmarshal(data, &dataStr)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(dataStr), dest)
}
