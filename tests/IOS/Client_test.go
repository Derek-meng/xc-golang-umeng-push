package IOS

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"xc-golang-umeng-push/src/Constants"
	"xc-golang-umeng-push/src/Service"
	"net/http"
	"testing"
)

var client *Service.IOSClient
var payload Service.Payload

func init() {
	client = Service.NewIOSClient("your app key", "your secret", Constants.TEST)
	alert := Service.AlertParams{
		"title", "subTitle", "Body",
	}
	aps := Service.ApsParams{Alert: alert}
	payload = Service.Payload{
		Aps: aps,
	}
}

func TestBroadcast(t *testing.T) {
	assertions := assert.New(t)
	push, err := client.Broadcast(payload)
	defer push.Close()
	if err != nil {
		t.Error(err)
	}
	if push.IsConnectSuccess() {
		if !push.IsErrorOccur() {
			assertions.NotEqual("", push.GetTaskId(), "task id should not be null ")
			assertions.Equal(http.StatusOK, push.GetHttpResponse().StatusCode, "http code should  be 200 ")
			assertions.Equal("", push.GetErrorCode(), "error code should  be null ")
			assertions.Equal("", push.GetErrorMessage(), "error message should  be null ")

		} else {
			assertions.NotEqual("", push.GetErrorCode(), "error code should not be null ")
			assertions.NotEqual("", push.GetErrorMessage(), "error message should not be null ")
			assertions.Equal("", push.GetTaskId(), "task id should  be null ")

		}
		fmt.Println(push.All())
	} else {
		fmt.Println(push.GetHttpResponse().StatusCode)
	}

}

func TestUniCast(t *testing.T) {
	assertions := assert.New(t)
	push, err := client.UniCast(payload, "456045aa01eb7f1fed5d5b62e3fd3aeded07f5a5bb36db3773ffa3eb2da7ba2f")
	defer push.Close()
	if err != nil {
		t.Error(err)
	}
	if push.IsConnectSuccess() {
		if !push.IsErrorOccur() {
			assertions.NotEqual("", push.GetMessageId(), "message id should not be null ")
			assertions.Equal(http.StatusOK, push.GetHttpResponse().StatusCode, "http code should  be 200 ")
			assertions.Equal("", push.GetErrorCode(), "error code should  be null ")
			assertions.Equal("", push.GetErrorMessage(), "error message should  be null ")

		} else {
			assertions.NotEqual("", push.GetErrorCode(), "error code should not be null ")
			assertions.NotEqual("", push.GetErrorMessage(), "error message should not be null ")
			assertions.Equal("", push.GetMessageId(), "task id should  be null ")

		}
		fmt.Println(push.All())
	} else {
		fmt.Println(push.GetHttpResponse().StatusCode)
	}
}

func TestListPush(t *testing.T) {
	assertions := assert.New(t)
	deviceToken := []string{
		"456045aa01eb7f1fed5d5b62e3fd3aeded07f5a5bb36db3773ffa3eb2da7ba2f",
		"9f0808e352c3a464e231928f5636650d168ab7bef1ccdaf9ad7aee244e0d0cce",
		"7149f4fe1a09dc29e2f36dfa981e76f2c16291946bfe701d859931cfe02d3e61",
	}

	push, err := client.ListPush(payload, deviceToken)
	defer push.Close()
	if err != nil {
		t.Error(err)
	}
	if push.IsConnectSuccess() {
		if !push.IsErrorOccur() {
			assertions.NotEqual("", push.GetMessageId(), "message id should not be null ")
			assertions.Equal(http.StatusOK, push.GetHttpResponse().StatusCode, "http code should  be 200 ")
			assertions.Equal("", push.GetErrorCode(), "error code should  be null ")
			assertions.Equal("", push.GetErrorMessage(), "error message should  be null ")

		} else {
			assertions.Equal(http.StatusBadRequest, push.GetHttpResponse().StatusCode, "http code should  be 400 ")
			assertions.NotEqual("", push.GetErrorCode(), "error code should not be null ")
			assertions.NotEqual("", push.GetErrorMessage(), "error message should not be null ")
			assertions.Equal("", push.GetMessageId(), "task id should  be null ")

		}
		fmt.Println(push.All())
	} else {
		fmt.Println(push.GetHttpResponse().StatusCode)
	}
}

func TestUpload(t *testing.T) {
	assertions := assert.New(t)
	deviceToken := []string{
		"456045aa01eb7f1fed5d5b62e3fd3aeded07f5a5bb36db3773ffa3eb2da7ba2f",
		"9f0808e352c3a464e231928f5636650d168ab7bef1ccdaf9ad7aee244e0d0cce",
		"7149f4fe1a09dc29e2f36dfa981e76f2c16291946bfe701d859931cfe02d3e61",
	}

	push, err := client.Upload(deviceToken)
	defer push.Close()
	if err != nil {
		t.Error(err)
	}
	if push.IsConnectSuccess() {
		if !push.IsErrorOccur() {
			assertions.NotEqual("", push.GetFileId(), "file id should not be null ")
			assertions.Equal(http.StatusOK, push.GetHttpResponse().StatusCode, "http code should  be 200 ")
			assertions.Equal("", push.GetErrorCode(), "error code should  be null ")
			assertions.Equal("", push.GetErrorMessage(), "error message should  be null ")

		} else {
			assertions.NotEqual("", push.GetErrorCode(), "error code should not be null ")
			assertions.NotEqual("", push.GetErrorMessage(), "error message should not be null ")
			assertions.Equal("", push.GetFileId(), "file id should  be null ")

		}
		fmt.Println(push.All())
	} else {
		fmt.Println(push.GetHttpResponse().StatusCode)
	}
}

