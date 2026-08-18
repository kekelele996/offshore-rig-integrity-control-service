package domain

import "reflect"

type ProviderHandle struct{ Provider PolicyProvider }

func (h ProviderHandle) Available() bool {
	if h.Provider == nil {
		return false
	}
	value := reflect.ValueOf(h.Provider)
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return !value.IsNil()
	default:
		return true
	}
}
