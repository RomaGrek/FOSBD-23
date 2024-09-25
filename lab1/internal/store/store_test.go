package store

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

// Функция тестирования процесса иницализации хранилища
func TestNewStore(t *testing.T) {
	tmpDBFile, err := os.CreateTemp("", "example-*.data")
	require.NoError(t, err)
	defer func() {
		err = os.Remove(tmpDBFile.Name())
		require.NoError(t, err)
	}()

	prevFileSize := getSizeFile(t, tmpDBFile)
	require.Equal(t, prevFileSize, int64(0))

	err = tmpDBFile.Close()
	require.NoError(t, err)

	stor := NewStore(tmpDBFile.Name(), testLogger(t))
	require.Equal(t, stor.globalDepth, defaultGlobalDepth)
	require.Equal(t, stor.pathToDB, tmpDBFile.Name())
	require.Equal(t, len(stor.dirList), 2)
	require.Equal(t, stor.endOffset, 2*pageSize)

	tmpDBFile, err = os.Open(tmpDBFile.Name())
	require.NoError(t, err)
	require.Equal(t, getSizeFile(t, tmpDBFile), 2*int64(pageSize))

	err = tmpDBFile.Close()
	require.NoError(t, err)
}

// Функция тестирования основных операций хранилища
func TestStore(t *testing.T) {
	tmpDBFile, err := os.CreateTemp("", "example-*.data")
	require.NoError(t, err)
	defer func() {
		err = os.Remove(tmpDBFile.Name())
		require.NoError(t, err)
	}()

	prevFileSize := getSizeFile(t, tmpDBFile)
	require.Equal(t, prevFileSize, int64(0))

	err = tmpDBFile.Close()
	require.NoError(t, err)

	stor := NewStore(tmpDBFile.Name(), testLogger(t))

	for _, tc := range []struct {
		name string
		key  string
		val  string
	}{
		{
			name: "first record",
			key:  "roma",
			val:  "dolznik",
		},
		{
			name: "second record",
			key:  "petia",
			val:  "kekus",
		},
		{
			name: "third record",
			key:  "ivan",
			val:  "kefkus",
		},
		{
			name: "fourth record",
			key:  "igor",
			val:  "feff",
		},
		{
			name: "five record",
			key:  "sima",
			val:  "fefff",
		},
		{
			name: "six record",
			key:  "sanek",
			val:  "fefffwff",
		},
		{
			name: "seven record",
			key:  "misha",
			val:  "fefffewff",
		},
		{
			name: "8 record",
			key:  "liza",
			val:  "fefppff",
		},
		{
			name: "9",
			key:  "gaika",
			val:  "itir",
		},
		{
			name: "9",
			key:  "poet",
			val:  "ffffff",
		},
		{
			name: "wefd",
			key:  "puskin",
			val:  "wefewfewf",
		},
		{
			name: "kok",
			key:  "kok",
			val:  "feee",
		},
		{
			name: "pow",
			key:  "feeeewf",
			val:  "feee",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err = stor.SetValue(tc.key, tc.val) // кладем значения
			require.NoError(t, err)

			val, err := stor.GetValue(tc.key) // получаем значения
			require.NoError(t, err)
			require.Equal(t, tc.val, val)
		})
	}
}

// Функция помошник для опеределения размера файла
func getSizeFile(t *testing.T, file *os.File) int64 {
	fInfo, err := file.Stat()
	require.NoError(t, err)
	return fInfo.Size()
}

func testLogger(t *testing.T) *zap.Logger {
	logger, err := zap.NewDevelopment()
	require.NoError(t, err)

	return logger
}
