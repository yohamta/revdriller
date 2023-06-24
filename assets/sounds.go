package assets

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

var (
	audioContext *audio.Context
)

type BGM int

const (
	BGMMain BGM = iota
)

type SE int

const (
	SEAttack SE = iota
	SEBreak
)

var (
	bgms = map[BGM]*audio.Player{}
	ses  = map[SE]*audio.Player{}
)

func PlayBGM(bgm BGM) {
	bgms[bgm].Play()
}

func PlaySE(se SE) {
	ses[se].Rewind()
	ses[se].Play()
}

const sampleRate = 22050

func loadSounds() {
	audioContext = audio.NewContext(sampleRate)

	bgms[BGMMain] = loadMP3(audioContext, "sounds/GJ_ChikenCrazy.mp3")
	ses[SEAttack] = loadWav(audioContext, "sounds/se_attack.wav")
	ses[SEBreak] = loadWav(audioContext, "sounds/se_break.wav")
}

func loadMP3(c *audio.Context, name string) *audio.Player {
	// TODO: migrate to DecodeWithSampleRate
	decoded, err := mp3.Decode(audioContext, bytes.NewReader(mustRead(name)))
	if err != nil {
		panic(fmt.Sprintf("failed to decode wav: %v", err))
	}
	infiniteStream := audio.NewInfiniteLoop(decoded, decoded.Length())
	player, _ := audio.NewPlayer(audioContext, infiniteStream)
	return player
}

func loadWav(c *audio.Context, name string) *audio.Player {
	// TODO: migrate to DecodeWithSampleRate
	decoded, err := wav.Decode(c, bytes.NewReader(mustRead(name)))
	if err != nil {
		panic(fmt.Sprintf("failed to decode wav: %v", err))
	}
	b, _ := ioutil.ReadAll(decoded)
	player := audio.NewPlayerFromBytes(audioContext, b)
	return player
}
