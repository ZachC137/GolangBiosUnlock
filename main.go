package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/PuerkitoBio/goquery"
)

var (
	Acer8DigitEndPoint = "https://www.biosbug.com/gets/check_acer_availability.php"
	Sony7DigitEndPoint = "https://www.biosbug.com/gets/check_sonyseven_availability.php"
	HpCo5DigitEndPoint = "https://www.biosbug.com/gets/check_HP_old.php"
	Bell8DigitEndPoint = "https://www.biosbug.com/gets/check_bell_availability.php"
	Sam16DigitEndPoint = "https://www.biosbug.com/gets/check_samsung_availability.php"
	Sam18DigitEndPoint = "https://www.biosbug.com/gets/check_samsung2_availability.php"
	Son16DigitEndPoint = "https://www.biosbug.com/gets/check_sony_availability.php"
	Asus8DigitEndPoint = "https://www.biosbug.com/gets/check_availability.php"
)

var (
	icns = IconUtil{}
)

func SetHeaders(req *http.Request) {
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("cookie", "testCookie=1")
	req.Header.Set("origin", "https://www.biosbug.com")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("referer", "https://www.biosbug.com/acer/")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="8", "Chromium";v="126", "Google Chrome";v="126"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
}

func SendUnlockRequest(endpoint string, code string) string {
	client := &http.Client{}
	var data = strings.NewReader(`date=` + code)
	req, err := http.NewRequest("POST", endpoint, data)
	if err != nil {
		log.Fatal(err)
	}
	SetHeaders(req)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(bodyText)
}

func parseCode(htmlContent string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}

	code := doc.Find("#demo3").Text()

	return code
}

func main() {

	// code := "DJP9-8DFJ-782D-X29J"
	// RawResult := SendUnlockRequest(Son16DigitEndPoint, code)
	// result := parseCode(RawResult)
	// fmt.Println("Found code:", result)
	Gui := app.New()
	GuiWindow := Gui.NewWindow("Golang Bios Unlock")
	GuiWindow.SetIcon(icns.Icons8("256", "key.png", "ink"))
	GuiWindow.Resize(fyne.NewSize(550, 500))
	GuiWindow.SetFixedSize(true)
	Acer8 := container.NewTabItem("Acer 8 Digit", Acer8DigitTab(GuiWindow))
	Asus8 := container.NewTabItem("Asus 8 Digit", Asus8DigitTab(GuiWindow))
	HPCo5 := container.NewTabItem("HP CO 5 Digit", HpCo5DigitTab(GuiWindow))
	HPBell8 := container.NewTabItem("HP Bell 8 Digit", Bell8DigitTab(GuiWindow))
	Sam16 := container.NewTabItem("Samsung 16 Digit", Sam16DigitTab(GuiWindow))
	Sam18 := container.NewTabItem("Samsung 18 Digit", Sam18DigitTab(GuiWindow))
	Sony7 := container.NewTabItem("Sony 7 Digit", Sony7DigitTab(GuiWindow))
	Sony16 := container.NewTabItem("Sony 16 Digit", Son16DigitTab(GuiWindow))
	MasterBiosUnlock := container.NewTabItem("Master Bios Unlock", MasterBiosUnlockTab(GuiWindow))
	tabs := container.NewAppTabs(
		Acer8,
		Asus8,
		HPCo5,
		HPBell8,
		Sam16,
		Sam18,
		Sony7,
		Sony16,
		MasterBiosUnlock,
	)

	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationLeading)

	GuiWindow.SetContent(tabs)
	GuiWindow.ShowAndRun()
}
