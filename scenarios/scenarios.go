package scenarios

import (
  "poc-replication/node"
)

func SetUpNodes(urls []string) []*node.Node {
  nodes := []*node.Node{}
  for _, url := range urls {
    n := node.New(url)

    for _, otherUrl := range urls {
      if otherUrl != n.Url {
        n.RegisterNode(otherUrl)
      }
    }

    nodes = append(nodes, n)
  }

  return nodes
}
