package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"io"
	"log"
	"net/http"
	"strings"
)

var (
	AcerLogoUri   = "https://i.postimg.cc/XJdzf0Zv/png-clipart-laptop-acer-aspire-predator-computer-projector-acer-text-computer-prev-ui.png"
	AsusLogoUri   = "https://i.postimg.cc/rsQZPPn5/asus-logo-prev-ui.png"
	HpLogoUri     = "https://i.postimg.cc/7YtSQkVK/466-hp-prev-ui.png"
	HpBellLogoUri = "https://i.postimg.cc/cC48X98q/packard-bell7897-logowik-com-prev-ui.png"
	SamLogoUri    = "https://i0.wp.com/www.alicjaszczypiorska.pl/wp-content/uploads/2018/02/samsung-logo-vector.png"
	SonyLogoUrl   = "https://cdn.freebiesupply.com/images/large/2x/sony-logo-png-transparent.png"
	Mkey          = "https://i.postimg.cc/MGh2sfsp/mkey-prev-ui.png"
)

var MasterEndpoint = "https://gtavisfreebro.pythonanywhere.com?code="

func Acer8DigitTab(mainWindow fyne.Window) *fyne.Container {
	MainContainer := container.NewVBox()
	//mainWindow.SetTitle("Acer 8 Digit Bios")
	AcerLogo := canvas.NewImageFromResource(icns.ImageFromUrl(AcerLogoUri))
	AcerLogo.SetMinSize(fyne.NewSize(260, 160))
	MainContainer.Add(AcerLogo)
	MainContainer.Add(widget.NewLabel("Acer 8-Digit BIOS Unlock Tool"))
	MainContainer.Add(widget.NewLabel("This is the Acer 8-Digit BIOS Unlock Tool"))
	MainContainer.Add(widget.NewLabel("Insert code from bios here, should be something like: 45454545"))
	codeEntry := widget.NewEntry()
	codeEntry.PlaceHolder = "Only 8-Digit Codes"
	MainContainer.Add(codeEntry)
	MainContainer.Add(widget.NewButton("Unlock", func() {
		code := codeEntry.Text
		if code != "" && len(code) == 8 {
			result := SendUnlockRequest(Acer8DigitEndPoint, code)
			dialog.ShowInformation("Result", parseCode(result), mainWindow)
		} else {
			dialog.ShowInformation("Error", "Please enter a valid code", mainWindow)
		}
	}))
	return MainContainer
}
func Sony7DigitTab(mainWindow fyne.Window) *fyne.Container {
	MainContainer := container.NewVBox()
	//mainWindow.SetTitle("Sony 7 Digit Bios")
	SonyLogo := canvas.NewImageFromResource(icns.ImageFromUrl(SonyLogoUrl))
	SonyLogo.SetMinSize(fyne.NewSize(260, 160))
	MainContainer.Add(SonyLogo)
	MainContainer.Add(widget.NewLabel("Sony 7-Digit BIOS Unlock Tool"))
	MainContainer.Add(widget.NewLabel("This is the Sony 7-Digit BIOS Unlock Tool"))
	MainContainer.Add(widget.NewLabel("Insert code from bios here, should be something like: 1234567"))
	codeEntry := widget.NewEntry()
	codeEntry.PlaceHolder = "Only 7-Digit Codes"
	MainContainer.Add(codeEntry)
	MainContainer.Add(widget.NewButton("Unlock", func() {
		code := codeEntry.Text
		if code != "" && len(code) == 7 {
			result := SendUnlockRequest(Sony7DigitEndPoint, code)
			dialog.ShowInformation("Result", parseCode(result), mainWindow)
		} else {
			dialog.ShowInformation("Error", "Please enter a valid code", mainWindow)
		}

	}))
	return MainContainer
}

func HpCo5DigitTab(mainWindow fyne.Window) *fyne.Container {
	MainContainer := container.NewVBox()
	//mainWindow.SetTitle("HP Copaq 5 Digit Bios")
	HpLogo := canvas.NewImageFromResource(icns.ImageFromUrl(HpLogoUri))
	HpLogo.SetMinSize(fyne.NewSize(256, 256))
	MainContainer.Add(HpLogo)
	MainContainer.Add(widget.NewLabel("HP CO 5-Digit BIOS Unlock Tool"))
	MainContainer.Add(widget.NewLabel("This is the HP CO 5-Digit BIOS Unlock Tool"))
	MainContainer.Add(widget.NewLabel("Insert code from bios here, should be something like: 12345"))
	codeEntry := widget.NewEntry()
	codeEntry.PlaceHolder = "Only 5-Digit Codes"
	MainContainer.Add(codeEntry)
	MainContainer.Add(widget.NewButton("Unlock", func() {
		code := codeEntry.Text
		if code != "" && len(code) == 5 {
			result := SendUnlockRequest(HpCo5DigitEndPoint, code)
			dialog.ShowInformation("Result", parseCode(result), mainWindow)
		} else {
			dialog.ShowInformation("Error", "Please enter a valid code", mainWindow)
		}

	}))
	return MainContainer
}

