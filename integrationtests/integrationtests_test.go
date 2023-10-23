package integrationtests

import (
  "slices"
  "testing"
  "poc-replication/scenarios"
  "poc-replication/node"
)

func TestNodeInitialState(t *testing.T) {
  n := node.New("url")
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
      if url != node.Url {
        if !slices.Contains(node.OtherNodesUrls, url) {
          t.Errorf("url %s is missing in node.OtherNodesUrls %v", url, node.OtherNodesUrls)
        } 
      }
    }
  }
}
