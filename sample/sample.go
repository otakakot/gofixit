package sample

type SampleInterface interface {
	// LongFunctionName is a function with a long name that takes multiple parameters.
	LongFunctionName(aaaaaaaaaa string, bbbbbbbbbb string, cccccccccc string, dddddddddd string, eeeeeeeeee string) (string, error)
	// LongFunctionNameNewLine is a function with a long name that takes multiple parameters.
	LongFunctionNameNewLine(
		aaaaaaaaaa string,
		bbbbbbbbbb string,
		cccccccccc string,
		dddddddddd string,
		eeeeeeeeee string,
	) (string, error)
	// ShortFunctionName is a function with a short name.
	Short(a string, b string) error
	// Empty is an empty function.
	Empty()
}
