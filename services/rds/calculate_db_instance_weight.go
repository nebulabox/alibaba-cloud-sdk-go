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

// CalculateDBInstanceWeight invokes the rds.CalculateDBInstanceWeight API synchronously
// api document: https://help.aliyun.com/api/rds/calculatedbinstanceweight.html
func (client *Client) CalculateDBInstanceWeight(request *CalculateDBInstanceWeightRequest) (response *CalculateDBInstanceWeightResponse, err error) {
	response = CreateCalculateDBInstanceWeightResponse()
	err = client.DoAction(request, response)
	return
}

// CalculateDBInstanceWeightWithChan invokes the rds.CalculateDBInstanceWeight API asynchronously
// api document: https://help.aliyun.com/api/rds/calculatedbinstanceweight.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CalculateDBInstanceWeightWithChan(request *CalculateDBInstanceWeightRequest) (<-chan *CalculateDBInstanceWeightResponse, <-chan error) {
	responseChan := make(chan *CalculateDBInstanceWeightResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CalculateDBInstanceWeight(request)
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

// CalculateDBInstanceWeightWithCallback invokes the rds.CalculateDBInstanceWeight API asynchronously
// api document: https://help.aliyun.com/api/rds/calculatedbinstanceweight.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CalculateDBInstanceWeightWithCallback(request *CalculateDBInstanceWeightRequest, callback func(response *CalculateDBInstanceWeightResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CalculateDBInstanceWeightResponse
		var err error
		defer close(result)
		response, err = client.CalculateDBInstanceWeight(request)
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

// CalculateDBInstanceWeightRequest is the request struct for api CalculateDBInstanceWeight
type CalculateDBInstanceWeightRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CalculateDBInstanceWeightResponse is the response struct for api CalculateDBInstanceWeight
type CalculateDBInstanceWeightResponse struct {
	*responses.BaseResponse
	RequestId string                           `json:"RequestId" xml:"RequestId"`
	Items     ItemsInCalculateDBInstanceWeight `json:"Items" xml:"Items"`
}

// CreateCalculateDBInstanceWeightRequest creates a request to invoke CalculateDBInstanceWeight API
func CreateCalculateDBInstanceWeightRequest() (request *CalculateDBInstanceWeightRequest) {
	request = &CalculateDBInstanceWeightRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CalculateDBInstanceWeight", "Rds", "openAPI")
	return
}

// CreateCalculateDBInstanceWeightResponse creates a response to parse from CalculateDBInstanceWeight response
func CreateCalculateDBInstanceWeightResponse() (response *CalculateDBInstanceWeightResponse) {
	response = &CalculateDBInstanceWeightResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
