package excelsheetcolumntitle

func ConvertToTitle(columnNumber int) string {
	result := []byte{}

	for columnNumber > 0 {
		columnNumber--

		remainder := columnNumber % 26

		ch := byte('A' + remainder)
		result = append(result, ch)
		columnNumber /= 26
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
