package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateTempDBInstance invokes the rds.CreateTempDBInstance API synchronously
// api document: https://help.aliyun.com/api/rds/createtempdbinstance.html
func (client *Client) CreateTempDBInstance(request *CreateTempDBInstanceRequest) (response *CreateTempDBInstanceResponse, err error) {
	response = CreateCreateTempDBInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTempDBInstanceWithChan invokes the rds.CreateTempDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/createtempdbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTempDBInstanceWithChan(request *CreateTempDBInstanceRequest) (<-chan *CreateTempDBInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateTempDBInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTempDBInstance(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateTempDBInstanceWithCallback invokes the rds.CreateTempDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/createtempdbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTempDBInstanceWithCallback(request *CreateTempDBInstanceRequest, callback func(response *CreateTempDBInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTempDBInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateTempDBInstance(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreateTempDBInstanceRequest is the request struct for api CreateTempDBInstance
type CreateTempDBInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RestoreTime          string           `position:"Query" name:"RestoreTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	BackupId             requests.Integer `position:"Query" name:"BackupId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CreateTempDBInstanceResponse is the response struct for api CreateTempDBInstance
type CreateTempDBInstanceResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	TempDBInstanceId string `json:"TempDBInstanceId" xml:"TempDBInstanceId"`
}

// CreateCreateTempDBInstanceRequest creates a request to invoke CreateTempDBInstance API
func CreateCreateTempDBInstanceRequest() (request *CreateTempDBInstanceRequest) {
	request = &CreateTempDBInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CreateTempDBInstance", "Rds", "openAPI")
	return
}

// CreateCreateTempDBInstanceResponse creates a response to parse from CreateTempDBInstance response
func CreateCreateTempDBInstanceResponse() (response *CreateTempDBInstanceResponse) {
	response = &CreateTempDBInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
