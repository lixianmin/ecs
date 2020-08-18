package ecs

/********************************************************************
created:    2020-08-18
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

// 因为golang的反射很慢，通过SendMessage()调用方法并不是推荐的方式。
// 另外，为了防止SendMessage()的过程中parts有变化，还需要取一个快照出来，有些慢
//func (entity *Entity) SendMessage(methodName string, args ...interface{}) {
//	var parts = entity.GetParts(nil)
//	sendMessage(parts, methodName, args...)
//}
//
//func sendMessage(parts []IPart, methodName string, args ...interface{}) {
//	var count = len(parts)
//	for i := 0; i < count; i++ {
//		var part = parts[i]
//		var value = reflect.ValueOf(part)
//		if !value.IsValid() {
//			continue
//		}
//
//		var method = value.MethodByName(methodName)
//		if !method.IsValid() {
//			continue
//		}
//
//		var args1 = fetchArgs(method, args...)
//		if args1 != nil {
//			method.Call(args1)
//		}
//	}
//}
//
//func fetchArgs(method reflect.Value, args ...interface{}) []reflect.Value {
//	var methodType = method.Type()
//	var numIn = methodType.NumIn()
//	if numIn == 0 {
//		return emptyArgs
//	}
//
//	// 如果输入参数多于需要的参数，则忽略多余的参数
//	var numArgs = len(args)
//	if numIn > numArgs {
//		return nil
//	}
//
//	var args1 = make([]reflect.Value, 0, numIn)
//	for i := 0; i < numIn; i++ {
//		if methodType.In(i) != reflect.TypeOf(args[i]) {
//			return nil
//		}
//
//		args1 = append(args1, reflect.ValueOf(args[i]))
//	}
//
//	return args1
//}
