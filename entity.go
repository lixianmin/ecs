package ecs

import (
	"reflect"
	"unsafe"
)

/********************************************************************
created:    2020-01-22
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

var emptyArgs = make([]reflect.Value, 0)

type Entity struct {
	keys  []int
	parts []IPart
}

func (entity *Entity) SetDefaultPart(key int, part IPart) IPart {
	var lastPart = entity.GetPart(key)
	if lastPart == nil && part != nil {
		lastPart = entity.AddPart(key, part)
	}

	return lastPart
}

func (entity *Entity) AddPart(key int, part IPart) IPart {
	if nil != part {
		entity.addPartImpl(key, part, unsafe.Pointer(entity))
		part.OnAdded()
		return part
	}

	return nil
}

func (entity *Entity) addPartImpl(key int, part IPart, pEntity unsafe.Pointer) {
	entity.keys = append(entity.keys, key)
	entity.parts = append(entity.parts, part)

	// 这里不能使用this指针这个entity，因为在SyncEntity调用时会被切断
	if setEntity, ok := part.(ISetEntity); ok {
		setEntity.SetEntity(pEntity)
	}

	sortKeysParts(entity.keys, entity.parts)
}

func (entity *Entity) RemovePart(key int) IPart {
	var part = entity.removePartImpl(key)
	if part != nil {
		part.OnRemoved()
	}

	return part
}

func (entity *Entity) removePartImpl(key int) IPart {
	var keys = entity.keys
	var parts = entity.parts

	if keys != nil {
		var index = BinarySearch(keys, key)
		if index >= 0 {
			var part = parts[index]

			var count = len(keys)
			for i := index; i < count-1; i++ {
				keys[i] = keys[i+1]
				parts[i] = parts[i+1]
			}

			entity.keys = keys[:count-1]
			entity.parts = parts[:count-1]
			sortKeysParts(entity.keys, entity.parts)

			return part
		}
	}

	return nil
}

func (entity *Entity) GetPart(key int) IPart {
	var keys = entity.keys
	if keys != nil {
		var index = BinarySearch(keys, key)
		if index >= 0 {
			var part = entity.parts[index]
			return part
		}
	}

	return nil
}

func (entity *Entity) ClearParts() {
	var parts = entity.parts
	if parts != nil {
		entity.keys = nil
		entity.parts = nil

		for _, part := range parts {
			part.OnRemoved()
		}
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
func (entity *Entity) GetParts(cache []IPart) []IPart {
	var parts = entity.parts
	var count = len(parts)
	if count > 0 {
		if cache != nil {
			cache = cache[:0]
		}

		for i := 0; i < count; i++ {
			cache = append(cache, parts[i])
		}
	}

	return cache
}

func sortKeysParts(keys []int, parts []IPart) {
	var count = len(keys)
	if keys == nil || count != len(parts) {
		return
	}

	for i := 1; i < count; i++ {
		for j := i; j > 0 && keys[j] < keys[j-1]; j-- {
			keys[j], keys[j-1] = keys[j-1], keys[j]
			parts[j], parts[j-1] = parts[j-1], parts[j]
		}
	}
}
