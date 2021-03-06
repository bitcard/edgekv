package main

// This example shows a ring of four dht nodes without storage
import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/ksonbol/edgekv/pkg/dht"
)

var (
	nodeAddr  = flag.String("node_addr", "localhost:5554", "The dht server address in the format of host:port")
	node2Addr = flag.String("node2_addr", "localhost:5555", "The helper node address in the format of host:port")
	node3Addr = flag.String("node3_addr", "localhost:5556", "The helper node address in the format of host:port")
	node4Addr = flag.String("node4_addr", "localhost:5557", "The helper node address in the format of host:port")
)

func main() {
	flag.Parse()
	fmt.Println("Starting the program")
	node := dht.NewLocalNode(*nodeAddr, "", nil, nil)
	node2 := dht.NewLocalNode(*node2Addr, "", nil, nil)
	node3 := dht.NewLocalNode(*node3Addr, "", nil, nil)
	node4 := dht.NewLocalNode(*node4Addr, "", nil, nil)
	fmt.Println("Created local nodes")
	m := map[int]*dht.Node{1: node, 2: node2, 3: node3, 4: node4}
	helperNode2 := dht.NewRemoteNode(node.Addr, node.ID, node2.Transport, nil)
	helperNode3 := dht.NewRemoteNode(node.Addr, node.ID, node3.Transport, nil)
	helperNode4 := dht.NewRemoteNode(node.Addr, node.ID, node4.Transport, nil)
	time.Sleep(2 * time.Second)
	node.Join(nil)
	node2.Join(helperNode2)
	node3.Join(helperNode3)
	node4.Join(helperNode4)
	fmt.Println("All nodes have joined the system")
	fmt.Println("Nodes IDs are: ")
	for k, v := range m {
		fmt.Println(k, v.ID)
	}
	fmt.Println()
	for {
		fmt.Println("Enter the node number (1-4) to print its info")
		var nodeNum int
		fmt.Scan(&nodeNum)
		n, ok := m[nodeNum]
		if !ok {
			log.Fatalf("Such node does not exist!")
		}
		fmt.Printf("Node %d\n with Addr %s, ID: %s\n", nodeNum, n.Addr, n.ID)
		// n.PrintFT()
		fmt.Printf("Successor: %s, Predecessor: %s\n", n.Successor().ID, n.Predecessor().ID)
		time.Sleep(5 * time.Second)
	}
}
