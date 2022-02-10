package modules

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"image/jpeg"
	"os"
)

var (
	positionsSlider  *widgets.QSlider
	sliderPos        int
	displayArea      *widgets.QWidget
	scene            *widgets.QGraphicsScene
	view             *widgets.QGraphicsView
	item             *widgets.QGraphicsPixmapItem
	mainApp          *widgets.QApplication
	snapshotFileName string
	qImg             = new(gui.QImage)
	cameraFeedLabel  *widgets.QLabel
	opt              = jpeg.Options{90} // for the best color, use png instead of jpeg
	stopCamera       = false            // to prevent segmentation fault
	snapPhoto        = false
	width, height    int
)

func InitWindow() {
	mainApp = widgets.NewQApplication(len(os.Args), os.Args)

	mainGUI().SetWindowTitle("Golang Qt OpenCV GUI application example")
	mainGUI().SetWindowState(core.Qt__WindowMaximized) // maximized on start
	mainGUI().Show()

	widgets.QApplication_Exec()
}
func setPosition(position int) {
	sliderPos = position
}

func mainGUI() *widgets.QWidget {
	scene = widgets.NewQGraphicsScene(nil)
	view = widgets.NewQGraphicsView(nil)
	displayArea = widgets.NewQWidget(nil, 0)
	cameraFeedLabel = widgets.NewQLabel(nil, core.Qt__Widget)

	view.SetScene(scene)

	// create a slider
	positionsSlider = widgets.NewQSlider2(core.Qt__Horizontal, nil)
	positionsSlider.SetRange(0, 100)
	positionsSlider.SetTickInterval(10)
	positionsSlider.SetValue(50)
	positionsSlider.ShowMaximized()

	positionsSlider.ConnectSliderMoved(setPosition)

	//create a button and connect the clicked signal
	var snapButton = widgets.NewQPushButton2("Take a snapshot", nil)
	snapButton.ConnectClicked(func(flag bool) {

		snapPhoto = true

	})

	//create a button and connect the clicked signal
	var quitButton = widgets.NewQPushButton2("Quit", nil)
	quitButton.ConnectClicked(func(flag bool) {

		stopCamera = true
		mainApp.Quit()
	})

	var layout = widgets.NewQVBoxLayout()

	layout.AddWidget(view, 0, core.Qt__AlignCenter)
	layout.AddWidget(positionsSlider, 0, core.Qt__AlignCenter)

	layout.AddWidget(snapButton, 0, core.Qt__AlignCenter)
	layout.AddWidget(quitButton, 0, core.Qt__AlignRight)

	displayArea.SetLayout(layout)

	return displayArea
}
