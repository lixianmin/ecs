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

func (entity *Entity) AddPart(key int, part IPart) IPart {
	if nil != part {
		entity.keys = append(entity.keys, key)
		entity.parts = append(entity.parts, part)

		if setEntity, ok := part.(ISetEntity); ok {
			setEntity.SetEntity(unsafe.Pointer(entity))
		}

		entity.sort()
		part.OnAdded()
		return part
	}

	return nil
}

func (entity *Entity) RemovePart(key int) IPart {
	var keys = entity.keys
	var parts = entity.parts

	if keys != nil {
		var index = BinarySearch(keys, key)
		if index >= 0 {
			var part = parts[index]
			part.OnRemoving()

			var count = len(keys)
			for i := index; i < count-1; i++ {
				keys[i] = keys[i+1]
				parts[i] = parts[i+1]
			}

			entity.keys = keys[:count-1]
			entity.parts = parts[:count-1]
			entity.sort()
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

func (entity *Entity) SendMessage(methodName string, args ...interface{}) {
	var parts = entity.parts
	var count = len(parts)
	for i := 0; i < count; i++ {
		var part = parts[i]
		var value = reflect.ValueOf(part)
		if !value.IsValid() {
			continue
		}

		var method = value.MethodByName(methodName)
		if !method.IsValid() {
			continue
		}

		var args1 = fetchArgs(method, args...)
		if args1 != nil {
			method.Call(args1)
		}
	}
}

func fetchArgs(method reflect.Value, args ...interface{}) []reflect.Value {
	var methodType = method.Type()
	var numIn = methodType.NumIn()
	if numIn == 0 {
		return emptyArgs
	}

	// 如果输入参数多于需要的参数，则忽略多余的参数
	var numArgs = len(args)
	if numIn > numArgs {
		return nil
	}

	var args1 = make([]reflect.Value, 0, numIn)
	for i := 0; i < numIn; i++ {
		if methodType.In(i) != reflect.TypeOf(args[i]) {
			return nil
		}

		args1 = append(args1, reflect.ValueOf(args[i]))
	}

	return args1
}

func (entity *Entity) ClearParts() {
	var parts = entity.parts
	if parts != nil {
		for _, part := range parts {
			part.OnRemoving()
		}

		entity.keys = nil
		entity.parts = nil
	}
}

func (entity *Entity) SnapParts(parts []IPart) []IPart {
	var srcParts = entity.parts
	var count = len(srcParts)
	if count > 0 {
		if parts != nil {
			parts = parts[:0]
		}

		for i := 0; i < count; i++ {
			parts = append(parts, srcParts[i])
		}
	}

	return parts
}

func (entity *Entity) sort() {
	var keys = entity.keys
	var parts = entity.parts

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
