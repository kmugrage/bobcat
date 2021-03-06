package interpreter

import "testing"

func Benchmark_LoadFile_For_OneThousandEntities(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance/1_thousand.lang", NewRootScope())
}

func Benchmark_LoadFile_For_TenThousandEntities(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance/10_thousand.lang", NewRootScope())
}

func Benchmark_LoadFile_For_FiftyThousandEntities(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance/50_thousand.lang", NewRootScope())
}

func Benchmark_LoadFile_For_SeventyThousandEntities(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance/70_thousand.lang", NewRootScope())
}

func Benchmark_LoadFile_For_EightyThousandEntities(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance/80_thousand.lang", NewRootScope())
}

func Benchmark_LoadFile_For_NinetyThousandEntities(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance/90_thousand.lang", NewRootScope())
}

func Benchmark_LoadFile_For_OnehundredThousandEntities(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance/100_thousand.lang", NewRootScope())
}

func Benchmark_LoadFile_For_OneMillionEntities(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance/1_million.lang", NewRootScope())
}

func Benchmark_LoadFile_For_OneThousandEntitiesWithDictionary(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance_with_dict/1_thousand.lang", NewRootScope())
}

func Benchmark_LoadFile_For_TenThousandEntitiesWithDictionary(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance_with_dict/10_thousand.lang", NewRootScope())
}

func Benchmark_LoadFile_For_FiftyThousandEntitiesWithDictionary(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance_with_dict/50_thousand.lang", NewRootScope())
}

func Benchmark_LoadFile_For_SeventyThousandEntitiesWithDictionary(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance_with_dict/70_thousand.lang", NewRootScope())
}

func Benchmark_LoadFile_For_EightyThousandEntitiesWithDictionary(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance_with_dict/80_thousand.lang", NewRootScope())
}

func Benchmark_LoadFile_For_NinetyThousandEntitiesWithDictionary(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance_with_dict/90_thousand.lang", NewRootScope())
}

func Benchmark_LoadFile_For_OnehundredThousandEntitiesWithDictionary(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance_with_dict/100_thousand.lang", NewRootScope())
}

func Benchmark_LoadFile_For_OneMillionEntitiesWithDictionary(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance_with_dict/1_million.lang", NewRootScope())
}

func Benchmark_LoadFile_For_OneThousandEntitiesWithCustomDictionary(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance_with_custom_dict/1_thousand.lang", NewRootScope())
}

func Benchmark_LoadFile_For_TenThousandEntitiesWithCustomDictionary(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance_with_custom_dict/10_thousand.lang", NewRootScope())
}

func Benchmark_LoadFile_For_FiftyThousandEntitiesWithCustomDictionary(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance_with_custom_dict/50_thousand.lang", NewRootScope())
}

func Benchmark_LoadFile_For_SeventyThousandEntitiesWithCustomDictionary(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance_with_custom_dict/70_thousand.lang", NewRootScope())
}

func Benchmark_LoadFile_For_EightyThousandEntitiesWithCustomDictionary(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance_with_custom_dict/80_thousand.lang", NewRootScope())
}

func Benchmark_LoadFile_For_NinetyThousandEntitiesWithCustomDictionary(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance_with_custom_dict/90_thousand.lang", NewRootScope())
}

func Benchmark_LoadFile_For_OnehundredThousandEntitiesWithCustomDictionary(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance_with_custom_dict/100_thousand.lang", NewRootScope())
}

func Benchmark_LoadFile_For_OneMillionEntitiesWithCustomDictionary(b *testing.B) {
	i := New()
	i.LoadFile("testdata/performance_with_custom_dict/1_million.lang", NewRootScope())
}
