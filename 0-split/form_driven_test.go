package split

import (
	"reflect"
	"testing"
)

// 官方标准库中有很多表格驱动测试的示例，例如fmt包中的测试代码
//var flagtests = []struct {
//	in  string
//	out string
//}{
//	{"%a", "[%a]"},
//	{"%-a", "[%-a]"},
//	{"%+a", "[%+a]"},
//	{"%#a", "[%#a]"},
//	{"% a", "[% a]"},
//	{"%0a", "[%0a]"},
//	{"%1.2a", "[%1.2a]"},
//	{"%-1.2a", "[%-1.2a]"},
//	{"%+1.2a", "[%+1.2a]"},
//	{"%-+1.2a", "[%+-1.2a]"},
//	{"%-+1.2abc", "[%+-1.2a]bc"},
//	{"%-1.2abc", "[%-1.2a]bc"},
//}
//
//func TestFlagParser(t *testing.T) {
//	var flagprinter flagPrinter
//	for _, tt := range flagtests {
//		t.Run(tt.in, func(t *testing.T) {
//			s := Sprintf(tt.in, &flagprinter)
//			if s != tt.out {
//				t.Errorf("got %q, want %q", s, tt.out)
//			}
//		})
//	}
//}

func TestSplitAll(t *testing.T) {
	// 定义测试表格
	// 这里使用匿名结构体定义了若干个测试用例
	// 并且为每个测试用例设置了一个名称
	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{"base case", "a:b:c", ":", []string{"a", "b", "c"}},
		{"wrong sep", "a:b:c", ",", []string{"a:b:c"}},
		{"more sep", "abcd", "bc", []string{"a", "d"}},
		{"leading sep", "沙河有沙又有河", "沙", []string{"", "河有", "又有河"}},
	}
	// 遍历测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Split(tt.input, tt.sep)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected:%#v, got:%#v", tt.want, got)
			}
		})
	}
}

func TestSplitAllParallel(t *testing.T) {
	t.Parallel() // 将 TLog 标记为能够与其他测试并行运行
	// 定义测试表格
	// 这里使用匿名结构体定义了若干个测试用例
	// 并且为每个测试用例设置了一个名称
	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{"base case", "a:b:c", ":", []string{"a", "b", "c"}},
		{"wrong sep", "a:b:c", ",", []string{"a:b:c"}},
		{"more sep", "abcd", "bc", []string{"a", "d"}},
		{"leading sep", "沙河有沙又有河", "沙", []string{"", "河有", "又有河"}},
	}
	// 遍历测试用例
	for _, tt := range tests {
		tt := tt                            // 注意这里重新声明tt变量（避免多个goroutine中使用了相同的变量）
		t.Run(tt.name, func(t *testing.T) { // 使用t.Run()执行子测试
			t.Parallel() // 将每个测试用例标记为能够彼此并行运行
			got := Split(tt.input, tt.sep)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected:%#v, got:%#v", tt.want, got)
			}
		})
	}
}
