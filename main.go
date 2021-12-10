package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"log"
	"time"
)

var endpoints = []string{
	"172.18.30.61:12379",
	"172.18.30.61:22379",
	"172.18.30.61:32379",
}

func main()  {
	cli,err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
		DialTimeout: 5 * time.Second,
	})

	check(err)
	defer cli.Close()

	//设置每次操作的超时时间
	//ctx,cancel := context.WithTimeout(context.Background(),5 * time.Second)
	//defer cancel()

	ctx := context.Background()

	//put key
	res,err := cli.Put(ctx,"jijin","aways up!",clientv3.WithPrevKV())
	check(err)

	fmt.Println(res)


	//watch key
	//返回一个接收通道，当被watch的key有修改时，通道都会接收到事件通知
	wc := cli.Watch(ctx,"jijin",clientv3.WithPrevKV())

	for re := range wc {
		for _,ev := range re.Events {
			fmt.Printf("type:%s,kv:%s,preKv:%s\n",ev.Type,ev.Kv,ev.PrevKv)
		}
	}

	//lease
	lease,err := cli.Grant(ctx,30)
	check(err)
	jsonByte,_ := json.Marshal(lease)
	fmt.Println(string(jsonByte))

	//绑定key
	cli.Put(ctx,"jijin","up",clientv3.WithLease(lease.ID))

	//查看lease信息，包括ttl和绑定的key集合
	leaseLive,err := cli.TimeToLive(ctx,lease.ID,clientv3.WithAttachedKeys())
	check(err)

	for _,key := range leaseLive.Keys {
		fmt.Println(string(key))
	}

	jsonByte,_ = json.Marshal(leaseLive)
	fmt.Println(string(jsonByte))
	//keys的类型是[][]byte，json编码时会对byte[]类型进行base64_encode，所以可以看到下面结果中的key并不是明文
	//{"cluster_id":17237436991929493444,"member_id":18249187646912138824,"revision":8903,"raft_term":9,"id":4703584957079197192,"ttl":29,"granted-ttl":30,"keys":["amlqaW4="]}

	//获取现存的所有lease，只返回lease id集合
	leases,err := cli.Leases(ctx)
	jsonByte,_ = json.Marshal(leases)
	fmt.Println(string(jsonByte))

}

func check(err error)  {
	if err != nil {
		log.Fatalln(err)
	}
}
