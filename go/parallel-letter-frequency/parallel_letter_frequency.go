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

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	channel := make(chan FreqMap, len(l))
	for _, words := range l {
		go func(w string) { channel <- Frequency(w) }(words)
	}

	// Reduce the results to a single map
	frequency := FreqMap{}
	for range l {
		for key, value := range <-channel {
			frequency[key] += value
		}
	}
	return frequency
}
