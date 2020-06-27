package ecs

import (
	"sync/atomic"
	"unsafe"
)

/********************************************************************
created:    2020-01-22
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

type Part struct {
	entity unsafe.Pointer
}

func (part *Part) OnAdded() {

}

func (part *Part) OnRemoved() {

}

func (part *Part) SetEntity(entity unsafe.Pointer) {
	atomic.StorePointer(&part.entity, entity)
}

func (part *Part) GetEntity() unsafe.Pointer {
	return atomic.LoadPointer(&part.entity)
}
