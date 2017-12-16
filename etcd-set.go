package main

import (
    "github.com/coreos/go-etcd/etcd"
    "log"
)

var etcdServers = []string{
    "http://127.0.0.1:2379",
}

func SetStatus_for_server(client *etcd.Client, status string){
    // Set "active" to path of "/state/server"
    _, err := client.Set("/state/server", status, 0)
    if err != nil {
        log.Fatal(err)
    }
}

func SetStatus_for_network(client *etcd.Client, status string){
    // Set "active" to path of "/state/network"
    _, err := client.Set("/state/network", status, 0)
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    // Create etcd client
    client := etcd.NewClient(etcdServers)
    SetStatus_for_server(client, "active")
    SetStatus_for_network(client, "down")
}