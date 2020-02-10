package cli

const (
	OUTPUT_DETAIL = "detail"
	OUTPUT_DOT    = "dot"
	OUTPUT_CSV    = "csv"
	OUTPUT_NULL   = "null"
)

func isValidOutputFormat(outputFormat string) bool {
	switch outputFormat {
	case OUTPUT_DETAIL, OUTPUT_DOT, OUTPUT_CSV, OUTPUT_NULL:
		return true
	}
	return false
}
