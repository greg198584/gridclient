package tools

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
)

func PrettyString(str []byte) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, str, "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
func GetValue(result []byte, index string) string {
	var results map[string]interface{}
	json.Unmarshal(result, &results)
	return results[index].(string)
}
func CreateJsonFile(filepath string, data any) (err error) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return
	}
	err = ioutil.WriteFile(filepath, file, 0644)
	return
}
func GetJsonFile(filename string) (file []byte, err error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return file, err
	}
	jsonFile, err := os.Open(filename)
	if err != nil {
		return file, err
	}
	file, err = ioutil.ReadAll(jsonFile)
	err = jsonFile.Close()
	return file, err
}
func AnyToByte(data any) []byte {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(data)
	return reqBodyBytes.Bytes()
}
func AnyToString(data any) string {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(data)
	return reqBodyBytes.String()
}
