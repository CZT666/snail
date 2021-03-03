package zk

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"snail/judger/settings"
	"time"
)

func connect(config *settings.ZKConfig) *zk.Conn {
	address := fmt.Sprintf("%s:%s", config.Host, config.Port)
	var hosts = []string{address}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		log.Printf("Connect zk failed: %v", err)
		return nil
	}
	createNode("root", config.Root, conn, 0)
	return conn
}

func register(config *settings.ZKConfig, conn *zk.Conn, host string, port string) {
	msg := fmt.Sprintf("%s:%s", host, port)
	createNode(msg, config.Node, conn, 3)
}

func createNode(msg string, path string, conn *zk.Conn, flag int32) {
	var data = []byte(msg)
	var acls = zk.WorldACL(zk.PermAll) //控制访问权限模式
	p, err := conn.Create(path, data, flag, acls)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("create:", p)
}

func InitZK(config *settings.ZKConfig, host string, port string) *zk.Conn {
	conn := connect(config)
	register(config, conn, host, port)
	return conn
}

func Close(conn *zk.Conn) {
	conn.Close()
}
