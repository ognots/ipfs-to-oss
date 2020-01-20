package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"git.jd.com/jcloud-api-gateway/jcloud-sdk-go/core"
	"git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/epns/apis"
	"git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/epns/client"
	"git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/epns/models"
)

func main() {
	const defaultRegion = "cn-south-1"
	const helpText = `
ipfs to jdcloud edge storage migration client

environment variables:
  JD_ACCESS_KEY_ID
  	access key id
  JD_SECRET_ACCESS_KEY
  	secrety access key
  JD_BUCKET
  	oss bucket
  JD_REGION
  	oss region
  JD_IPFS_REGION
  	ipfs cluster region

flags:`
	ak := os.Getenv("JD_ACCESS_KEY_ID")
	sk := os.Getenv("JD_SECRET_ACCESS_KEY")
	bucketName := os.Getenv("JD_BUCKET")
	ossRegion := os.Getenv("JD_REGION")
	ipfsRegion := os.Getenv("JD_IPFS_REGION")
	help := flag.Bool("help", false, "print help information")
	cid := flag.String("cid", "", " ipfs cid to migrate to oss bucket - REQUIRED")
	ossKey := flag.String("output", "", "key name in oss bucket cid. if unset uses cid")
	flag.Parse()

	if *help {
		fmt.Println(helpText)
		flag.PrintDefaults()
		return
	}
	if ak == "" {
		log.Fatal("JD_ACCESS_KEY_ID environment is required")
	}
	if sk == "" {
		log.Fatal("JD_SECRET_ACCESS_KEY environment is required")
	}
	if bucketName == "" {
		log.Fatal("JD_BUCKET environment is required")
	}
	if ossRegion == "" {
		ossRegion = defaultRegion
	}
	if ipfsRegion == "" {
		ipfsRegion = ossRegion
	}
	if *cid == "" {
		log.Fatal("cid flag is required")
	}
	if *ossKey == "" {
		ossKey = cid
	}

	log.Printf("migrating ipfs cid '%s' as '%s' to bucket '%s'", *cid, *ossKey, bucketName)
	log.Printf("ipfs cluster region: %s, oss region: %s", ipfsRegion, ossRegion)

	epnsClient := client.NewEpnsClient(core.NewCredential(ak, sk))

	ipfs2OssTaskID, err := AddFileToOSS(epnsClient, ipfsRegion, *cid, *ossKey, ossRegion, bucketName)
	if err != nil {
		log.Fatal(err)
		return
	}
	if checkTaskSuccess(epnsClient, ipfs2OssTaskID, ipfsRegion) {
		log.Printf("task with id(%s) is successful \n", ipfs2OssTaskID)
	}
}

func AddFileToOSS(epnsClient *client.EpnsClient, ipfsRegion, cid, ossKey, region, bucketName string) (string, error) {
	req := apis.NewCreateMigrationTaskRequestWithoutParam()

	req.SetRegionId(ipfsRegion)
	req.SetOssRegion(ipfsRegion)
	req.SetOssBucket(bucketName)
	req.SetGroup(bucketName)

	req.Items = []models.MigrationItem{
		newMigrationItem(ossKey, cid),
	}
	req.AddHeader("migrationType", "2")

	resp, _ := epnsClient.CreateMigrationTask(req)
	if resp.Error.Code != 0 {
		log.Print(resp.Error.Details)
		return "", errors.New(resp.Error.Message)
	}
	if len(resp.Result.TaskIds) > 0 {
		return resp.Result.TaskIds[0], nil
	}
	return "", errors.New("no tasks found")
}

func checkTaskSuccess(epnsClient *client.EpnsClient, taskID, region string) bool {
	timeout := time.NewTimer(30 * time.Second)
	for {
		select {
		case <-timeout.C:
			return false
		default:
			taskInfo := getTaskInfo(epnsClient, taskID, region)
			if taskInfo.State == 3 { // 3 is success flag
				timeout.Stop()
				return true
			}
			time.Sleep(5 * time.Second)
		}
	}
}

func getTaskInfo(epnsClient *client.EpnsClient, taskId, region string) models.TaskInfo {
	req := apis.NewGetTasksRequest(region)
	req.SetFilter(taskId)

	req.AddHeader("migrationType", "1")
	resp, _ := epnsClient.GetTasks(req)

	return resp.Result.TaskInfos[0]
}

func newMigrationItem(key, cid string) models.MigrationItem {
	return models.MigrationItem{
		OssKey:     key,
		ObjectKey:  &key,
		ObjectHash: &cid,
	}
}
