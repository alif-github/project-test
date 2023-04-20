package dao

import "fmt"

func ParameterQuery(rangeData int, startIndex *int) (result string) {
	for i := 0; i < rangeData; i++ {
		result += fmt.Sprintf(`$%d`, *startIndex)
		if i < rangeData-1 {
			result += ", "
		}

		*startIndex++
	}

	return
}
