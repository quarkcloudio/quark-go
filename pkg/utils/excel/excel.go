package excel

// 生成 Excel 中的列标签
func GenerateColumnLabel(index int) string {

	var columnLabel string

	for index > 0 {
		mod := (index - 1) % 26
		columnLabel = string(rune('A'+mod)) + columnLabel
		index = (index - 1) / 26
	}

	return columnLabel
}
