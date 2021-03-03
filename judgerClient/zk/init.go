package zk

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"math/rand"
	"time"
)

func ConnectZK() (*zk.Conn, error) {
	// TODO 配置启动
	address := fmt.Sprintf("%s:%s", "127.0.0.1", "2181")
	var hosts = []string{address}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		log.Printf("Connect zk failed: %v", err)
		return nil, err
	}
	return conn, nil
}

func GetOneNode(conn *zk.Conn) string {
	nodeList, _ := getNode(conn)
	return loadBalance(nodeList)
}

func getNode(conn *zk.Conn) ([]string, error) {
	root := "/Judger"
	var nodeList []string
	children, _, err := conn.Children(root)
	if err != nil {
		log.Printf("get zk node failed: %v\n", err)
		return nil, err
	}
	for _, child := range children {
		log.Printf("get node: %v\n", child)
		childPath := root + "/" + child
		content, _, err := conn.Get(childPath)
		if err != nil {
			log.Printf("get child node failed: %v\n", err)
			continue
		}
		nodeList = append(nodeList, string(content))
	}
	log.Printf("zk node list:%v\n", nodeList)
	return nodeList, nil
}

func loadBalance(nodeList []string) string {
	llen := len(nodeList)
	index := rand.Intn(llen)
	return nodeList[index]
}
