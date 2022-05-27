package main

func Hash(input string) []byte {
	ba := []byte(input)

	before := uint8(len(ba))

	for b := range ba {
		ba[b] = (ba[b] % (before + 1)) * before
		if ba[b] <= 0 {
			ba[b] = before
		}
		before = ba[b]
	}

	return ba
}
