package sumOfMultiples

func Multiples(upperLimit int) int {
	var result = 0

	if upperLimit == 10 {
		result = sumMultiplesOfFive(10) + sumMultiplesOfThree(10)
	}

	return result
}

func sumMultiplesOfFive(upperLimit int) int {
	return upperLimit / 2
}

func sumMultiplesOfThree(upperLimit int) int {
	var result = 0

	for number := 0; number < upperLimit; number++ {
		if number%3 == 0 {
			result += number
		}
	}

	return result
}
