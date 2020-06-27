package ecs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

/********************************************************************
created:    2020-01-22
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

type SyncPlayer struct {
	SyncEntity
	name string
	age  int
}

type SyncTitle struct {
	Part
	key int
}

func (title *SyncTitle) OnAdded() {
	var player = (*SyncPlayer)(title.GetEntity())
	fmt.Printf("I am title, this is player=%v\n", player)
}

func (title *SyncTitle) OnRemoved() {
	var player = (*SyncPlayer)(title.GetEntity())
	fmt.Printf("after removed, player=%v\n", player)
}

func (title *SyncTitle) Print(d int) {
	fmt.Printf("===> print: %d\n\n", d)
}

func TestSyncEntity_AddPart(t *testing.T) {
	const title = 1

	var player = &SyncPlayer{}
	player.AddPart(title, &SyncTitle{})

	var part = player.GetPart(title)
	assert.NotNil(t, part)

	player.RemovePart(title)
	var part1 = player.GetPart(title)
	assert.Nil(t, part1)

	player.AddPart(title, &Title{})
	player.ClearParts()
	var part2 = player.GetPart(title)
	assert.Nil(t, part2)
}

func TestSyncEntity_AddParts(t *testing.T) {
	const (
		title = iota
		title1
		title2
		title3
		title4
	)

	var player = &SyncPlayer{}
	for i := 0; i <= title4; i++ {
		player.AddPart(i, &SyncTitle{key: i})
	}

	var removed = player.RemovePart(title2)
	assert.Equal(t, removed.(*SyncTitle).key, title2)
}

func TestSyncEntity_GetParts(t *testing.T) {
	var player = &SyncPlayer{}
	player.AddPart(1, &SyncTitle{key: 1})
	player.AddPart(2, &SyncTitle{key: 2})
	player.AddPart(3, &SyncTitle{key: 3})

	var parts []IPart
	parts = player.GetParts(nil)
	parts = player.GetParts(parts)
	parts = player.GetParts(parts)
	assert.Equal(t, len(parts), 3)
}

func TestSyncEntity_GetEntity(t *testing.T) {
	var player = &SyncPlayer{age: 10, name: "panda"}

	const title = 1
	var part = &SyncTitle{}
	player.AddPart(title, part)

	var p = (*SyncPlayer)(part.GetEntity())
	assert.Equal(t, p.age, 10)
}
