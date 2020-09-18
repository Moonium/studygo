package splitstring

import (
	"reflect"
	"testing"
)

// func TestSplit(t *testing.T) {
// 	ret := Split("babcbef", "b")
// 	want := []string{"", "a", "c", "ef"}
// 	if !reflect.DeepEqual(ret, want) { // 切片不能直接比较
// 		t.Errorf("期望得到：%v 但是得到：%v\n", want, ret)
// 	}
// }

// func Test2Split(t *testing.T) {
// 	ret := Split("a:b:c", ":")
// 	want := []string{"a", "b", "c"}
// 	if !reflect.DeepEqual(ret, want) {
// 		t.Errorf("期望得到：%v 但是得到：%v\n", want, ret)
// 	}
// }

// func Test3Split(t *testing.T) {
// 	ret := Split("abcdef", "bc")
// 	want := []string{"a", "def"}
// 	if !reflect.DeepEqual(ret, want) {
// 		t.Fatalf("期望得到：%v 但是得到：%v\n", want, ret)
// 	}
// }

// 测试组
// func TestSplit(t *testing.T) {
// 	type testCase struct {
// 		str  string
// 		sep  string
// 		want []string
// 	}

// 	testGroup := []testCase{
// 		testCase{"babcbef", "b", []string{"", "a", "c", "ef"}},
// 		testCase{"a:b:c", ":", []string{"a", "b", "c"}},
// 		testCase{"abcdef", "bc", []string{"a", "def"}},
// 		testCase{"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
// 	}

// 	for _, tc := range testGroup {
// 		got := Split(tc.str, tc.sep)
// 		if !reflect.DeepEqual(got, tc.want) {
// 			t.Fatalf("期望得到：%#v 但是得到：%#v\n", tc.want, got)
// 		}
// 	}
// }

// 子测试

func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroup := map[string]testCase{
		"case_1": testCase{"babcbef", "b", []string{"", "a", "c", "ef"}},
		"case_2": testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		"case_3": testCase{"abcdef", "bc", []string{"a", "def"}},
		"case_4": testCase{"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("期望得到：%#v 但是得到：%#v\n", tc.want, got)
			}
		})
	}
}

// BenchmarkSplit 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}

// fib_test.go

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }