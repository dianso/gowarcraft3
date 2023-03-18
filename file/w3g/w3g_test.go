// Author:  Niels A.D.
// Project: gowarcraft3 (https://github.com/nielsAD/gowarcraft3)
// License: Mozilla Public License, v2.0

package w3g_test

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"testing"

	"github.com/nielsAD/gowarcraft3/file/w3g"
	"github.com/nielsAD/gowarcraft3/protocol"
	"github.com/nielsAD/gowarcraft3/protocol/w3gs"
)

func b64(s string) []byte {
	b, _ := base64.StdEncoding.DecodeString(s)
	return b
}

func Example() {
	replay, err := w3g.Open("./test_130.w3g")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(replay.HostPlayer.Name)

	// output:
	// niels
}

func TestFiles(t *testing.T) {
	var files = []struct {
		file   string
		replay w3g.Replay
	}{
		{
			"test_102.w3g",
			w3g.Replay{
				Header: w3g.Header{
					GameVersion: w3gs.GameVersion{Product: w3gs.ProductROC, Version: 2},
					BuildNumber: 4531,
					DurationMS:  441925,
				},
				GameInfo: w3g.GameInfo{
					HostPlayer: w3g.PlayerInfo{
						ID:   1,
						Name: "Go4WC3.Sapor",
					},
					GameName: "final",
					GameSettings: w3gs.GameSettings{
						GameSettingFlags: w3gs.SettingSpeedFast | w3gs.SettingTerrainDefault | w3gs.SettingObsFull | w3gs.SettingTeamsTogether | w3gs.SettingTeamsFixed,
						MapWidth:         124,
						MapHeight:        124,
						MapXoro:          325121041,
						MapPath:          "Maps\\(4)LostTemple.w3m",
						HostName:         "Go4WC3.Sapor",
					},
					GameFlags:  w3gs.GameFlagCustomGame | w3gs.GameFlagSignedMap | w3gs.GameFlagPrivateGame,
					NumSlots:   12,
					LanguageID: 7206592,
				},
				SlotInfo: w3g.SlotInfo{
					SlotInfo: w3gs.SlotInfo{
						Slots: []w3gs.SlotData{
							w3gs.SlotData{
								PlayerID:       2,
								DownloadStatus: 100,
								SlotStatus:     w3gs.SlotOccupied,
								Team:           12,
								Color:          12,
								Race:           w3gs.RaceRandom,
								ComputerType:   w3gs.ComputerNormal,
								Handicap:       100,
							}},
						RandomSeed: 45792916,
						SlotLayout: w3gs.LayoutMelee,
						NumPlayers: 4,
					},
				},
				PlayerInfo: []*w3g.PlayerInfo{
					&w3g.PlayerInfo{
						ID:   2,
						Name: "Go4WC3.Desann",
					},
				},
				Records: []w3g.Record{
					&w3g.TimeSlot{TimeSlot: w3gs.TimeSlot{TimeIncrementMS: 100}},
					&w3g.TimeSlot{TimeSlot: w3gs.TimeSlot{TimeIncrementMS: 100, Actions: []w3gs.PlayerAction{w3gs.PlayerAction{PlayerID: 10, Data: b64("FgIBAIlLAACMSwAAFgEFAKtLAACuSwAAw0sAAMZLAADbSwAA3ksAAPNLAAD2SwAAC0wAAA5MAAA=")}}}},
					&w3g.TimeSlot{TimeSlot: w3gs.TimeSlot{TimeIncrementMS: 100, Actions: []w3gs.PlayerAction{w3gs.PlayerAction{PlayerID: 9, Data: b64("FgIBAMBKAADASgAAFgEFABRLAAAXSwAAKksAAC1LAABASwAAQ0sAAFZLAABZSwAAbEsAAG9LAAAZAA==")}, w3gs.PlayerAction{PlayerID: 10, Data: b64("EggDAA0AAACAxAAA2EXfQAAA30AAAA==")}}}},
				},
			},
		},
		{
			"test_126.w3g",
			w3g.Replay{
				Header: w3g.Header{
					GameVersion: w3gs.GameVersion{Product: w3gs.ProductTFT, Version: 26},
					BuildNumber: 6059,
					DurationMS:  237175,
				},
				GameInfo: w3g.GameInfo{
					HostPlayer: w3g.PlayerInfo{
						ID:          1,
						Name:        "ForFunyo",
						Race:        w3gs.RaceUndead,
						JoinCounter: 3052346,
					},
					GameName: "BNet",
					GameSettings: w3gs.GameSettings{
						GameSettingFlags: w3gs.SettingSpeedFast | w3gs.SettingTerrainExplored | w3gs.SettingObsNone | w3gs.SettingTeamsTogether | w3gs.SettingTeamsFixed,
						MapWidth:         0,
						MapHeight:        0,
						MapXoro:          4294967295,
						MapPath:          "Maps\\FrozenThrone\\(4)TurtleRock.w3x",
						HostName:         "Battle.net",
						MapSha1:          [20]byte{0, 145, 213, 254, 70, 60, 116, 93, 232, 133, 235, 135, 140, 210, 168, 35, 212, 189, 5, 97},
					},
					GameFlags:  w3gs.GameFlagCustomGame,
					NumSlots:   2,
					LanguageID: 1636528,
				},
				SlotInfo: w3g.SlotInfo{
					SlotInfo: w3gs.SlotInfo{
						Slots: []w3gs.SlotData{
							w3gs.SlotData{
								PlayerID:       2,
								DownloadStatus: 255,
								SlotStatus:     w3gs.SlotOccupied,
								Team:           1,
								Color:          0,
								Race:           w3gs.RaceUndead,
								ComputerType:   w3gs.ComputerNormal,
								Handicap:       100,
							}},
						RandomSeed: 77005536,
						SlotLayout: w3gs.LayoutLadder,
						NumPlayers: 4,
					},
				},
				PlayerInfo: []*w3g.PlayerInfo{
					&w3g.PlayerInfo{
						ID:          2,
						Name:        "Fighting-",
						Race:        w3gs.RaceUndead,
						JoinCounter: 31881612,
					},
				},
				Records: []w3g.Record{
					&w3g.TimeSlot{TimeSlot: w3gs.TimeSlot{TimeIncrementMS: 251, Actions: []w3gs.PlayerAction{w3gs.PlayerAction{PlayerID: 1, Data: b64("EgAAAwANAP//////////AAAgxQAApEU3QQAAQ0EAAA==")}}}},
					&w3g.TimeSlot{TimeSlot: w3gs.TimeSlot{TimeIncrementMS: 250, Actions: []w3gs.PlayerAction{w3gs.PlayerAction{PlayerID: 1, Data: b64("EgAAAwANAP//////////AAAgxQAApEU3QQAAQ0EAAA==")}, w3gs.PlayerAction{PlayerID: 1, Data: b64("EgAAAwANAP//////////AAAgxQAApEU3QQAAQ0EAAA==")}}}},
					&w3g.TimeSlot{TimeSlot: w3gs.TimeSlot{TimeIncrementMS: 250}},
				},
			},
		},
		{
			"test_130.w3g",
			w3g.Replay{
				Header: w3g.Header{
					GameVersion:  w3gs.GameVersion{Product: w3gs.ProductTFT, Version: 10030},
					BuildNumber:  6061,
					DurationMS:   640650,
					SinglePlayer: true,
				},
				GameInfo: w3g.GameInfo{
					HostPlayer: w3g.PlayerInfo{
						ID:   1,
						Name: "niels",
					},
					GameName: "Local Game",
					GameSettings: w3gs.GameSettings{
						GameSettingFlags: w3gs.SettingSpeedFast | w3gs.SettingTerrainDefault | w3gs.SettingObsNone | w3gs.SettingTeamsTogether | w3gs.SettingTeamsFixed,
						MapWidth:         116,
						MapHeight:        84,
						MapXoro:          2599102717,
						MapPath:          "Maps/FrozenThrone//(2)EchoIsles.w3x",
						HostName:         "niels",
						MapSha1:          [20]byte{107, 111, 100, 67, 248, 197, 26, 44, 89, 111, 217, 78, 123, 106, 91, 101, 208, 6, 70, 129},
					},
					GameFlags: w3gs.GameFlagCustomGame | w3gs.GameFlagSignedMap,
					NumSlots:  24,
				},
				SlotInfo: w3g.SlotInfo{
					SlotInfo: w3gs.SlotInfo{
						Slots: []w3gs.SlotData{
							w3gs.SlotData{
								PlayerID:       0,
								DownloadStatus: 100,
								SlotStatus:     w3gs.SlotOccupied,
								Computer:       true,
								Team:           1,
								Color:          1,
								Race:           w3gs.RaceRandom | w3gs.RaceSelectable,
								ComputerType:   w3gs.ComputerNormal,
								Handicap:       100,
							}},
						RandomSeed: 40053178,
						SlotLayout: w3gs.LayoutMelee,
						NumPlayers: 2,
					},
				},
				PlayerInfo: []*w3g.PlayerInfo{
					&w3g.PlayerInfo{
						ID:   1,
						Name: "niels",
					},
				},
				Records: []w3g.Record{
					&w3g.TimeSlot{TimeSlot: w3gs.TimeSlot{TimeIncrementMS: 0, Actions: []w3gs.PlayerAction{w3gs.PlayerAction{PlayerID: 1, Data: b64("EhgAAwANAP//////////AACwxQAAYEXMMQAAzDEAAA==")}}}},
					&w3g.TimeSlot{TimeSlot: w3gs.TimeSlot{TimeIncrementMS: 100}},
					&w3g.TimeSlot{TimeSlot: w3gs.TimeSlot{TimeIncrementMS: 100}},
				},
			},
		},
		{
			"test_132.w3g",
			w3g.Replay{
				Header: w3g.Header{
					GameVersion: w3gs.GameVersion{Product: w3gs.ProductTFT, Version: 10032},
					BuildNumber: 6105,
					DurationMS:  503575,
				},
				GameInfo: w3g.GameInfo{
					HostPlayer: w3g.PlayerInfo{
						ID:   2,
						Name: "TheBiGsLeeP#220",
					},
					GameName: "BNet",
					GameSettings: w3gs.GameSettings{
						GameSettingFlags: w3gs.SettingSpeedFast | w3gs.SettingTerrainDefault | w3gs.SettingObsEnabled | w3gs.SettingObsReferees | w3gs.SettingTeamsTogether | w3gs.SettingTeamsFixed,
						MapXoro:          4294967295,
						MapPath:          "Maps/Download/ddee38acb17ca0c372e271ef0dbe684a78d42eb4/(2)Amazonia.w3x",
						HostName:         "Battle.net",
						MapSha1:          [20]byte{221, 238, 56, 172, 177, 124, 160, 195, 114, 226, 113, 239, 13, 190, 104, 74, 120, 212, 46, 180},
					},
					GameFlags: w3gs.GameFlagCreatorBlizzard | w3gs.GameFlagObsFull | w3gs.GameFlags(0x10),
					NumSlots:  3,
				},
				SlotInfo: w3g.SlotInfo{
					SlotInfo: w3gs.SlotInfo{
						Slots: []w3gs.SlotData{
							w3gs.SlotData{
								PlayerID:       3,
								DownloadStatus: 100,
								SlotStatus:     w3gs.SlotOccupied,
								Team:           1,
								Color:          1,
								Race:           w3gs.RaceUndead,
								Handicap:       100,
							}},
						RandomSeed: 1131023344,
						SlotLayout: w3gs.LayoutMelee,
						NumPlayers: 2,
					},
				},
				PlayerInfo: []*w3g.PlayerInfo{
					&w3g.PlayerInfo{
						ID:   3,
						Name: "Серник#26",
					},
				},
				PlayerExtra: []*w3g.PlayerExtra{
					&w3g.PlayerExtra{PlayerExtra: w3gs.PlayerExtra{
						Type: w3gs.PlayerProfile,
						Profiles: []w3gs.PlayerDataProfile{
							w3gs.PlayerDataProfile{
								PlayerID:  3,
								BattleTag: "Серник#2653",
								Clan:      "clan",
								Portrait:  "p052",
								Realm:     w3gs.RealmEurope,
							},
							w3gs.PlayerDataProfile{
								PlayerID:  2,
								BattleTag: "TheBiGsLeeP#2208",
								Clan:      "clan",
								Portrait:  "p029",
								Realm:     w3gs.RealmEurope,
							},
						},
					},
					}},
				Records: []w3g.Record{
					&w3g.TimeSlot{TimeSlot: w3gs.TimeSlot{TimeIncrementMS: 51, Actions: []w3gs.PlayerAction{w3gs.PlayerAction{PlayerID: 2, Data: b64("EhgAAwANAP//////////op92xdiHf8X//////////w==")}}}},
					&w3g.TimeSlot{TimeSlot: w3gs.TimeSlot{TimeIncrementMS: 51}},
					&w3g.TimeSlot{TimeSlot: w3gs.TimeSlot{TimeIncrementMS: 52}},
				},
			},
		},
	}

	for _, f := range files {
		rep, err := w3g.Open("./" + f.file)
		if err != nil {
			t.Fatal("Loading file", err)
		}

		var trunc = *rep
		if len(trunc.PlayerInfo) > 1 {
			trunc.PlayerInfo = trunc.PlayerInfo[1:2]
		}
		if len(trunc.PlayerExtra) > 1 {
			trunc.PlayerExtra = trunc.PlayerExtra[1:2]
		}
		if len(trunc.Slots) > 1 {
			trunc.Slots = trunc.Slots[1:2]
		}
		trunc.Records = trunc.Records[20:23]

		if !reflect.DeepEqual(f.replay, trunc) {
			t.Logf("REF: %+v\n", f.replay)
			t.Logf("OUT: %+v\n", trunc)
			t.Fatal(f.file, "Replay is not deep equal")
		}

		var b protocol.Buffer
		if err := rep.Encode(&b); err != nil {
			t.Fatal("Encode", err)
		}

		rep2, err := w3g.Decode(&b)
		if err != nil {
			t.Fatal("Decode", err)
		}

		if !reflect.DeepEqual(rep, rep2) {
			t.Fatal(f.file, "Replays not deep equal after encode/decode")
		}
	}
}
