package scrapeit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bbemis017/ApartmentNotifier/util"
)

// Represents a ScrapeIt.net Job
type JobStruct struct {
	templateId int                    // Id number of template that is being used for the job
	cacheOn    bool                   // If true the rendered html will be cached
	taskId     string                 //id of started job
	status     string                 // current status of job
	data       map[string]interface{} // raw results of job
}

// Job Status
type Status string

const (
	STATUS_SUCCESS   = "SUCCESS"
	STATUS_ERROR     = "ERROR"
	STATUS_FAILURE   = "FAILURE"
	STATUS_PARSING   = "PARSING"
	STATUS_RENDERING = "RENDERING"
	STATUS_LOADING   = "LOADING"
	STATUS_PENDING   = "PENDING"
)

// Creates a new Job Object for the provided template id
// Parameters
//    - templateId, id of template on ScrapeIt.net
//    - cacheOn, TODO this option has not been implemented
func NewJob(templateId int, cacheOn bool) JobStruct {
	return JobStruct{
		templateId: templateId,
		cacheOn:    cacheOn, //TODO this option has not been implemented
	}
}

// Starts a Job for the template
func (job *JobStruct) Start() (string, error) {
	fmt.Println("Start")

	body := map[string]string{
		"cache_on": strconv.FormatBool(job.cacheOn),
	}

	result_map, err := sendRequest(
		"POST",
		"job/"+fmt.Sprint(job.templateId),
		body,
	)
	if err != nil {
		return "", err
	}

	job.taskId = result_map["id"].(string)

	return job.taskId, nil
}

// Queries the Status of the Job from the API
func (job *JobStruct) Status() (string, error) {
	if isFinalState(job.status) {
		return job.status, nil
	}

	result_map, err := sendRequest(
		"GET",
		"job/"+job.taskId+"/status",
		map[string]string{},
	)

	if err != nil {
		return "", err
	}

	job.status = result_map["state"].(string)

	if job.status == STATUS_SUCCESS {
		job.data = result_map["data"].(map[string]interface{})
	}

	return job.status, nil
}

// Waits until the job has completed and returns the data
func (job *JobStruct) AwaitResult() (map[string]interface{}, error) {
	fmt.Println("Wait for Results")

	status, err := job.Status()
	job.status = status
	if err != nil {
		return nil, err
	}

	for !isFinalState(job.status) {
		time.Sleep(3 * time.Second)

		status, err2 := job.Status()
		job.status = status
		if err != nil {
			return nil, err2
		}
		fmt.Println(job.status)
	}

	if job.status == STATUS_SUCCESS {
		return job.data, nil
	} else {
		return nil, errors.New("error occurred Status: " + job.status)
	}
}

// Gets result of the Job
func (job *JobStruct) GetResult() (map[string]interface{}, error) {
	if job.status == STATUS_SUCCESS {
		return job.data, nil
	}
	return nil, errors.New("No Result Available, Status " + job.status)
}

// Sends an HTTP Request to the ScrapeIt.net API with proper Authentication
// Paramters:
// - method http method to use in request
// - url   path to resource to send request to, note that hostname/api will be prepended for you
// - body  parameters to send via the request body
// Returns
// - map[string]interface{} containing scraped data
// - error if non 200 http status is returned
func sendRequest(method string, url string, body map[string]string) (map[string]interface{}, error) {
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

	var result_interface interface{}
	json.Unmarshal(bytes, &result_interface)
	return result_interface.(map[string]interface{}), nil
}

// Checks if a Job Status is a final state
// if this method returns true it means ScrapeIt.net is no longer running the job
func isFinalState(status string) bool {
	switch status {
	case STATUS_SUCCESS:
		return true
	case STATUS_ERROR:
		return true
	case STATUS_FAILURE:
		return true
	default:
		return false
	}
}
