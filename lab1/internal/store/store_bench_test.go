package store

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

// Бенчмарк на загрузку одного значения при изначально пустой бд
func BenchmarkSetValue1(b *testing.B) {
	tmpDBFile, err := os.CreateTemp("", "example-*.data")
	require.NoError(b, err)
	defer func() {
		err = os.Remove(tmpDBFile.Name())
		require.NoError(b, err)
	}()

	err = tmpDBFile.Close()
	require.NoError(b, err)

	logger, _ := zap.NewDevelopment()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		stor := NewStore(tmpDBFile.Name(), logger)
		b.StartTimer()
		
		err = stor.SetValue("key", "val")
		_ = err
	}
}

// Бенчмарк на загрузку 5-ти значений
func BenchmarkSetValue5(b *testing.B) {
	tmpDBFile, err := os.CreateTemp("", "example-*.data")
	require.NoError(b, err)
	defer func() {
		err = os.Remove(tmpDBFile.Name())
		require.NoError(b, err)
	}()

	err = tmpDBFile.Close()
	require.NoError(b, err)

	logger, _ := zap.NewDevelopment()

	testKeys := []string{"roma", "lesha", "vlad", "pema", "linux"}
	testVal := []string{"leop", "doner", "pilorama", "agent", "windows"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		stor := NewStore(tmpDBFile.Name(), logger)
		b.StartTimer()

		for i := 0; i < 5; i++ {
			err = stor.SetValue(testKeys[i], testVal[i])
			_ = err
		}
	}
}

// Бенчмарк на загрузку 10-ти значений
func BenchmarkSetValue10(b *testing.B) {
	tmpDBFile, err := os.CreateTemp("", "example-*.data")
	require.NoError(b, err)
	defer func() {
		err = os.Remove(tmpDBFile.Name())
		require.NoError(b, err)
	}()

	err = tmpDBFile.Close()
	require.NoError(b, err)

	logger, _ := zap.NewDevelopment()

	testKeys := []string{"roma", "lesha", "vlad", "pema", "linux", "chek", "poet", "lev", "volk", "cats"}
	testVal := []string{"leop", "doner", "pilorama", "agent", "windows", "mavos", "genos", "lol", "kek", "dogs"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		stor := NewStore(tmpDBFile.Name(), logger)
		b.StartTimer()

		for i := 0; i < 10; i++ {
			err = stor.SetValue(testKeys[i], testVal[i])
			_ = err
		}
	}
}

// Бенчмарк на загрузку 15-ти значений
func BenchmarkSetValue15(b *testing.B) {
	tmpDBFile, err := os.CreateTemp("", "example-*.data")
	require.NoError(b, err)
	defer func() {
		err = os.Remove(tmpDBFile.Name())
		require.NoError(b, err)
	}()

	err = tmpDBFile.Close()
	require.NoError(b, err)

	logger, _ := zap.NewDevelopment()

	testKeys := []string{"roma", "lesha", "vlad", "pema", "linux", "chek", "poet", "lev", "volk", "cats", "torvals", "viking", "micrk", "leon", "five"}
	testVal := []string{"leop", "doner", "pilorama", "agent", "windows", "mavos", "genos", "lol", "kek", "dogs", "cmel", "shmel", "tron", "lhal", "drakon"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		stor := NewStore(tmpDBFile.Name(), logger)
		b.StartTimer()
		
		for i := 0; i < 15; i++ {
			err = stor.SetValue(testKeys[i], testVal[i])
			_ = err
		}
	}
}

// Бенчмарк на загрузку 20-ти значений
func BenchmarkSetValue20(b *testing.B) {
	tmpDBFile, err := os.CreateTemp("", "example-*.data")
	require.NoError(b, err)
	defer func() {
		err = os.Remove(tmpDBFile.Name())
		require.NoError(b, err)
	}()

	err = tmpDBFile.Close()
	require.NoError(b, err)

	logger, _ := zap.NewDevelopment()

	testKeys := []string{"roma", "lesha", "vlad", "pema", "linux", "chek", "poet", "lev", "volk", "cats", "torvals", "viking", "micrk", "leon", "five", "mem", "orbidol", "zabolel", "vizdorovel", "eooe"}
	testVal := []string{"leop", "doner", "pilorama", "agent", "windows", "mavos", "genos", "lol", "kek", "dogs", "cmel", "shmel", "tron", "lhal", "drakon", "eee", "ooo", "kkkk", "eeeee", "wefwef"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		stor := NewStore(tmpDBFile.Name(), logger)
		b.StartTimer()
		
		for i := 0; i < 20; i++ {
			err = stor.SetValue(testKeys[i], testVal[i])
			_ = err
		}
	}
}

