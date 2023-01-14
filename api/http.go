package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/greg198584/gridclient/tools"
	"io/ioutil"
	"net/http"
	"time"
)

func RequestApi(method string, url string, data []byte) (result []byte, statusCode int, err error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	cli := &http.Client{
		Timeout: time.Duration(4) * time.Second,
	}
	resp, err := cli.Do(req)
	defer cli.CloseIdleConnections()
	if err != nil {
		return
	}
	result, err = ioutil.ReadAll(resp.Body)
	tools.Log(fmt.Sprintf("< request api status [%d] [%s %s] >", resp.StatusCode, method, url), "", false)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		jsonPretty, _ := tools.PrettyString(result)
		fmt.Println(jsonPretty)
		var results map[string]interface{}
		json.Unmarshal(result, &results)

		message := "erreur"
		if results["err_message"] != nil {
			message = results["err_message"].(string)
		}
		if results["Message"] != nil {
			message = results["Message"].(string)
		}
		tools.Error(fmt.Sprintf("erreur: %s", message))
	}
	return result, resp.StatusCode, err
}
