package count_servers_that_communicate

import "testing"

func TestServer(t *testing.T) {
	t.Log(countServers([][]int{{1, 0}, {1, 1}}))
}
