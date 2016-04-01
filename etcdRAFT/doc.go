//	Usage instruction :
// 		./etcdRAFT --idx 1 --workerport ":12345"
//		./etcdRAFT --idx 2 --workerport ":12346" --clusterIP ":12345"
//		./etcdRAFT --idx 3 --workerport ":12347" --clusterIP ":12345"
//
//		Each process will propose a different key value pair to be stored.
//		The cluster runs until the three of them reach concensus over
//		all the three values.
//
//		Can be extended to any number of nodes. Requies small change in code
//
package main
