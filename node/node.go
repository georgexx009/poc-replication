package node

import "fmt"

func New(url string) *Node {
  return &Node{
    Url: url,
    State: "follower",
  }
}

type Node struct {
  Url string
  State string
  OtherNodesUrls []string
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
