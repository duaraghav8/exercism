package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune for the given list of
// strings. Each string is analyzed concurrently.
func ConcurrentFrequency(strings []string) FreqMap {
	freqCh := make(chan FreqMap)

	for _, input := range strings {
		// Submit input for analysis
		go func(input string, outputCh chan<- FreqMap) {
			outputCh <- Frequency(input)
		}(input, freqCh)
	}

	// Collect the response from the first goroutine to finish
	resp := <-freqCh

	// Collect responses from all subsequent goroutines and merge them
	// into resp.
	for i := 1; i < len(strings); i++ {
		freq := <-freqCh

		for r, f := range freq {
			// If rune r already has an entry in resp, just
			// add to its value.
			// Otherwise, create an entry for r with value f
			if _, ok := resp[r]; ok {
				resp[r] += f
			} else {
				resp[r] = f
			}
		}
	}

	return resp
}