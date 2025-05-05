package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"reflect"
)

/** 对 widget.List 进一步封装常用属性 */

type ListView struct {
	GetData func() interface{}
	*widget.List
}

func NewListView(
	setData func() interface{},
	setItemView func() fyne.CanvasObject,
	OnUpdate func(id widget.ListItemID, object fyne.CanvasObject, lv *ListView)) (*ListView, error) {

	lv := &ListView{}
	lv.GetData = setData
	if _, e := lv.toSlice(setData()); e != nil {
		return nil, e
	}

	mList := widget.NewList(func() int {
		data, _ := lv.toSlice(lv.GetData())
		return len(data)
	}, func() fyne.CanvasObject {
		return setItemView()
	}, func(id widget.ListItemID, object fyne.CanvasObject) {})
	mList.UpdateItem = func(id widget.ListItemID, item fyne.CanvasObject) {
		if OnUpdate != nil {
			OnUpdate(id, item, lv)
		}
	}
	lv.List = mList
	return lv, nil
}

//func (l *ListView) DeleteItem(id widget.ListItemID) {
//	newListData := []interface{}{}
//	data, _ := l.toSlice(l.GetData())
//	for i := 0; i < len(data); i++ {
//		if i == id {
//			continue
//		}
//		newListData = append(newListData, data[i])
//	}
//	if len(newListData) > 0 || (len(data) == 1 && id == 0) { // last delete
//		data = newListData
//		l.List.Refresh()
//	}
//}

func (l *ListView) DisableItemClick() {
	l.List.OnSelected = func(id widget.ListItemID) {
		l.List.Unselect(id)
	}
}

func (l *ListView) NotifyDataChange() {
	l.List.Refresh()
}

func (l *ListView) toSlice(value interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return nil, fmt.Errorf("invalid data")
	}
	result := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		result[i] = v.Index(i).Interface()
	}
	return result, nil
}
