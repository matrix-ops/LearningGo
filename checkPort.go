package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"
)

var (
	FailedReason error
)

func checkURLStatus(URL string) (int, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		FailedReason = err
		return 0, err
	}
	client := &http.Client{
		Timeout: time.Second * 3,
	}
	resp, err := client.Do(req)
	if err != nil {
		FailedReason = err
		return 0, err
	}
	err = nil
	validStatusCodes := []int{200, 301, 302}
	for _, validValue := range validStatusCodes {
		if resp.StatusCode != validValue {
			err = errors.New(resp.Status)
			FailedReason = err
		} else {
			break
		}
	}
	return resp.StatusCode, err
}
func startKubectlPortForward() {
	killKubectl := exec.Command("/usr/bin/killall", "kubectl")
	killKubectl.Run()
	cmd := exec.Command("/usr/bin/kubectl", "port-forward", "svc/xxl-job-admin", "--address", "0.0.0.0", "38888:8080")
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil
	if err := cmd.Start(); err != nil {
		FailedReason = err
		log.Fatalf("cmd.Start: %v", err)
	}
}
func main() {
	statusCode, err := checkURLStatus("PUT_YOURURL_THERE")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Application Failed , Reason :", FailedReason)
		startKubectlPortForward()
	} else {
		fmt.Println("Application  Running.", statusCode)
	}

}
