package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/helper/myfont"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"math/rand"
)

var (
	sizeW float32 = 640 * 3.4
	sizeH float32 = 480 * 4.5
)

type ListDataObj struct {
	Id   string `json:"id"`
	Col1 string
	Col2 string
	Col3 string
}

func main() {
	myApp := app.NewWithID("io.fyne.demo")
	myApp.SetIcon(theme.FyneLogo())
	myApp.Settings().SetTheme(myfont.NewLightTheme())

	mainWindow := myApp.NewWindow("ListView Demo")

	var listData = []ListDataObj{
		{
			Id:   "1",
			Col1: "22",
			Col2: "33",
			Col3: "44",
		},
		{
			Id:   "2",
			Col1: "a22",
			Col2: "a33",
			Col3: "a44",
		},
		{
			Id:   "3",
			Col1: "b22",
			Col2: "b33",
			Col3: "b44",
		},
	}

	listView, err := NewListView(func() interface{} {
		return listData
	}, func() fyne.CanvasObject {
		return container.NewHBox(
			widget.NewLabel("Col1"),
			widget.NewLabel("Col2"),
			widget.NewLabel("Col3"),
			widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {}))
	}, func(id widget.ListItemID, object fyne.CanvasObject, lv *ListView) {
		fmt.Println("id:", id, listData[id].Col1)
		hBox := object.(*fyne.Container)
		(hBox.Objects[0]).(*widget.Label).SetText(fmt.Sprintf("%d", id))
		(hBox.Objects[1]).(*widget.Label).SetText(listData[id].Col1)
		(hBox.Objects[2]).(*widget.Label).SetText(listData[id].Col2)
		(hBox.Objects[3]).(*widget.Button).OnTapped = func() {
			fmt.Println("del:", id)
			newListData := []ListDataObj{}
			for i := 0; i < len(listData); i++ {
				if i == id {
					continue
				}
				newListData = append(newListData, listData[i])
			}
			if len(newListData) > 0 || (len(listData) == 1 && id == 0) {
				listData = newListData
				lv.NotifyDataChange()
			}
		}
	})
	if err != nil {
		panic(err)
	}

	vbox := container.NewBorder(widget.NewButton("addOne", func() {
		listData = append(listData, ListDataObj{
			Id:   fmt.Sprintf("%d", rand.Int63n(1000)+3),
			Col1: "v",
			Col2: "v",
			Col3: "v",
		})
		listView.NotifyDataChange()
	}), nil, nil, nil, listView)
	listView.DisableItemClick()
	mainWindow.SetContent(vbox)
	mainWindow.Resize(fyne.NewSize(sizeW/3, sizeH/4))
	mainWindow.ShowAndRun()
}
