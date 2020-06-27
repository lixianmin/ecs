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

func (title *Title) Print(d int) {
	fmt.Printf("===> print: %d\n\n", d)
}

func TestEntity_AddPart(t *testing.T) {
	const title = 1

	var player = &Player{}
	player.AddPart(title, &Title{})

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
	assert.Equal(t, removed.(*Title).key, title2)
}

//func TestEntity_SendMessage(t *testing.T) {
//	var player = &Player{}
//	player.AddPart(1, &Title{})
//
//	player.SendMessage("Print")             // 参数个数不对
//	player.SendMessage("Print", 1)          // 正常调用
//	player.SendMessage("Print", "2")        // 参数类型不对
//	player.SendMessage("Print", 3, 123)     // 参数个数不对
//	player.SendMessage("InvalidMethod", 10) // 不存在的方法
//}

func TestEntity_GetParts(t *testing.T) {
	var player = &Player{}
	player.AddPart(1, &Title{key: 1})
	player.AddPart(2, &Title{key: 2})
	player.AddPart(3, &Title{key: 3})

	var parts []IPart
	parts = player.GetParts(nil)
	parts = player.GetParts(parts)
	parts = player.GetParts(parts)
	assert.Equal(t, len(parts), 3)
}