func TestCustomizedPush(t *testing.T) {
	assertions := assert.New(t)
	fileIds := []string{
		"PF955921560487812770",
	}
	push, err := client.CustomizedPush(payload, "13246", "", fileIds)
	defer push.Close()
	if err != nil {
		t.Error(err)
	}
	if push.IsConnectSuccess() {
		if !push.IsErrorOccur() {
			assertions.NotEqual("", push.GetTaskId(), "task id should not be null ")

			assertions.Equal(http.StatusOK, push.GetHttpResponse().StatusCode, "http code should  be 200 ")
			assertions.Equal("", push.GetErrorCode(), "error code should  be null ")
			assertions.Equal("", push.GetErrorMessage(), "error message should  be null ")

		} else {
			assertions.NotEqual("", push.GetErrorCode(), "error code should not be null ")
			assertions.NotEqual("", push.GetErrorMessage(), "error message should not be null ")
			assertions.Equal("", push.GetTaskId(), "task id should  be null ")

		}
		fmt.Println(push.All())
	} else {
		fmt.Println(push.GetHttpResponse().StatusCode)
	}
}

func TestGroupPush(t *testing.T) {
	assertions := assert.New(t)
	push, err := client.GroupPush(payload, `{"where":"9f0808e352c3a464e231928f5636650d168ab7bef1ccdaf9ad7aee244e0d0cce"}`)
	defer push.Close()
	if err != nil {
		t.Error(err)
	}
	if push.IsConnectSuccess() {
		if !push.IsErrorOccur() {
			assertions.NotEqual("", push.GetTaskId(), "task id should not be null ")

			assertions.Equal(http.StatusOK, push.GetHttpResponse().StatusCode, "http code should  be 200 ")
			assertions.Equal("", push.GetErrorCode(), "error code should  be null ")
			assertions.Equal("", push.GetErrorMessage(), "error message should  be null ")

		} else {
			assertions.NotEqual("", push.GetErrorCode(), "error code should not be null ")
			assertions.NotEqual("", push.GetErrorMessage(), "error message should not be null ")
			assertions.Equal("", push.GetTaskId(), "task id should  be null ")

		}
		fmt.Println(push.All())
	} else {
		fmt.Println(push.GetHttpResponse().StatusCode)
	}
}

func TestStatus(t *testing.T) {

	assertions := assert.New(t)
	taskId := "ucghza0156048783213001"
	push, err := client.PushStatus(taskId)
	defer push.Close()
	if err != nil {
		fmt.Println(err)
	}
	if push.IsConnectSuccess() {
		if !push.IsErrorOccur() {
			assertions.Equal(http.StatusOK, push.GetHttpResponse().StatusCode, "http code should  be 200 ")
			assertions.Equal("", push.GetErrorCode(), "error code should  be null ")
			assertions.Equal("", push.GetErrorMessage(), "error message should  be null ")

		} else {
			assertions.NotEqual("", push.GetErrorCode(), "error code should not be null ")
			assertions.NotEqual("", push.GetErrorMessage(), "error message should not be null ")

		}
		var expectInt int

		assertions.IsType(expectInt, push.GetSentCount(), "should be int")
		assertions.IsType(expectInt, push.GetOpenCount(), "should be int")
		assertions.IsType(expectInt, push.GetMessageStatus(), "should be int")

		fmt.Println(push.All())
	} else {
		fmt.Println(push.GetHttpResponse().StatusCode)
	}

}

func TestFilePush(t *testing.T) {

	assertions := assert.New(t)
	fileIds := []string{
		"PF955921560487812770",
	}

	push, err := client.FilePush(payload, fileIds)
	defer push.Close()
	if err != nil {
		t.Error(err)
	}
	if push.IsConnectSuccess() {
		if !push.IsErrorOccur() {
			assertions.NotEqual("", push.GetTaskId(), "task id should not be null ")

			assertions.Equal(http.StatusOK, push.GetHttpResponse().StatusCode, "http code should  be 200 ")
			assertions.Equal("", push.GetErrorCode(), "error code should  be null ")
			assertions.Equal("", push.GetErrorMessage(), "error message should  be null ")

		} else {
			assertions.NotEqual("", push.GetErrorCode(), "error code should not be null ")
			assertions.NotEqual("", push.GetErrorMessage(), "error message should not be null ")
			assertions.Equal("", push.GetTaskId(), "task id should  be null ")

		}
		fmt.Println(push.All())
	} else {
		fmt.Println(push.GetHttpResponse().StatusCode)
	}

}

func TestChancelPush(t *testing.T) {

	assertions := assert.New(t)
	taskId := "PF955921560487812770"
	push, err := client.ChancelPush(taskId)
	defer push.Close()
	if err != nil {
		t.Error(err)
	}
	if push.IsConnectSuccess() {
		if !push.IsErrorOccur() {
			assertions.NotEqual("", push.GetTaskId(), "task id should not be null ")
			assertions.Equal(http.StatusOK, push.GetHttpResponse().StatusCode, "http code should  be 200 ")
			assertions.Equal("", push.GetErrorCode(), "error code should  be null ")
			assertions.Equal("", push.GetErrorMessage(), "error message should  be null ")

		} else {
			assertions.NotEqual("", push.GetErrorCode(), "error code should not be null ")
			assertions.NotEqual("", push.GetErrorMessage(), "error message should not be null ")
			assertions.Equal("", push.GetTaskId(), "task id should  be null ")

		}
		fmt.Println(push.All())
	} else {
		fmt.Println(push.GetHttpResponse().StatusCode)
	}

}
