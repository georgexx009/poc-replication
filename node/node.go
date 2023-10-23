package node

import (
	"fmt"
	"time"

  "mock-network-golang/basenode"
  "mock-network-golang/network"
)

func New(url string, n *network.Network) *Node {
  return &Node{
    State: "follower",
    HeartbeatTimeout: time.Second * 5,
    baseNode: *basenode.New(url, n),
  }
}

type Node struct {
  State string
  OtherNodesUrls []string
  HeartbeatTimeout time.Duration
  baseNode basenode.Basenode
  timeoutChannel chan bool
}

func (n *Node) GetUrl() string {
  return n.baseNode.HostUrl
}

func (n *Node) RegisterNode(url string) {
  n.OtherNodesUrls = append(n.OtherNodesUrls, url) 
}

func (n *Node) RunProcess(name string) error {
  switch name {
  case "heartbeatListener":
    if (n.State == "follower") {
      n.heartbeatListener()  
      return nil
    }
    return fmt.Errorf("%s cannot be run in state %s", name, n.State)
  case "heartbeatPublisher":
    if (n.State == "leader") {
      n.heartbeatPublisher()  
      return nil
    }
    return fmt.Errorf("%s cannot be run in state %s", name, n.State)
  case "election":
    if (n.State == "candidate") {
      n.election()  
      return nil
    }
    return fmt.Errorf("%s cannot be run in state %s", name, n.State)
  }

  return fmt.Errorf("unsupported process %s", name)
}

func (n *Node) heartbeatListener() {
  
}

func (n *Node) heartbeatPublisher() {

}

func (n *Node) election() {

}
