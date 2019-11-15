package scdn

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

// DeleteScdnDomain invokes the scdn.DeleteScdnDomain API synchronously
// api document: https://help.aliyun.com/api/scdn/deletescdndomain.html
func (client *Client) DeleteScdnDomain(request *DeleteScdnDomainRequest) (response *DeleteScdnDomainResponse, err error) {
	response = CreateDeleteScdnDomainResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteScdnDomainWithChan invokes the scdn.DeleteScdnDomain API asynchronously
// api document: https://help.aliyun.com/api/scdn/deletescdndomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteScdnDomainWithChan(request *DeleteScdnDomainRequest) (<-chan *DeleteScdnDomainResponse, <-chan error) {
	responseChan := make(chan *DeleteScdnDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteScdnDomain(request)
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

// DeleteScdnDomainWithCallback invokes the scdn.DeleteScdnDomain API asynchronously
// api document: https://help.aliyun.com/api/scdn/deletescdndomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteScdnDomainWithCallback(request *DeleteScdnDomainRequest, callback func(response *DeleteScdnDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteScdnDomainResponse
		var err error
		defer close(result)
		response, err = client.DeleteScdnDomain(request)
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

// DeleteScdnDomainRequest is the request struct for api DeleteScdnDomain
type DeleteScdnDomainRequest struct {
	*requests.RpcRequest
	OwnerAccount    string           `position:"Query" name:"OwnerAccount"`
	DomainName      string           `position:"Query" name:"DomainName"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SecurityToken   string           `position:"Query" name:"SecurityToken"`
}

// DeleteScdnDomainResponse is the response struct for api DeleteScdnDomain
type DeleteScdnDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteScdnDomainRequest creates a request to invoke DeleteScdnDomain API
func CreateDeleteScdnDomainRequest() (request *DeleteScdnDomainRequest) {
	request = &DeleteScdnDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DeleteScdnDomain", "", "")
	return
}

// CreateDeleteScdnDomainResponse creates a response to parse from DeleteScdnDomain response
func CreateDeleteScdnDomainResponse() (response *DeleteScdnDomainResponse) {
	response = &DeleteScdnDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
