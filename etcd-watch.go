package main

import (
    "os"
    "os/signal"
    "github.com/coreos/go-etcd/etcd"
    "log"
)

var etcdServers = []string{
    "http://127.0.0.1:2379",
}

func main() {
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    etcdResponseChan := make(chan *etcd.Response)
    client := etcd.NewClient(etcdServers)
    go func(){
        _, err := client.Watch("/state", 0, true, etcdResponseChan, nil)
        if err != nil {
            log.Fatal(err)
        }
    }()
    log.Println("Waiting for an update...")

    go func(){
        for {
            response := <-etcdResponseChan
            log.Printf("Got updated state : path=[%s], value=[%s]\n",
                        response.Node.Key, response.Node.Value)
        }
    }()
    <-c
}