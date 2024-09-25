package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	lenKV          = 1365
	maxLenUsrKey   = 127
	maxLenUsrValue = 1235
)

// Функция проверки корректности парсинга и распарсинга
// Концепция проста - сначала сообщение кодируется в бинарный вид.
// Потом анмаршалиться и если ключ и значение совпали - значит все ок.
func TestParseKV(t *testing.T) {
	for _, tc := range []struct {
		name  string
		key   string
		value string
	}{
		{
			name:  "just word",
			key:   "rome",
			value: "dopsa",
		},
		{
			name:  "just russian word",
			key:   "рома",
			value: "должник",
		},
		{
			name:  "empty value",
			key:   "test-key",
			value: "",
		},
		{
			name:  "max key",
			key:   strings.Repeat("s", maxLenUsrKey),
			value: "test",
		},
		{
			name:  "max value",
			key:   "test",
			value: strings.Repeat("s", maxLenUsrValue),
		},
		{
			name:  "max key and value",
			key:   strings.Repeat("s", maxLenUsrKey),
			value: strings.Repeat("s", maxLenUsrValue),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			dataKV, err := MarshalKV(tc.key, tc.value)
			require.NoError(t, err)
			require.Equal(t, lenKV, len(dataKV))
			require.Equal(t, lenKV, cap(dataKV))

			parsedKey, parsedValue, err := UnmarshalKV(dataKV)
			require.NoError(t, err)
			require.Equal(t, tc.key, parsedKey)
			require.Equal(t, tc.value, parsedValue)
		})
	}
}
