// Author:  Niels A.D.
// Project: gowarcraft3 (https://github.com/nielsAD/gowarcraft3)
// License: Mozilla Public License, v2.0
package capi_test

import (
	"reflect"
	"testing"

	"github.com/nielsAD/gowarcraft3/protocol"
	"github.com/nielsAD/gowarcraft3/protocol/capi"
)

func TestPackets(t *testing.T) {
	var types = map[string]interface{}{
		"Testapi.RandomCommand" + capi.CmdRequestSuffix: &map[string]interface{}{"key": "value"},
		capi.CmdAuthenticate + capi.CmdResponseSuffix:   &capi.Response{},
		capi.CmdAuthenticate + capi.CmdRequestSuffix: &capi.Authenticate{
			APIKey: "[API KEY]",
		},
		capi.CmdConnect + capi.CmdRequestSuffix:    &capi.Connect{},
		capi.CmdDisconnect + capi.CmdRequestSuffix: &capi.Disconnect{},
		capi.CmdSendMessage + capi.CmdRequestSuffix: &capi.SendMessage{
			Message: "[MESSAGE]",
		},
		capi.CmdSendEmote + capi.CmdRequestSuffix: &capi.SendEmote{
			Message: "[EMOTE MESSAGE]",
		},
		capi.CmdSendWhisper + capi.CmdRequestSuffix: &capi.SendWhisper{
			Message: "[MESSAGE]",
			UserID:  "[USER ID]",
		},
		capi.CmdKickUser + capi.CmdRequestSuffix: &capi.KickUser{
			UserID: "[USER ID]",
		},
		capi.CmdBanUser + capi.CmdRequestSuffix: &capi.BanUser{
			UserID: "[USER ID]",
		},
		capi.CmdUnbanUser + capi.CmdRequestSuffix: &capi.UnbanUser{
			Username: "[TOON NAME]",
		},
		capi.CmdSetModerator + capi.CmdRequestSuffix: &capi.SetModerator{
			UserID: "[USER ID]",
		},
		capi.CmdConnectEvent + capi.CmdRequestSuffix: &capi.ConnectEvent{
			Channel: "Op Lodle",
		},
		capi.CmdDisconnectEvent + capi.CmdRequestSuffix: &capi.DisconnectEvent{},
		capi.CmdMessageEvent + capi.CmdRequestSuffix: &capi.MessageEvent{
			UserID:  "[USER ID]",
			Message: "[MESSAGE]",
			Type:    capi.MessageServerInfo,
		},
		capi.CmdUserUpdateEvent + capi.CmdRequestSuffix: &capi.UserUpdateEvent{
			UserID:   "[USER ID]",
			Username: "[TOON NAME]",
			Flags:    capi.UserFlagModerator | capi.UserFlagMuteWhisper,
			Attributes: capi.UserAttributes{
				ProgramID: "W3XP",
				Rate:      "1",
				Rank:      "2",
				Wins:      "3",
			},
		},
		capi.CmdUserLeaveEvent + capi.CmdRequestSuffix: &capi.UserLeaveEvent{},
	}

	for cmd, payl := range types {
		var buf = protocol.Buffer{Bytes: make([]byte, 0, 2048)}

		var pkt = &capi.Packet{
			Command:   cmd,
			RequestID: 123,
			Status:    &capi.Success,
			Payload:   payl,
		}

		if err := capi.SerializePacket(&buf, pkt); err != nil {
			t.Log(reflect.TypeOf(payl))
			t.Fatal(err)
		}

		var pkt2, err = capi.DeserializePacket(&buf)
		if err != nil {
			t.Log(reflect.TypeOf(payl))
			t.Fatal(err)
		}
		if buf.Size() > 0 {
			t.Fatalf("DeserializePacket size mismatch for %v", reflect.TypeOf(payl))
		}
		if reflect.TypeOf(pkt2.Payload) != reflect.TypeOf(payl) {
			t.Fatalf("DeserializePacket type mismatch %v != %v", reflect.TypeOf(pkt2.Payload), reflect.TypeOf(payl))
		}
		if !reflect.DeepEqual(pkt, pkt2) {
			t.Logf("I: %+v", pkt)
			t.Logf("O: %+v", pkt2)
			t.Errorf("DeserializePacket value mismatch for %v", reflect.TypeOf(payl))
		}
	}
}

var testPkt = capi.Packet{
	Command:   capi.CmdMessageEvent + capi.CmdRequestSuffix,
	RequestID: 123,
	Status:    &capi.Success,
	Payload: &capi.MessageEvent{
		UserID:  "[USER ID]",
		Message: "[MESSAGE]",
		Type:    capi.MessageServerInfo,
	},
}

func BenchmarkSerializePacket(b *testing.B) {
	var w = &protocol.Buffer{}

	capi.SerializePacket(w, &testPkt)

	b.SetBytes(int64(w.Size()))
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		w.Truncate()
		capi.SerializePacket(w, &testPkt)
	}
}

func BenchmarkDeserializePacket(b *testing.B) {
	var pbuf = protocol.Buffer{Bytes: make([]byte, 0, 2048)}
	capi.SerializePacket(&pbuf, &testPkt)

	b.SetBytes(int64(pbuf.Size()))
	b.ResetTimer()

	var r = &protocol.Buffer{}
	for n := 0; n < b.N; n++ {
		r.Bytes = pbuf.Bytes
		capi.DeserializePacket(r)
	}
}
