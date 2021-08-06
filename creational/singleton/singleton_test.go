package singleton

import (
	"testing"
)

// 测试饿汉式 并发安全
func BenchmarkGetInstance(b *testing.B) {
	// 并发测试
	b.RunParallel(func(pb *testing.PB) {
		b.Logf("%p", GetInstance())
	})
}

// 测试懒汉模式 并发不安全
func BenchmarkGetLazyInstance(b *testing.B) {
	// 通过答打印的 返回的每一个输出地址，就知道是并发不安全的
	b.RunParallel(func(pb *testing.PB) {
		b.Logf("%p", GetLazyInstance())
	})
}

// 利用 双重检测
func BenchmarkGetLazyInstanceV2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		b.Logf("%p", GetLazyInstanceV2())
	})
}

// 利用 golang的特性 once.Do
func BenchmarkGetLazyInstanceV3(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		b.Logf("%p", GetLazyInstanceV3())
	})
}
