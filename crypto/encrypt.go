package crypto

const clau = "FRIENDSOFGO"

func Decrypt(input string) string {
	s := ""
	k := 0
	for i := 0; i < len(input); i++ {
		if input[i] < 'A' || input[i] > 'Z' {
			s += string(input[i])
			continue
		}

		s += string('A' + ((input[i] - clau[k%len(clau)] + 26) % 26))
		k++
	}

	return s
}

func Encrypt(input string) string {
	s := ""
	k := 0
	for i := 0; i < len(input); i++ {
		if input[i] < 65 || input[i] > 90 {
			s += string(input[i])
			continue
		}

		s += string(65 + (input[i] + clau[k%len(clau)] + 26) % 26)
		k++
	}

	return s
}

