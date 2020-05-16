package ecs

import (
	"fmt"
	"testing"
)

/********************************************************************
created:    2020-01-22
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

type Player struct {
	Entity
	name string
	age  int
}

type Title struct {
	Part
	key int
}

func (title *Title) OnAdded() {
	var player = (*Player)(title.GetEntity())
	fmt.Printf("I am title, this is player=%v\n", player)
}

func (title *Title) OnRemoving() {
	var player = (*Player)(title.GetEntity())
	fmt.Printf("before removed, player=%v\n", player)
}

func TestEntity_AddPart(t *testing.T) {
	const title = 1

	var player = &Player{}
	player.AddPart(title, &Title{})

	var part = player.GetPart(title)
	if part == nil {
		t.Errorf("part should be got")
	}
	fmt.Println(part)

	player.RemovePart(title)
	var part1 = player.GetPart(title)
	if part1 != nil {
		t.Errorf("part1 should be removed")
	}

	player.AddPart(title, &Title{})
	player.ClearParts()
	var part2 = player.GetPart(title)
	if part2 != nil {
		t.Errorf("player should has no parts now")
	}
}

func TestEntity_AddParts(t *testing.T) {
	const (
		title = iota
		title1
		title2
		title3
		title4
	)

	var player = &Player{}
	for i := 0; i <= title4; i++ {
		player.AddPart(i, &Title{key: i})
	}

	var removed = player.RemovePart(title2)
	if removed.(*Title).key != title2 {
		t.Errorf("remove error")
	}
}
