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

package apis

import (
    "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/core"
    epns "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/epns/models"
)

type CreateMigrationTaskRequest struct {

    core.JDCloudRequest

    /* Region ID，例如：cn-north-1  */
    RegionId string `json:"regionId"`

    /* ipfs系统中的文件组，一般对应为OSS的bucket  */
    Group string `json:"group"`

    /* 迁移文件所在OSS的region，例如：cn-north-1  */
    OssRegion string `json:"ossRegion"`

    /* 文件所在的bucket  */
    OssBucket string `json:"ossBucket"`

    /*  (Optional) */
    Items []epns.MigrationItem `json:"items"`
}

/*
 * param regionId: Region ID，例如：cn-north-1 (Required)
 * param group: ipfs系统中的文件组，一般对应为OSS的bucket (Required)
 * param ossRegion: 迁移文件所在OSS的region，例如：cn-north-1 (Required)
 * param ossBucket: 文件所在的bucket (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateMigrationTaskRequest(
    regionId string,
    group string,
    ossRegion string,
    ossBucket string,
) *CreateMigrationTaskRequest {

	return &CreateMigrationTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/tasks",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Group: group,
        OssRegion: ossRegion,
        OssBucket: ossBucket,
	}
}

/*
 * param regionId: Region ID，例如：cn-north-1 (Required)
 * param group: ipfs系统中的文件组，一般对应为OSS的bucket (Required)
 * param ossRegion: 迁移文件所在OSS的region，例如：cn-north-1 (Required)
 * param ossBucket: 文件所在的bucket (Required)
 * param items:  (Optional)
 */
func NewCreateMigrationTaskRequestWithAllParams(
    regionId string,
    group string,
    ossRegion string,
    ossBucket string,
    items []epns.MigrationItem,
) *CreateMigrationTaskRequest {

    return &CreateMigrationTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/tasks",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Group: group,
        OssRegion: ossRegion,
        OssBucket: ossBucket,
        Items: items,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateMigrationTaskRequestWithoutParam() *CreateMigrationTaskRequest {

    return &CreateMigrationTaskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/tasks",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID，例如：cn-north-1(Required) */
func (r *CreateMigrationTaskRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param group: ipfs系统中的文件组，一般对应为OSS的bucket(Required) */
func (r *CreateMigrationTaskRequest) SetGroup(group string) {
    r.Group = group
}

/* param ossRegion: 迁移文件所在OSS的region，例如：cn-north-1(Required) */
func (r *CreateMigrationTaskRequest) SetOssRegion(ossRegion string) {
    r.OssRegion = ossRegion
}

/* param ossBucket: 文件所在的bucket(Required) */
func (r *CreateMigrationTaskRequest) SetOssBucket(ossBucket string) {
    r.OssBucket = ossBucket
}

/* param items: (Optional) */
func (r *CreateMigrationTaskRequest) SetItems(items []epns.MigrationItem) {
    r.Items = items
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateMigrationTaskRequest) GetRegionId() string {
    return r.RegionId
}

type CreateMigrationTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateMigrationTaskResult `json:"result"`
}

type CreateMigrationTaskResult struct {
    TaskIds []string `json:"taskIds"`
}