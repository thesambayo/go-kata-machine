package helpers

import (
	"errors"
	"path"
	"path/filepath"
	"runtime"
)

// func GetDSA() []string {
// 	return []string{
// 		"LinearSearchList",
// 		"BinarySearchList",
// 		"TwoCrystalBalls",
// 		"BubbleSort",
// 		"Stack",
// 		"Queue",
// 		//"SinglyLinkedList",
// 		//"DoublyLinkedList",
// 		//"ArrayList",
// 		//"MazeSolver",
// 		//"QuickSort",
// 		//"DFSOnBST",
// 		//"LRU",
// 		//"BTPreOrder",
// 		//"BTInOrder",
// 		//"BTPostOrder",
// 		//"BTBFS",
// 		//"CompareBinaryTrees",
// 		//"DFSOnBST",
// 		//"DFSGraphList",
// 		//"Trie",
// 		//"BFSGraphMatrix",
// 		//"Map",
// 		//"MinHeap",
// 	}
// }

func GetRootDir() (string, error) {
	_, runtimeFile, _, ok := runtime.Caller(0)
	// fmt.Println(runtimeFile)
	if !ok {
		return "", errors.New("unable to find root directory")
	}
	runtimeDir := path.Join(path.Dir(runtimeFile))
	return filepath.Dir(runtimeDir), nil
}
