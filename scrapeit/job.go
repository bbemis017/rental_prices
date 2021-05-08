package scrapeit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/bbemis017/ApartmentNotifier/util"
)

type (
	JobStruct struct {
		templateId int
		cacheOn    bool
		taskId     string `json:"id"`
		status     string
		data       map[string]interface{}
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
	if job.status == "SUCCESS" || job.status == "ERROR" || job.status == "FAILURE" {
		return job.status, nil
	}

	bytes, err := sendRequestBytes(
		"GET",
		"job/"+job.taskId+"/status",
		map[string]string{},
	)

	if err != nil {
		return "", err
	}

	var v interface{}
	json.Unmarshal(bytes, &v)
	result_map := v.(map[string]interface{})

	job.status = result_map["state"].(string)

	if job.status == "SUCCESS" {
		job.data = result_map["data"].(map[string]interface{})
	}

	return job.status, nil
}

func (job *JobStruct) AwaitResult() (map[string]interface{}, error) {
	fmt.Println("Wait for Results")

	status, _ := job.Status()
	fmt.Println(status)

	for status != "SUCCESS" && status != "ERROR" && status != "FAILURE" {
		time.Sleep(3 * time.Second)

		status, _ = job.Status()
		fmt.Println(status)
	}

	if status == "SUCCESS" {
		return job.data, nil
	} else {
		return nil, errors.New("error occurred Status: " + job.status)
	}
}

func (job *JobStruct) GetResult() (map[string]interface{}, error) {
	if job.status == "SUCCESS" {
		return job.data, nil
	}
	return nil, errors.New("No Result Available, Status " + job.status)
}

func sendRequestBytes(method string, url string, body map[string]string) ([]byte, error) {
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
		return nil, err
	}

	// if request does not return a 200 status code like 200 or 202
	if !strings.HasPrefix(res.Status, "20") {
		return nil, errors.New("HTTP Request returned" + res.Status)
	}

	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)
	return bytes, nil
}

func send_request(method string, url string, body map[string]string, result_struct interface{}) error {

	// Unmarshal results into result struct
	bytes, err := sendRequestBytes(method, url, body)
	if err != nil {
		return err
	}

	json.Unmarshal(bytes, &result_struct)
	return nil
}