func Bell8DigitTab(mainWindow fyne.Window) *fyne.Container {
	MainContainer := container.NewVBox()
	//mainWindow.SetTitle("HP Bell 8 Digit Bios")
	HpBellLogo := canvas.NewImageFromResource(icns.ImageFromUrl(HpBellLogoUri))
	HpBellLogo.SetMinSize(fyne.NewSize(260, 260))
	MainContainer.Add(HpBellLogo)
	MainContainer.Add(widget.NewLabel("Bell 8-Digit BIOS Unlock Tool"))
	MainContainer.Add(widget.NewLabel("This is the Bell 8-Digit BIOS Unlock Tool"))
	MainContainer.Add(widget.NewLabel("Insert code from bios here, should be something like: 12345678"))
	codeEntry := widget.NewEntry()
	codeEntry.PlaceHolder = "Only 8-Digit Codes"
	MainContainer.Add(codeEntry)
	MainContainer.Add(widget.NewButton("Unlock", func() {
		code := codeEntry.Text
		if code != "" && len(code) == 8 {
			result := SendUnlockRequest(Bell8DigitEndPoint, code)
			dialog.ShowInformation("Result", parseCode(result), mainWindow)
		} else {
			dialog.ShowInformation("Error", "Please enter a valid code", mainWindow)
		}

	}))
	return MainContainer
}

func Sam16DigitTab(mainWindow fyne.Window) *fyne.Container {
	MainContainer := container.NewVBox()
	//mainWindow.SetTitle("Samsung 16 Digit Bios")
	SamLogo := canvas.NewImageFromResource(icns.ImageFromUrl(SamLogoUri))
	SamLogo.SetMinSize(fyne.NewSize(260, 160))
	MainContainer.Add(SamLogo)
	MainContainer.Add(widget.NewLabel("Samsung 16-Digit BIOS Unlock Tool"))
	MainContainer.Add(widget.NewLabel("This is the Samsung 16-Digit BIOS Unlock Tool"))
	MainContainer.Add(widget.NewLabel("Insert code from bios here, should be like: 07088120410C0000"))
	MainContainer.Add(widget.NewLabel("Please note that the code should be entered without the hyphens"))
	MainContainer.Add(widget.NewLabel("This tool may not work for all models"))
	codeEntry := widget.NewEntry()
	codeEntry.PlaceHolder = "Only 16-Digit Codes"
	MainContainer.Add(codeEntry)
	MainContainer.Add(widget.NewButton("Unlock", func() {
		code := codeEntry.Text
		if code != "" && len(code) == 16 {
			result := SendUnlockRequest(Sam16DigitEndPoint, code)
			dialog.ShowInformation("Result", parseCode(result), mainWindow)
		} else {
			dialog.ShowInformation("Error", "Please enter a valid code", mainWindow)
		}

	}))
	return MainContainer
}

func Sam18DigitTab(mainWindow fyne.Window) *fyne.Container {
	MainContainer := container.NewVBox()
	//mainWindow.SetTitle("Samsung 18 Digit Bios")
	SamLogo := canvas.NewImageFromResource(icns.ImageFromUrl(SamLogoUri))
	SamLogo.SetMinSize(fyne.NewSize(260, 160))
	MainContainer.Add(SamLogo)
	MainContainer.Add(widget.NewLabel("Samsung 18-Digit BIOS Unlock Tool"))
	MainContainer.Add(widget.NewLabel("This is the Samsung 18-Digit BIOS Unlock Tool"))
	MainContainer.Add(widget.NewLabel("Insert code from bios here, should be like: 3104CD2BDB962ABD900"))
	MainContainer.Add(widget.NewLabel("Please note that the code should be entered without the hyphens"))
	MainContainer.Add(widget.NewLabel("This tool may not work for all models"))
	codeEntry := widget.NewEntry()
	codeEntry.PlaceHolder = "Only 18-Digit Codes"
	MainContainer.Add(codeEntry)
	MainContainer.Add(widget.NewButton("Unlock", func() {
		code := codeEntry.Text
		if code != "" && len(code) == 18 {
			result := SendUnlockRequest(Sam18DigitEndPoint, code)
			dialog.ShowInformation("Result", parseCode(result), mainWindow)
		} else {
			dialog.ShowInformation("Error", "Please enter a valid code", mainWindow)
		}

	}))
	return MainContainer
}

