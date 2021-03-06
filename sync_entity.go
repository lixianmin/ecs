package ecs

import (
	"sync"
	"unsafe"
)

/********************************************************************
created:    2020-06-27
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

type SyncEntity struct {
	e Entity
	m sync.RWMutex
}

func (entity *SyncEntity) SetDefaultPart(key int, part IPart) IPart {
	var lastPart = entity.GetPart(key)
	if lastPart == nil && part != nil {
		entity.m.Lock()
		lastPart = entity.e.GetPart(key)
		var isNew = lastPart == nil
		if isNew {
			entity.e.addPartImpl(key, part, unsafe.Pointer(entity))
		}
		entity.m.Unlock()

		// 事件方法必须放到lock外面，否则有可能死锁
		if isNew {
			lastPart = part
			part.OnAdded()
		}
	}

	return lastPart
}

func (entity *SyncEntity) AddPart(key int, part IPart) IPart {
	if nil != part {
		entity.m.Lock()
		entity.e.addPartImpl(key, part, unsafe.Pointer(entity))
		entity.m.Unlock()

		// 事件方法必须放到lock外面，否则有可能死锁
		part.OnAdded()
		return part
	}

	return nil
}

func (entity *SyncEntity) RemovePart(key int) IPart {
	entity.m.Lock()
	var part = entity.e.removePartImpl(key)
	entity.m.Unlock()

	// 事件方法必须放到lock外面，否则有可能死锁
	if part != nil {
		part.OnRemoved()
	}

	return part
}

func (entity *SyncEntity) GetPart(key int) IPart {
	entity.m.RLock()
	var part = entity.e.GetPart(key)
	entity.m.RUnlock()

	return part
}

func (entity *SyncEntity) ClearParts() {
	entity.m.Lock()
	var parts = entity.e.parts
	if parts != nil {
		entity.e.keys = nil
		entity.e.parts = nil
	}
	entity.m.Unlock()

	for _, part := range parts {
		part.OnRemoved()
	}
}

// cache参数是外面提供一个数组，entity下所有的parts都被会填充到这个cache中并返回
//
// 1. 单次调用场景：比如在OnAdded()方法缓存某些parts组件，以避免多次调用时每次获取组件的开销。因为是一次性调用，
// 因此直接传cache=nil就好
//
// 2. 反复调用场景：，循环中调用AddPart()、RemovePart()会影响存储结构，所以遍历快照是更合理的选择。此时应
// 该在entity对象中缓存一个cache数组，并在每次调用的时候作为参数传递。
//
// 问题1： 如何支持gameObject.GetComponents<Type>()这类调用？
// 答：放弃吧，在golang可以考虑遍历并测试类型的方案，但这样最简单也得写一个filter匿名函数来做这件事情，实测会发现
//    跟直接写一个func的代码量相仿。
func (entity *SyncEntity) GetParts(cache []IPart) []IPart {
	entity.m.RLock()
	cache = entity.e.GetParts(cache)
	entity.m.RUnlock()
	return cache
}
