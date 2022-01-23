// Author:  Niels A.D.
// Project: gowarcraft3 (https://github.com/nielsAD/gowarcraft3)
// License: Mozilla Public License, v2.0

// w3gsclient is a mocked Warcraft III game client that can be used to add dummy players to games.
package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"

	"github.com/nielsAD/gowarcraft3/network"
	"github.com/nielsAD/gowarcraft3/network/dummy"
	"github.com/nielsAD/gowarcraft3/network/lan"
	"github.com/nielsAD/gowarcraft3/network/peer"
	"github.com/nielsAD/gowarcraft3/protocol/w3gs"
)

var (
	findlan  = flag.Bool("lan", false, "Find a game on LAN")
	gametft  = flag.Bool("tft", true, "Search for TFT or ROC games (only used when searching local)")
	gamevers = flag.Uint("v", uint(w3gs.CurrentGameVersion), "Game version (only used when searching local)")
	entrykey = flag.Uint("e", 0, "Entry key (only used when entering local game)")

	hostcounter = flag.Uint("c", 1, "Host counter")
	dialpeers   = flag.Bool("dial", true, "Dial peers")
	listen      = flag.Int("l", 0, "Listen on port (0 to pick automatically)")

	playername = flag.String("n", "fakeplayer", "Player name")
)

var logOut = log.New(color.Output, "", log.Ltime)
var logErr = log.New(color.Error, "", log.Ltime)
var stdin = bufio.NewReader(os.Stdin)

