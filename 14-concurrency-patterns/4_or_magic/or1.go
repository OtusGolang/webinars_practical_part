package main

func or1(channels ...<-chan struct{}) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			for _, ch := range channels {
				select {
				case <-ch:
					return
				default:
				}
			}
		}
	}()

	return done
}
