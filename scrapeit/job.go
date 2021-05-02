package scrapeit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/bbemis017/ApartmentNotifier/util"
)

type (
	JobStruct struct {
		templateId int
		cacheOn    bool
		TaskId     string `json:"id"`
	}
)

func NewJob(templateId int, cacheOn bool) JobStruct {
	return JobStruct{
		templateId: templateId,
		cacheOn:    cacheOn,
	}
}

func (job JobStruct) Start() (string, error) {
	fmt.Println("Start")

	requestBody, _ := json.Marshal(map[string]string{})

	client := &http.Client{}
	req, _ := http.NewRequest(
		"POST",
		util.GetEnvOrFail(util.ENV_SCRAPEIT_NET_HOST)+"job/"+fmt.Sprint(job.templateId),
		bytes.NewBuffer(requestBody),
	)

	req.Header.Set("Authorization", "KEY "+util.GetEnvOrFail(util.ENV_SCRAPEIT_NET_KEY))
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if !strings.HasPrefix(res.Status, "202") {
		return "", errors.New("HTTP Request returned" + res.Status)
	}

	defer res.Body.Close()

	bytes, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(bytes, &job)

	return job.TaskId, nil
}

func Status() (string, error) {
	return "", nil
}

func GetResult() (string, error) {
	return "", nil
}
