package node

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
