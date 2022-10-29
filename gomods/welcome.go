package snats

import (
	"os"
	"net/url"
	"log"
	"fmt"
	"errors"
	"encoding/json"
	"fyne.io/fyne/v2"
//	"fyne.io/fyne/v2/canvas"
//	"fyne.io/fyne/v2/cmd/fyne_demo/data"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)
var Password string
var Server string
var Caroot string
var Queue string
var Queuepassword string

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}

func welcomeScreen(_ fyne.Window) fyne.CanvasObject {
//	logo := canvas.NewImageFromResource(data.FyneScene)
//	logo.FillMode = canvas.ImageFillContain
//	if fyne.CurrentDevice().IsMobile() {
//		logo.SetMinSize(fyne.NewSize(96, 96))
//	} else {
//		logo.SetMinSize(fyne.NewSize(128, 128))
//	}
	enckey := widget.NewEntry()
	enckey.SetPlaceHolder("Enter Password For Config Encryption")
	enckey.SetText(Password)
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("New Horizons 3000 Secure Communications", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		enckey,
		widget.NewButton("Save", func() {	
			log.Println(enckey.Text)
			Password = enckey.Text
			loadfile()
		}),	
//		logo,
		container.NewHBox(
			widget.NewHyperlink("Documentation", parseURL("https://newhorizons3000,org/")),

		),
		widget.NewLabel(""), // balance the header on the tutorial screen we leave blank on this content
	))

	
}
func loadfile() {
			 _, err := os.Stat("servers.json")
    		if errors.Is(err, os.ErrNotExist) {
        		fmt.Println("file does not exist")
				file, err := os.Create("server.json")
				if err != nil {
        			log.Fatal(err)
    			}
				enc := json.NewEncoder(file)
				str :=  map[string]string {"password":  Password, "server": "NoneEntered", "caroot": "NoCertificate", "queue": "NoQueue", "queuepassword": "NoQueuePassword" }
				enc.Encode((str))

    		} else {
    		    fmt.Println("file exists")
    		}

	}
