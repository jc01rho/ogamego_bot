package Test

import (
	"github.com/0xE232FE/ogame.mod"
	"reflect"
	"testing"
)

func TestGetBotForTest(t *testing.T) {
	var tests = []struct {
		name string
		want *ogame.OGame
	}{
		// TODO: Add test cases.
		{"test",nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBotForTest(); !reflect.DeepEqual(got, tt.want) {

				t.Logf("GetBotForTest() = %v, want %v", got, tt.want)
			}
		})
	}
}