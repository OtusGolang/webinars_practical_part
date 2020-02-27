package main

func primeFinder(done <-chan interface{}, intStream <-chan int) <-chan interface{} {
	primeStream := make(chan interface{})
	go func() {
		defer close(primeStream)
		for v := range intStream {
			select {
			case <-done:
				return
			default:
				if isPrime(v) {
					select {
					case <-done:
						return
					case primeStream <- v:
					}
				}
			}
		}
	}()
	return primeStream
}

// very slow calculation
func isPrime(v int) bool {
	for i := 2; i < v-1; i++ {
		if v%i == 0 {
			return false
		}
	}
	return true
}
