/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 14:53 2019-09-27
 */
package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"

	"log"
	"time"
)

func main() {
	var (
		config clientv3.Config  // etcd config
		client *clientv3.Client // etcd client
		err error
	)
	config = clientv3.Config{
		Endpoints:[]string{"127.0.0.1:2379"}, // 集群地址
		DialTimeout: 5 * time.Millisecond,
	}

	log.Println(1)

	if client,err = clientv3.New(config);err != nil {
		fmt.Println(err)
		return
	}

	log.Println(2)

	// 创建读取etcd的键值对
	kv := clientv3.NewKV(client)

	response, err := kv.Put(context.TODO(), "/name", "one")
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(response.Header.Revision)

	getResponse, err := kv.Get(context.TODO(), "/name")
	if err != nil {
		panic(err)
	}

	log.Println(getResponse.Kvs)
}