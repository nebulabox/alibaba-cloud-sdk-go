package green

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

// DescribeImageUploadInfo invokes the green.DescribeImageUploadInfo API synchronously
// api document: https://help.aliyun.com/api/green/describeimageuploadinfo.html
func (client *Client) DescribeImageUploadInfo(request *DescribeImageUploadInfoRequest) (response *DescribeImageUploadInfoResponse, err error) {
	response = CreateDescribeImageUploadInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeImageUploadInfoWithChan invokes the green.DescribeImageUploadInfo API asynchronously
// api document: https://help.aliyun.com/api/green/describeimageuploadinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeImageUploadInfoWithChan(request *DescribeImageUploadInfoRequest) (<-chan *DescribeImageUploadInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeImageUploadInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeImageUploadInfo(request)
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

// DescribeImageUploadInfoWithCallback invokes the green.DescribeImageUploadInfo API asynchronously
// api document: https://help.aliyun.com/api/green/describeimageuploadinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeImageUploadInfoWithCallback(request *DescribeImageUploadInfoRequest, callback func(response *DescribeImageUploadInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeImageUploadInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeImageUploadInfo(request)
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

// DescribeImageUploadInfoRequest is the request struct for api DescribeImageUploadInfo
type DescribeImageUploadInfoRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
}

// DescribeImageUploadInfoResponse is the response struct for api DescribeImageUploadInfo
type DescribeImageUploadInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Accessid  string `json:"Accessid" xml:"Accessid"`
	Policy    string `json:"Policy" xml:"Policy"`
	Signature string `json:"Signature" xml:"Signature"`
	Folder    string `json:"Folder" xml:"Folder"`
	Host      string `json:"Host" xml:"Host"`
	Expire    int    `json:"Expire" xml:"Expire"`
}

// CreateDescribeImageUploadInfoRequest creates a request to invoke DescribeImageUploadInfo API
func CreateDescribeImageUploadInfoRequest() (request *DescribeImageUploadInfoRequest) {
	request = &DescribeImageUploadInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "DescribeImageUploadInfo", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeImageUploadInfoResponse creates a response to parse from DescribeImageUploadInfo response
func CreateDescribeImageUploadInfoResponse() (response *DescribeImageUploadInfoResponse) {
	response = &DescribeImageUploadInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
