package ecs

import "unsafe"

/********************************************************************
created:    2020-01-23
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

type IPart interface {
	OnAdded()
	OnRemoving()
}

type ISetEntity interface {
	SetEntity(entity unsafe.Pointer)
}

type IGetEntity interface {
	GetEntity() unsafe.Pointer
}
