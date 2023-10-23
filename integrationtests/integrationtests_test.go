package integrationtests

import (
  "slices"
  "testing"

  "poc-replication/scenarios"
  "poc-replication/node"
  "mock-network-golang/network"
)

func TestNodeInitialState(t *testing.T) {
  n := node.New("url", network.New())
  if n.State != "follower" {
    t.Errorf("inital state should be follower, not %s", n.State)
  }
}

func TestScenariosSetUpNodes(t *testing.T) {
  urls := []string{"url-1", "url-2", "url-3"}
  nodes := scenarios.SetUpNodes(urls)
  for _, node := range nodes {
    if len(node.OtherNodesUrls) != (len(urls) - 1) {
      t.Errorf("the length from OtherNodesUrls should be %q, and got %q", len(urls) - 1, len(node.OtherNodesUrls))
    }

    for _, url := range urls {
      if url != node.GetUrl() {
        if !slices.Contains(node.OtherNodesUrls, url) {
          t.Errorf("url %s is missing in node.OtherNodesUrls %v", url, node.OtherNodesUrls)
        } 
      }
    }
  }
}

func TestRunProcess(t *testing.T) {
  n := node.New("url", network.New())
  if err := n.RunProcess("heartbeatListener"); err != nil {
    t.Error("heartbeatListener can only be run in follower state")
  }
  
  n.State = "leader"
  if err := n.RunProcess("heartbeatPublisher"); err != nil {
    t.Error("heartbeatPublisher can only be run in leader state")
  }

  n.State = "candidate"
  if err := n.RunProcess("election"); err != nil {
    t.Error("heartbeatPublisher can only be run in candidate state")
  }
}
