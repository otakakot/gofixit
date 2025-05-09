package sample

import "github.com/otakakot/gofixit/sample/domain"

const cccccccccccccccccccc = "c"

// This file contains a sample of a function with a long name and multiple parameters.
type FunctionsInterfaceLLL interface {
	// LongFunctionName is a function with a long name that takes multiple parameters.
	LongFunctionName(aaaaaaaaaa string, bbbbbbbbbb string, cccccccccc string, dddddddddd string, eeeeeeeeee string) (string, error)
	LongFunctionNameWithOriginal(aaaaaaaaaa string, bbbbbbbbbb string, cccccccccc string, ddddddddddd string, eeeeeeeeee domain.Original) error
	LongFunctionNameWithPointer(aaaaaaaaaa []string, bbbbbbbbbb string, cccccccccc string, ddddddddddd string, eeeeeeeeee *domain.Original) error
	// LongFunctionNameNewLine is a function with a long name that takes multiple parameters.
	LongFunctionNameNewLine(
		aaaaaaaaaa string,
		bbbbbbbbbb string,
		cccccccccc string,
		dddddddddd string,
		eeeeeeeeee string,
	) (string, error)
	// Short is a function with a short name.
	Short(a string, b string) error
	// Empty is an empty function.
	Empty()
}

// MethodLLL is a struct that implements the FunctionsInterfaceLLL interface. 1111111111111111111111111111111111111111111111111111111111111
type MethodLLL struct{}

// LongFunctionName is a method with a long name that takes multiple parameters.
func (m *MethodLLL) LongFunctionNameStruct(aaaaaaaaaa string, bbbbbbbbbb string, cccccccccc string, dddddddddd string, eeeeeeeeee string) (string, error) {
	return "", nil
}

// LongFunctionNameNewLine is a method with a long name that takes multiple parameters.
func (m *MethodLLL) LongFunctionNameNewLine(
	aaaaaaaaaa string,
	bbbbbbbbbb string,
	cccccccccc string,
	dddddddddd string,
	eeeeeeeeee string,
) (string, error) {
	return "", nil
}

// Short is a function with a short name.
func (m *MethodLLL) Short(a string, b string) error {
	return nil
}

// Empty is an empty method.
func (m *MethodLLL) Empty() {}

// LongFunctionName is a function with a long name that takes multiple parameters.
func LongFunctionNameFunc(aaaaaaaaaa string, bbbbbbbbbb []string, cccccccccc string, dddddddddd string, eeeeeeeeee string) (string, error) {
	return "", nil
}

// LongFunctionNameNewLine is a function with a long name that takes multiple parameters.
func LongFunctionNameNewLineFunc(
	aaaaaaaaaa string,
	bbbbbbbbbb string,
	cccccccccc string,
	dddddddddd string,
	eeeeeeeeee string,
) (string, error) {
	return "", nil
}

// Short is a function with a short name.
func ShortFunc(a string, b string) error {
	return nil
}

// Empty is an empty function.
func EmptyFunc() {}

// 1234567890
// 1111111111222222222233333333334444444444455555555556666666666777777777788888888889999999999000000000011111111111111111111
// 111111111122222222223333333333444444444445555555555666666666677777777778888888888999999999900000000001111111111111111111
// 1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901
// 1234567890
func LLL() {
	if err := ShortFunc("aaaa", "bbbb"); err != nil {
		panic(err)
	}

	res1, err1 := LongFunctionNameFunc("aaaaaaaaaaaaaaaaaaaa", []string{"bbbbbbbbbbbbbbbbbbbb"}, "cccccccccccccccccccc", "dddddddddddddddddddd", "eeeeeeeeeeeeeeeeeeee")
	if err1 != nil {
		panic(err1)
	}

	println(res1)

	// 1111111111222222222233333333334444444444455555555556666666666777777777788888888889999999999000000000011111111111111111111
	res2, err2 := LongFunctionNameNewLineFunc(
		"aaaaaaaaaaaaaaaaaaaa",
		"bbbbbbbbbbbbbbbbbbbb",
		"cccccccccccccccccccc",
		"dddddddddddddddddddd",
		"eeeeeeeeeeeeeeeeeeee",
	)
	if err2 != nil {
		panic(err2)
	}

	println(res2)

	dddddddddddddddddddd := "d"

	res3, err3 := LongFunctionNameFunc("aaaaaaaaaaaaaaaaaaaa", []string{"bbbbbbbbbbbbbbbbbbbb"}, cccccccccccccccccccc, dddddddddddddddddddd, "eeeeeeeeeeeeeeeeeeee")
	if err3 != nil {
		panic(err3)
	}

	println(res3)

	s := &S{
		Arg: "sample",
	}

	res4, err4 := LongFunctionNameNewLineFunc(Var(), s.Method(), s.Arg, "dddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd", "eeeeeeeeeeeeeeeeeeee")
	if err4 != nil {
		panic(err4)
	}

	println(res4)
}

// 1234567890

func Var() string {
	return "sample"
}

type S struct {
	Arg string
}

func (s *S) Method() string {
	return "sample"
}

func (s *S) LongFunctionName(
	a string,
	b string,
	c string,
	d string,
	e string,
) (string, error) {
	return "", nil
}

func (s *S) LongFunctionNameNewLine(
	c string,
	d string,
	e string,
) (string, error) {
	res, err := s.LongFunctionName("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", c, d, e)
	if err != nil {
		return "", err
	}

	return res, nil
}

type FunctionsInterfaceLLL2 interface {
	LongFunctionNameWithOriginal(s string, b string, c string, d string, e string) error
}

type FunctionsInterfaceLLL2Imp struct{}

func (f *FunctionsInterfaceLLL2Imp) LongFunctionNameWithOriginal(s string, b string, c string, d string, e string) error {
	return nil
}

type TTT struct {
	lll FunctionsInterfaceLLL2
}

func (t *TTT) Func() {
	if err := t.lll.LongFunctionNameWithOriginal("aaaaaaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbbbbbb", "cccccccccccccccccccc", "dddddddddddddddddddd", "eeeeeeeeeeeeeeeeeeee"); err != nil {
		panic(err)
	}
}