func main() {
	flag.Parse()

	logOut.SetPrefix(fmt.Sprintf("[%v] ", *playername))
	logErr.SetPrefix(fmt.Sprintf("[%v] ", *playername))

	var addr string
	var hc = uint32(*hostcounter)
	var ek = uint32(*entrykey)

	if *findlan {
		// Search local game for 75 seconds
		var ctx, cancel = context.WithTimeout(context.Background(), 75*time.Second)

		var p = w3gs.ProductTFT
		if !*gametft {
			p = w3gs.ProductROC
		}

		var err error
		addr, hc, ek, err = lan.FindGame(ctx, w3gs.GameVersion{Product: p, Version: uint32(*gamevers)})
		cancel()

		if err != nil {
			logErr.Fatal("Could not find local game: ", err)
		}
	} else {
		addr = strings.Join(flag.Args(), ":")
		if addr == "" {
			addr = "127.0.0.1:6112"
		}
	}

	logOut.Println(color.MagentaString("Joining lobby at %s (ID: %d, key: %d)", addr, hc, ek))
	d, err := dummy.Join(addr, *playername, hc, ek, *listen, w3gs.Encoding{GameVersion: uint32(*gamevers)})
	if err != nil {
		logErr.Fatal("Join error: ", err)
	}

	d.DialPeers = *dialpeers
	logOut.Println(color.MagentaString("Joined lobby with (ID: %d)", d.PlayerInfo.PlayerID))

	d.On(&network.AsyncError{}, func(ev *network.Event) {
		var err = ev.Arg.(*network.AsyncError)
		logErr.Println(color.RedString("[ERROR] %s", err.Error()))
	})
	d.On(&peer.Registered{}, func(ev *network.Event) {
		var reg = ev.Arg.(*peer.Registered)
		var pi = &reg.Peer.PlayerInfo

		logOut.Println(color.YellowString("%s has joined the game (ID: %d)", pi.PlayerName, pi.PlayerID))

		reg.Peer.On(&network.AsyncError{}, func(ev *network.Event) {
			var err = ev.Arg.(*network.AsyncError)
			logErr.Println(color.RedString("[ERROR] [PEER%d] %s", pi.PlayerID, err.Error()))
		})
	})
	d.On(&peer.Deregistered{}, func(ev *network.Event) {
		var reg = ev.Arg.(*peer.Deregistered)
		logOut.Println(color.YellowString("%s has left the game (ID: %d)", reg.Peer.PlayerInfo.PlayerName, reg.Peer.PlayerInfo.PlayerID))
	})
	d.On(&peer.Connected{}, func(ev *network.Event) {
		var e = ev.Arg.(*peer.Connected)
		if e.Dial {
			logOut.Println(color.MagentaString("Established peer connection to %s (ID: %d)", e.Peer.PlayerInfo.PlayerName, e.Peer.PlayerInfo.PlayerID))
		} else {
			logOut.Println(color.MagentaString("Accepted peer connection from %s (ID: %d)", e.Peer.PlayerInfo.PlayerName, e.Peer.PlayerInfo.PlayerID))
		}
	})
	d.On(&peer.Disconnected{}, func(ev *network.Event) {
		var e = ev.Arg.(*peer.Disconnected)
		logOut.Println(color.MagentaString("Peer connection to %s (ID: %d) closed", e.Peer.PlayerInfo.PlayerName, e.Peer.PlayerInfo.PlayerID))
	})

	d.On(&w3gs.PlayerKicked{}, func(ev *network.Event) {
		logOut.Println(color.MagentaString("Kicked from lobby"))
	})
	d.On(&w3gs.CountDownStart{}, func(ev *network.Event) {
		logOut.Println(color.CyanString("Countdown started"))
	})
	d.On(&w3gs.CountDownEnd{}, func(ev *network.Event) {
		logOut.Println(color.CyanString("Countdown ended, loading game"))
	})

	d.On(&w3gs.StartLag{}, func(ev *network.Event) {
		var lag = ev.Arg.(*w3gs.StartLag)

		var laggers []string
		for _, l := range lag.Players {
			var name = ""
			if l.PlayerID == d.PlayerInfo.PlayerID {
				name = d.PlayerInfo.PlayerName
			} else {
				if peer := d.Peer(l.PlayerID); peer != nil {
					name = peer.PlayerInfo.PlayerName
				}
			}
			if name == "" {
				continue
			}
			laggers = append(laggers, name)
		}

		logOut.Println(color.CyanString("Lag: %v", laggers))
	})
	d.On(&w3gs.StopLag{}, func(ev *network.Event) {
		var lag = ev.Arg.(*w3gs.StopLag)
		var peer = d.Peer(lag.PlayerID)
		if peer == nil {
			return
		}

		logOut.Println(color.CyanString("%s (ID: %d) stopped lagging", peer.PlayerInfo.PlayerName, peer.PlayerInfo.PlayerID))
	})

	d.On(&dummy.Say{}, func(ev *network.Event) {
		var say = ev.Arg.(*dummy.Say)
		logOut.Printf("[CHAT] %s (ID: %d): %s\n", d.PlayerInfo.PlayerName, d.PlayerInfo.PlayerID, say.Content)
	})
	d.On(&dummy.Chat{}, func(ev *network.Event) {
		var chat = ev.Arg.(*dummy.Chat)
		if chat.Content == "" || chat.Sender == nil {
			return
		}

		logOut.Printf("[CHAT] %s (ID: %d): %s\n", chat.Sender.PlayerName, chat.Sender.PlayerID, chat.Content)
		if chat.Sender.PlayerID != 1 || chat.Content[0] != '.' {
			return
		}

		var cmd = strings.Fields(chat.Content)
		switch strings.ToLower(cmd[0]) {
		case ".say":
			d.Say(strings.Join(cmd[1:], " "))
		case ".leave":
			d.Leave(w3gs.LeaveLost)
		case ".race":
			if len(cmd) != 2 {
				d.Say("use like: .race [str]")
				break
			}

			switch strings.ToLower(cmd[1]) {
			case "h", "hu", "hum", "human":
				d.ChangeRace(w3gs.RaceHuman)
			case "o", "orc":
				d.ChangeRace(w3gs.RaceOrc)
			case "u", "ud", "und", "undead":
				d.ChangeRace(w3gs.RaceUndead)
			case "n", "ne", "elf", "nightelf":
				d.ChangeRace(w3gs.RaceNightElf)
			case "r", "rnd", "rdm", "random":
				d.ChangeRace(w3gs.RaceRandom)
			default:
				d.Say("Invalid race")
			}
		case ".team":
			if len(cmd) != 2 {
				d.Say("use like: .team [int]")
				break
			}
			if t, err := strconv.ParseUint(cmd[1], 0, 8); err == nil && t >= 1 {
				d.ChangeTeam(uint8(t - 1))
			}
		case ".color":
			if len(cmd) != 2 {
				d.Say("use like: .color [int]")
				break
			}
			if c, err := strconv.ParseUint(cmd[1], 0, 8); err == nil && c >= 1 {
				d.ChangeColor(uint8(c - 1))
			}
		case ".handicap":
			if len(cmd) != 2 {
				d.Say("use like: .handicap [int]")
				break
			}
			if h, err := strconv.ParseUint(cmd[1], 0, 8); err == nil {
				d.ChangeHandicap(uint8(h))
			}
		}
	})

	go func() {
		for {
			line, err := stdin.ReadString('\n')
			if err != nil {
				d.Close()
				break
			}

			if err := d.Say(strings.TrimRight(line, "\r\n")); err != nil {
				logErr.Println(color.RedString("[ERROR] %s", err.Error()))
			}
		}
	}()

	if err := d.Run(); err != nil && !network.IsCloseError(err) {
		logErr.Println(color.RedString("[ERROR] %s", err.Error()))
	}
}