// Бенчмарк на получение значения когда в бд 1 значение
func BenchmarkGetValue1(b *testing.B) {
	tmpDBFile, err := os.CreateTemp("", "example-*.data")
	require.NoError(b, err)
	defer func() {
		err = os.Remove(tmpDBFile.Name())
		require.NoError(b, err)
	}()

	err = tmpDBFile.Close()
	require.NoError(b, err)

	logger, _ := zap.NewDevelopment()
	stor := NewStore(tmpDBFile.Name(), logger)
	
	err = stor.SetValue("key", "val")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val, err := stor.GetValue("key")
		_, _ = val, err
	}
}

// Бенчмарк на получение значения когда в бд 5 значений
func BenchmarkGetValue5(b *testing.B) {
	tmpDBFile, err := os.CreateTemp("", "example-*.data")
	require.NoError(b, err)
	defer func() {
		err = os.Remove(tmpDBFile.Name())
		require.NoError(b, err)
	}()

	err = tmpDBFile.Close()
	require.NoError(b, err)

	logger, _ := zap.NewDevelopment()
	stor := NewStore(tmpDBFile.Name(), logger)

	testKeys := []string{"roma", "lesha", "vlad", "pema", "linux"}
	testVal := []string{"leop", "doner", "pilorama", "agent", "windows"}

	for i := 0; i < 5; i++ {
		err = stor.SetValue(testKeys[i], testVal[i])
		require.NoError(b, err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val, err := stor.GetValue("roma")
		_, _ = val, err
	}
}

// Бенчмарк на получение значения когда в бд 10 значений
func BenchmarkGetValue10(b *testing.B) {
	tmpDBFile, err := os.CreateTemp("", "example-*.data")
	require.NoError(b, err)
	defer func() {
		err = os.Remove(tmpDBFile.Name())
		require.NoError(b, err)
	}()

	err = tmpDBFile.Close()
	require.NoError(b, err)

	logger, _ := zap.NewDevelopment()
	stor := NewStore(tmpDBFile.Name(), logger)

	testKeys := []string{"roma", "lesha", "vlad", "pema", "linux", "chek", "poet", "lev", "volk", "cats"}
	testVal := []string{"leop", "doner", "pilorama", "agent", "windows", "mavos", "genos", "lol", "kek", "dogs"}

	for i := 0; i < 10; i++ {
		err = stor.SetValue(testKeys[i], testVal[i])
		require.NoError(b, err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val, err := stor.GetValue("volk")
		_, _ = val, err
	}
}

// Бенчмарк на получение значения когда в бд 15 значений
func BenchmarkGetValue15(b *testing.B) {
	tmpDBFile, err := os.CreateTemp("", "example-*.data")
	require.NoError(b, err)
	defer func() {
		err = os.Remove(tmpDBFile.Name())
		require.NoError(b, err)
	}()

	err = tmpDBFile.Close()
	require.NoError(b, err)

	logger, _ := zap.NewDevelopment()
	stor := NewStore(tmpDBFile.Name(), logger)

	testKeys := []string{"roma", "lesha", "vlad", "pema", "linux", "chek", "poet", "lev", "volk", "cats", "torvals", "viking", "micrk", "leon", "five"}
	testVal := []string{"leop", "doner", "pilorama", "agent", "windows", "mavos", "genos", "lol", "kek", "dogs", "cmel", "shmel", "tron", "lhal", "drakon"}

	for i := 0; i < 15; i++ {
		err = stor.SetValue(testKeys[i], testVal[i])
		require.NoError(b, err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val, err := stor.GetValue("volk")
		_, _ = val, err
	}
}

// Бенчмарк на получение значения когда в бд 20 значений
func BenchmarkGetValue20(b *testing.B) {
	tmpDBFile, err := os.CreateTemp("", "example-*.data")
	require.NoError(b, err)
	defer func() {
		err = os.Remove(tmpDBFile.Name())
		require.NoError(b, err)
	}()

	err = tmpDBFile.Close()
	require.NoError(b, err)

	logger, _ := zap.NewDevelopment()
	stor := NewStore(tmpDBFile.Name(), logger)

	testKeys := []string{"roma", "lesha", "vlad", "pema", "linux", "chek", "poet", "lev", "volk", "cats", "torvals", "viking", "micrk", "leon", "five", "mem", "orbidol", "zabolel", "vizdorovel", "eooe"}
	testVal := []string{"leop", "doner", "pilorama", "agent", "windows", "mavos", "genos", "lol", "kek", "dogs", "cmel", "shmel", "tron", "lhal", "drakon", "eee", "ooo", "kkkk", "eeeee", "wefwef"}

	for i := 0; i < 20; i++ {
		err = stor.SetValue(testKeys[i], testVal[i])
		require.NoError(b, err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val, err := stor.GetValue("volk")
		_, _ = val, err
	}
}