package main

import (
	"os"

	_ "github.com/akiyosi/qt/interop"

	"github.com/akiyosi/qt/core"
	_ "github.com/akiyosi/qt/gui"
	_ "github.com/akiyosi/qt/multimedia"
	_ "github.com/akiyosi/qt/quick"
	_ "github.com/akiyosi/qt/quickcontrols2"
	"github.com/akiyosi/qt/widgets"
)

func main() {

	// enable high dpi scaling
	// useful for devices with high pixel density displays
	// such as smartphones, retina displays, ...
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	widgets.NewQApplication(len(os.Args), os.Args)

	widgets.QApplication_Exec()
}
