package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

type IconUtil struct{}

func (iu IconUtil) IconFromBytes(IconName string, IconBytes []byte) fyne.Resource {
	return fyne.NewStaticResource(IconName, IconBytes)
}

func (iu IconUtil) IconFromRepo(name string) fyne.Resource {
	IconsRepoBaseUrl := "https://raw.githubusercontent.com/ZachC137/MICNS/e2b386038b1465856a3f46735a384b80d5511655/"
	// http get request to the repo to get the icon
	resp, err := http.Get(IconsRepoBaseUrl + name + ".png")
	if err != nil {
		log.Fatalf("Failed to fetch icon: %v", err)
	}
	IconBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read icon response body: %v", err)
	}
	return fyne.NewStaticResource(name, IconBytes)
}

func (iu IconUtil) IconByteLoader(IconName string, IconsFolder string) []byte {
	if IconsFolder == "" {
		IconsFolder = "assets/"
	}
	IconBytes, err := ioutil.ReadFile(IconsFolder + IconName + ".png")
	if err != nil {
		log.Fatalf("Failed to load icon: %v", err)
	}
	return IconBytes
}

func (iu IconUtil) Icon(name string) fyne.Resource {
	return fyne.NewStaticResource(name, iu.IconByteLoader(name, ""))
}

func (iu IconUtil) Icons8(uuid string, name string, category string) fyne.Resource {
	if category == "" {
		category = "color"
	}
	BaseUrl := fmt.Sprintf("https://img.icons8.com/%s/%s/%s", category, uuid, name)
	resp, err := http.Get(BaseUrl)
	if err != nil {
		log.Fatalf("Failed to fetch icon: %v", err)
	}
	IconBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read icon response body: %v", err)
	}
	return fyne.NewStaticResource(name, IconBytes)
}
func (iu IconUtil) ImageFromUrl(ImageUrl string) fyne.Resource {

	resp, err := http.Get(ImageUrl)
	if err != nil {
		log.Fatalf("Failed to fetch icon: %v", err)
	}
	IconBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read icon response body: %v", err)
	}
	return fyne.NewStaticResource(string(rand.Int()), IconBytes)
}

type ImageUtil struct{}

func NewUuid() string {
	uuid_base, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	return uuid_base.String()
}

func (iu ImageUtil) LoadImageFromBytes(name string, data []byte) fyne.Resource {
	return fyne.NewStaticResource(name, data)
}
func (iu ImageUtil) LoadImageFromUri(name string, uri string) fyne.Resource {
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalf("Failed to fetch image: %v", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read image response body: %v", err)
	}
	return iu.LoadImageFromBytes(name, data)
}
func (iu ImageUtil) NewCanvasImageUri(w float32, h float32, uri string) *fyne.Container {
	Image_ := canvas.NewImageFromResource(iu.LoadImageFromUri(NewUuid(), uri))
	// Image_.FillMode = canvas.ImageFillContain
	Image_.SetMinSize(fyne.Size{Width: w, Height: h})
	Image_.Resize(fyne.Size{Width: w, Height: h})
	view := container.NewHBox(Image_)
	return view
}

func (iu ImageUtil) NewCanvasImageFile(w float32, h float32, filePath string) *fyne.Container {
	Image_ := canvas.NewImageFromFile(filePath)
	Image_.SetMinSize(fyne.Size{Width: w, Height: h})
	Image_.Resize(fyne.Size{Width: w, Height: h})
	view := container.NewHBox(Image_)
	return view
}
