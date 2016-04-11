//	Usage instruction :
// 		./etcdRAFT --idx 1 --workerport ":12345"
//		./etcdRAFT --idx 2 --workerport ":12346" --clusterIP ":12345"
//		./etcdRAFT --idx 3 --workerport ":12347" --clusterIP ":12345"
//		./etcdRAFT --idx 4 --workerport ":12348" --clusterIP ":12345"
//		./etcdRAFT --idx 5 --workerport ":12349" --clusterIP ":12345"
//
//		Each process will propose a different key value pair to be stored.
//		The cluster reaches concensus over the proposed values
//
//		Can be extended to any number of nodes.
//
package main
