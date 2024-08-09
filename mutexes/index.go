package mutexes

import "fmt"

func Index() {
	fmt.Println("----- Mutexes Demo -----")

	MutexDemo()
	AtomicDemo()
	ReadWriteDemo()
}
