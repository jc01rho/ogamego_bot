package OgameUtil

import (
	"bitbucket.org/jc01rho/ogame"
	"github.com/jc01rho/ogamego_bot/OGameBotRoutine"
	"github.com/jc01rho/ogamego_bot/Test"
	"reflect"
	"testing"
)

func TestGetNextResBuildingOnline(t *testing.T) {
	type args struct {
		bot *ogame.OGame
	}
	tests := []struct {
		name  string
		args  args
		want  *ogame.Planet
		want1 *ogame.BaseBuilding
	}{
		// TODO: Add test cases.
		{"testOnline", struct{ bot *ogame.OGame }{bot: Test.GetBotForTest()}, nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := OGameBotRoutine.GetNextResBuilding(tt.args.bot)
			if !reflect.DeepEqual(got.ID, ogame.PlanetID(34080740)) {
				t.Errorf("GetNextResBuilding() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1.Base.ID, ogame.MetalMineID) {
				t.Errorf("GetNextResBuilding() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, int64(27)) {
				t.Errorf("GetNextResBuilding() got2 = %v, want %v", got2, tt.want)
			}

		})
	}
}
