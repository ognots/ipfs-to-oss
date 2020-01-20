// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package client

import (
    "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/core"
    epns "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/epns/apis"
    "encoding/json"
    "errors"
)

type EpnsClient struct {
    core.JDCloudClient
}

func NewEpnsClient(credential *core.Credential) *EpnsClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("epns.jdcloud-api.com")

    return &EpnsClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "epns",
            Revision:    "1.0.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *EpnsClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *EpnsClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *EpnsClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 获取Task列表详情
 */
func (c *EpnsClient) GetTasks(request *epns.GetTasksRequest) (*epns.GetTasksResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &epns.GetTasksResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* getMigrationFlag
 */
func (c *EpnsClient) GetMigrationFlag(request *epns.GetMigrationFlagRequest) (*epns.GetMigrationFlagResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &epns.GetMigrationFlagResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取task详情
 */
func (c *EpnsClient) GetTaskInfo(request *epns.GetTaskInfoRequest) (*epns.GetTaskInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &epns.GetTaskInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取用户下的文件列表
 */
func (c *EpnsClient) GetFiles(request *epns.GetFilesRequest) (*epns.GetFilesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &epns.GetFilesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* deleteMigrationFlag
 */
func (c *EpnsClient) DeleteMigrationFlag(request *epns.DeleteMigrationFlagRequest) (*epns.DeleteMigrationFlagResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &epns.DeleteMigrationFlagResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* createMigrationTask
 */
func (c *EpnsClient) CreateMigrationTask(request *epns.CreateMigrationTaskRequest) (*epns.CreateMigrationTaskResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &epns.CreateMigrationTaskResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* createMigrationFlag
 */
func (c *EpnsClient) CreateMigrationFlag(request *epns.CreateMigrationFlagRequest) (*epns.CreateMigrationFlagResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &epns.CreateMigrationFlagResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

