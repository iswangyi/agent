package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

var (
	config  clientv3.Config
	client  *clientv3.Client
	err     error
	kv      clientv3.KV
	putResp *clientv3.PutResponse
)

var taskID = "0001"

func RunTask() string {
	mMap := make(map[string]interface{})
	mMap["cmd"] = "./tmp.sh"
	mMap["taskID"] = taskID
	mMap["status"] = "running"
	mMap["executor"] = "bash"
	mMap["express"] = "* * * * *"

	v, err := json.Marshal(mMap)
	fmt.Println(string(v), err)
	return string(v)
}
func main() {
	config = clientv3.Config{
		Endpoints:   []string{"0.0.0.0:2379"},
		DialTimeout: 5 * time.Second,
	}
	if client, err = clientv3.New(config); err != nil {
		return
	}

	kv = clientv3.NewKV(client)
	jsonStr := RunTask()
	//TODO() 占位
	putResp, _ := kv.Put(
		context.TODO(),
		"/cron/jobs/127.0.0.1/"+taskID,
		jsonStr,
		clientv3.WithPrevKV())
	fmt.Println(putResp)

}
