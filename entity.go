package ecs

import (
	"unsafe"
)

/********************************************************************
created:    2020-01-22
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

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
