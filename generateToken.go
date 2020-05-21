package main
import (
    "fmt"
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk"
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)
func generateToken() {
    client, err := sdk.NewClientWithAccessKey("cn-shanghai", "<您的AccessKey Id>", "<您的AccessKey Secret>")
    if err != nil {
        panic(err)
    }
    request := requests.NewCommonRequest()
    request.Method = "POST"
    request.Domain = "nls-meta.cn-shanghai.aliyuncs.com"
    request.ApiName = "CreateToken"
    request.Version = "2019-02-28"
    response, err := client.ProcessCommonRequest(request)
    if err != nil {
        panic(err)
    }
    fmt.Print(response.GetHttpStatus())
    fmt.Print(response.GetHttpContentString())
}