func Son16DigitTab(mainWindow fyne.Window) *fyne.Container {
	MainContainer := container.NewVBox()
	//mainWindow.SetTitle("Sony 16 Digit Bios")
	SonyLogo := canvas.NewImageFromResource(icns.ImageFromUrl(SonyLogoUrl))
	SonyLogo.SetMinSize(fyne.NewSize(260, 160))
	MainContainer.Add(SonyLogo)
	MainContainer.Add(widget.NewLabel("Sony 16-Digit BIOS Unlock Tool"))
	MainContainer.Add(widget.NewLabel("This is the Sony 16-Digit BIOS Unlock Tool"))
	MainContainer.Add(widget.NewLabel("Insert code from bios here, should be like: DJP9-8DFJ-782D-X29J"))
	codeEntry := widget.NewEntry()
	codeEntry.PlaceHolder = "Only 16-Digit Codes"
	MainContainer.Add(codeEntry)
	MainContainer.Add(widget.NewButton("Unlock", func() {
		code := codeEntry.Text
		if code != "" && len(code) == 19 && strings.Count(code, "-") == 3 {
			result := SendUnlockRequest(Son16DigitEndPoint, code)
			dialog.ShowInformation("Result", parseCode(result), mainWindow)
		} else {
			dialog.ShowInformation("Error", "Please enter a valid code", mainWindow)
		}

	}))
	return MainContainer
}

func Asus8DigitTab(mainWindow fyne.Window) *fyne.Container {
	MainContainer := container.NewVBox()
	//mainWindow.SetTitle("Asus 8 Digit Bios")
	AsusLogo := canvas.NewImageFromResource(icns.ImageFromUrl(AsusLogoUri))
	AsusLogo.SetMinSize(fyne.NewSize(260, 200))
	MainContainer.Add(AsusLogo)
	MainContainer.Add(widget.NewLabel("Asus 8-Digit BIOS Unlock Tool"))
	MainContainer.Add(widget.NewLabel("This is the Asus 8-Digit BIOS Unlock Tool"))
	MainContainer.Add(widget.NewLabel("Insert date from bios here, should be something like: 2002-01-02"))
	MainContainer.Add(widget.NewLabel("This tool may not work for all models"))
	codeEntry := widget.NewEntry()
	codeEntry.PlaceHolder = "Only 8-Digit Codes"
	MainContainer.Add(codeEntry)
	MainContainer.Add(widget.NewButton("Unlock", func() {
		code := codeEntry.Text
		if code != "" && len(code) == 10 && strings.Count(code, "-") == 2 {
			result := SendUnlockRequest(Asus8DigitEndPoint, code)
			dialog.ShowInformation("Result", parseCode(result), mainWindow)
		} else {
			dialog.ShowInformation("Error", "Please enter a valid code", mainWindow)
		}

	}))

	return MainContainer
}

func MasterBiosUnlockRequest(code string) []interface{} {
	client := &http.Client{}
	req, err := http.NewRequest("GET", MasterEndpoint+code, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var jsonArray []interface{}
	err = json.Unmarshal([]byte(bodyText), &jsonArray)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	return jsonArray
}

func MasterBiosUnlockTab(mainWindow fyne.Window) *fyne.Container {
	MainContainer := container.NewVBox()
	MasterKeyLogo := canvas.NewImageFromResource(icns.ImageFromUrl(Mkey))
	MasterKeyLogo.SetMinSize(fyne.NewSize(100, 256))
	MainContainer.Add(MasterKeyLogo)
	MainContainer.Add(widget.NewLabel("Master Key"))
	MainContainer.Add(widget.NewLabel("This is a bios master key finder for other unlisted platforms"))
	MainContainer.Add(widget.NewLabel("Please enter a code provided by your platform in length/format"))
	MainContainer.Add(widget.NewLabel("This may not work for all models"))
	codeEntry := widget.NewEntry()
	codeEntry.PlaceHolder = "Any Format / Length Codes"
	MainContainer.Add(codeEntry)
	MainContainer.Add(widget.NewButton("Unlock", func() {
		var MasterResult = ""
		code := codeEntry.Text
		if code != "" {
			result := MasterBiosUnlockRequest(code)
			fmt.Println(result)
			if len(result) > 0 {
				for _, item := range result {
					switch v := item.(type) {
					case map[string]interface{}:
						vendor := v["vendor"].(string)
						biosCode := v["biosCode"]

						var biosCodeStrings []string
						if b, ok := biosCode.([]interface{}); ok {
							for _, bc := range b {
								biosCodeStrings = append(biosCodeStrings, bc.(string))
							}
						} else if s, ok := biosCode.(string); ok {
							biosCodeStrings = []string{s}
						} else {
							log.Fatalf("Invalid biosCode format: %v", biosCode)
						}

						biosCodeDisplay := strings.Join(biosCodeStrings, ", ")
						//fmt.Printf("Vendor: %s, BIOS Codes: %s\n", vendor, biosCodeDisplay)
						MasterResult += fmt.Sprintf("Vendor: %s, BIOS Codes: %s\n", vendor, biosCodeDisplay)
					default:
						fmt.Printf("Unexpected JSON format: %+v", item)
					}
				}

				dialog.ShowInformation("Result", MasterResult, mainWindow)
			} else {
				dialog.ShowInformation("Error", "Unable to find any matching platforms", mainWindow)
			}

		} else {
			dialog.ShowInformation("Error", "You have not entered a bios code", mainWindow)
		}

	}))
	return MainContainer
}
