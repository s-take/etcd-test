package main

import (
    "github.com/coreos/go-etcd/etcd"
    "log"
    "fmt"
)

var etcdServers = []string{
    "http://127.0.0.1:2379",
}

func main() {
    client := etcd.NewClient(etcdServers)
    resp, err := client.Get("/state", false, true)
    if err != nil {
        log.Fatal(err)
    }
    for _, n := range resp.Node.Nodes {
        fmt.Printf("%s: %s\n", n.Key, n.Value)
    }
}