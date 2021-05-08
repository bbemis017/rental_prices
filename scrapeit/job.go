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
		taskId     string `json:"id"`
		status     string
	}
)

func NewJob(templateId int, cacheOn bool) JobStruct {
	return JobStruct{
		templateId: templateId,
		cacheOn:    cacheOn,
	}
}

func (job *JobStruct) Start() (string, error) {
	fmt.Println("Start")

	type Results struct {
		TaskId string `json:"id"`
	}
	var results Results

	err := send_request(
		"POST",
		"job/"+fmt.Sprint(job.templateId),
		map[string]string{},
		&results,
	)
	if err != nil {
		return "", err
	}

	job.taskId = results.TaskId

	return job.taskId, nil
}

func (job *JobStruct) Status() (string, error) {
	if job.status == "SUCCESS" || job.status == "ERROR" {
		return job.status, nil
	}

	type Results struct {
		State string `json:"state"`
	}
	var results Results

	err := send_request(
		"GET",
		"job/"+job.taskId+"/status",
		map[string]string{},
		&results,
	)
	if err != nil {
		return "", err
	}

	job.status = results.State

	return job.status, nil
}

func GetResult() (string, error) {
	return "", nil
}

func send_request(method string, url string, body map[string]string, result_struct interface{}) error {
	// Encode Object into json bytes
	encoded_body, _ := json.Marshal(body)

	path := util.GetEnvOrFail(util.ENV_SCRAPEIT_NET_HOST) + url

	req, _ := http.NewRequest(
		method,
		path,
		bytes.NewBuffer(encoded_body),
	)

	req.Header.Set("Authorization", "KEY "+util.GetEnvOrFail(util.ENV_SCRAPEIT_NET_KEY))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	// if request does not return a 200 status code like 200 or 202
	if !strings.HasPrefix(res.Status, "20") {
		return errors.New("HTTP Request returned" + res.Status)
	}

	defer res.Body.Close()

	// Unmarshal results into result struct
	bytes, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(bytes, &result_struct)

	return nil
}
