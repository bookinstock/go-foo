package main

import "foo/concurrency"

func main() {
	// err := mod_b.PrintB()
	// if err != nil {
	// 	fmt.Printf("%+v\n", err)
	// }

	// concurrency.RunPubSub()

	concurrency.RunPriorityFanIn()
}
