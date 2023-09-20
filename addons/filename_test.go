package addons

import (
	"testing"
)

func TestGetAddonType(t *testing.T) {
	var testData = []struct {
		filename string
		flags    AddonType
	}{
		{"K_kart.pk3", KartFlag},
		{"KL_lua.lua", KartFlag | LuaFlag},
		{"SMRFCL_MyCreativeFileName-v2.0.pk3", SinglePlayerFlag | MatchFlag | RaceFlag | FlagFlag | CharFlag | LuaFlag},
		{"KRBCL_MyKartContent-v1.0.pk3", KartFlag | RaceFlag | BattleFlag | CharFlag | LuaFlag},
		{"KR_MyMap-v1.wad", KartFlag | RaceFlag},
		{"KC_MyCharacters-v2.pk3", KartFlag | CharFlag},
		{"KRB_MyMapPack_v1.wad", KartFlag | RaceFlag | BattleFlag},
		{"KRBC_MyMapAndCharacterPack_v2.2.pk3", KartFlag | RaceFlag | BattleFlag | CharFlag},
		{"KRBCL_MySuperBigGameplayModPack-v8.8.pk3", KartFlag | RaceFlag | BattleFlag | CharFlag | LuaFlag},
	}

	for _, data := range testData {
		ret := GetAddonType(data.filename)
		if ret != data.flags {
			t.Fatalf("Expected flags %b but got %b for filename %s", data.flags, ret, data.filename)
		}
	}
}

func TestGetAddonVersion(t *testing.T) {
	var testData = []struct {
		file    string
		version []uint
	}{
		{"L_My_Gameplay_Mod-v1.wad", []uint{1}},
		{"S_My-Ported-Map_v2.wad", []uint{2}},
		{"C_MyCharacter_v1.4.pk3", []uint{1, 4}},
		{"SCL_MyLuaCharacterMap_v5.pk3", []uint{5}},
		{"KBL_MK8D_Urchin_Underpass_v1.5.2.wad", []uint{1, 5, 2}},
		{"KL_DriftNitro_v2.2.12.pk3", []uint{2, 2, 12}},
		{"KC_CCCP_v1.2.1.pk3", []uint{1, 2, 1}},
		// some weirdos
		{"KBL_Reef_v1h.pk3", []uint{1}},
		{"KCL_PizzaDarCharPack-R12.pk3", []uint{12}},
		//{"KC_AaronPack-v1-5.pk3  ", []uint{1, 5}}, too wrong
		{"KC_Amigo_PD25.pk3", []uint{25}},
		{"KC_BONANZABROS_v1.0 .pk3", []uint{1, 0}},
		{"KC_Cole.pk3", []uint{}},
		{"KC_Dominos_Funny_HehehahaV4.pk3", []uint{4}},
		{"KC_EfiniRx7-V1.1.pk3", []uint{1, 1}},
		{"KC_FlameRunnerFunky_InsDrift.pk3", []uint{}},
		{"KC_GranbluePack-v1.01.pk3", []uint{1, 1}},
		{"KC_JokerCharsPack-v6.1.PLUS.pk3", []uint{6, 1}},
		{"KC_Linerider-v.1.wad", []uint{1}},
		{"KC_MLP_Pack_v1a.pk3", []uint{1}}, // man screw letters
		// {"KC_PlomChars-v1-1b.pk3", []uint{1, 1}}, to wrong
		{"KC_SDHornet_V1noengine.pk3", []uint{1}},
	}

	for _, d := range testData {
		ret := GetAddonVersion(d.file)
		fail := func() {
			t.Fatalf("Expected version %v, but got %v for filename %s", d.version, ret, d.file)
		}

		if len(ret) != len(d.version) {
			fail()
		}

		for i := range ret {
			if ret[i] != d.version[i] {
				fail()
			}
		}
	}
}
