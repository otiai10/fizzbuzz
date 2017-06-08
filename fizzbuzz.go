package fizzbuzz

import "fmt"

// GetOutput 関数
// パッケージ "fizzbuzz" の、外部から参照される関数の名前は、
// かならず大文字から始める必要があります。
func GetOutput(i int) string {
	fizz := shouldFizz(i)
	buzz := shouldBuzz(i)
	// Go言語のswitch-case文は、明示的にbreakしなくても
	// 下のcase句へ落ちて **いかない** という特徴があります。
	switch {
	case fizz && buzz:
		return "FizzBuzz"
	case fizz:
		return "fizz"
	case buzz:
		return "buzz"
	default:
		return fmt.Sprintf("%v", i)
	}
}

// 逆に、外部から fizzbuzz.XXX という形で参照される必要の無いものは
// その名前を小文字から始めればよいです。
// いわゆるパッケージプライベートな関数です。
func shouldFizz(i int) bool {
	return i%3 == 0
}
func shouldBuzz(i int) bool {
	return i%5 == 0
}

// fizzbuzzパッケージを使う側から `fizzbuzz.shouldFizz(100)`
// と参照しようとするとコンパイルエラーになります。
