package main

import (
	"github.com/kbinani/screenshot"
	"image/jpeg"
	"os"
	"fmt"
	"time"
	"log"
)

func main() {
	
	captureDirName := "captures"

	_, err := os.Stat(captureDirName)
	if err != nil{
		if os.IsNotExist(err) {
			log.Println("[MSG] Create Directory")
			err = os.MkdirAll(captureDirName, os.ModePerm)
			if err != nil{
				log.Fatal(err)
			}
		}
	}

	ticker := time.NewTicker(time.Second * 2)
	deleteCounter := 0
	fileName := ""

	for currentTime := range ticker.C {
		deleteCounter = deleteCounter + 2;
		if(deleteCounter > 60){
			log.Println("[MSG_] Remove Directory")
			err := os.RemoveAll(captureDirName)
			if err != nil{
				log.Fatal(err)
			}
			log.Println("[MSG_] Create Directory")
			err = os.MkdirAll(captureDirName, 0600)
			if err != nil{
				log.Fatal(err)
			}
			deleteCounter = 0
		}
		bounds := screenshot.GetDisplayBounds(0)		// first display

		img, err := screenshot.CaptureRect(bounds)		// detect한 영역 만큼 캡쳐
		if err != nil {
			//panic(err)
			log.Fatal(err)
		}					

		equipment_name := "dev001"
		fileName = fmt.Sprintf("%s/%s_%s.jpg",captureDirName, equipment_name, currentTime.Format("20060102_150405"))
		file, _ := os.Create(fileName)
		
		//imageOpt := jpeg.Options
		var imageOpt jpeg.Options
		imageOpt.Quality = 60

		jpeg.Encode(file, img, &imageOpt)
		file.Close()

		log.Printf("[MSG_] captured.. %s\n", fileName)
	}
	
	
}

// package main

// import (
// 	"github.com/kbinani/screenshot"
// 	"image/png"
// 	"os"
// 	"fmt"
// )

// func main() {
// 	n := screenshot.NumActiveDisplays()

// 	// Only Capture on first display

// 	for i := 0; i < n; i++ {
// 		bounds := screenshot.GetDisplayBounds(i)

// 		img, err := screenshot.CaptureRect(bounds)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
// 		file, _ := os.Create(fileName)
// 		defer file.Close()
// 		png.Encode(file, img)

// 		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"image"
// 	"reflect"
// 	"syscall"
// 	"unsafe"
// 	"image/png"
// 	"os"
// 	"github.com/disintegration/gift"
// )

// func main() {
// 	// for windows 8 or later
// 	//procSetProcessDpiAwareness.Call(uintptr(2)) // PROCESS_PER_MONITOR_DPI_AWARE

// 	fmt.Println("RUN")
	
// 	img, err := capture()
// 	if err != nil {
// 		fmt.Println("error",err)
//         return
//     }

// 	file, _ := os.Create("test.png")
// 	defer file.Close()
// 	//png.Encode(file, img)
// 	png.Encode(file, img)
// 	fmt.Println("file created..")

// }

// func capture() (image.Image, error) {
// 	// Find the window
// 	handle, err := findWindow()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Determine the full width and height of the window
// 	rect, err := windowRect(handle)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Capture!
// 	return captureWindow(handle, rect)
// }

// var (
// 	modUser32         = syscall.NewLazyDLL("User32.dll")
// 	procFindWindow    = modUser32.NewProc("FindWindowW")
// 	procGetClientRect = modUser32.NewProc("GetClientRect")
// 	procGetDC         = modUser32.NewProc("GetDC")
// 	procReleaseDC     = modUser32.NewProc("ReleaseDC")

// 	modGdi32                   = syscall.NewLazyDLL("Gdi32.dll")
// 	procBitBlt                 = modGdi32.NewProc("BitBlt")
// 	procCreateCompatibleBitmap = modGdi32.NewProc("CreateCompatibleBitmap")
// 	procCreateCompatibleDC     = modGdi32.NewProc("CreateCompatibleDC")
// 	procCreateDIBSection       = modGdi32.NewProc("CreateDIBSection")
// 	procDeleteDC               = modGdi32.NewProc("DeleteDC")
// 	procDeleteObject           = modGdi32.NewProc("DeleteObject")
// 	procGetDeviceCaps          = modGdi32.NewProc("GetDeviceCaps")
// 	procSelectObject           = modGdi32.NewProc("SelectObject")

// 	modShcore = syscall.NewLazyDLL("Shcore.dll")
// 	procSetProcessDpiAwareness = modShcore.NewProc("SetProcessDpiAwareness")
// )

// const (
// 	// GetDeviceCaps constants from Wingdi.h
// 	deviceCaps_HORZRES = 8
// 	deviceCaps_VERTRES = 10
// 	deviceCaps_LOGPIXELSX = 88
// 	deviceCaps_LOGPIXELSY = 90

// 	// BitBlt constants
// 	bitBlt_SRCCOPY = 0x00CC0020
// )

// // Windows RECT structure
// type win_RECT struct {
// 	Left, Top, Right, Bottom int32
// }

// // http://msdn.microsoft.com/en-us/library/windows/desktop/dd183375.aspx
// type win_BITMAPINFO struct {
// 	BmiHeader win_BITMAPINFOHEADER
// 	BmiColors *win_RGBQUAD
// }

// // http://msdn.microsoft.com/en-us/library/windows/desktop/dd183376.aspx
// type win_BITMAPINFOHEADER struct {
// 	BiSize          uint32
// 	BiWidth         int32
// 	BiHeight        int32
// 	BiPlanes        uint16
// 	BiBitCount      uint16
// 	BiCompression   uint32
// 	BiSizeImage     uint32
// 	BiXPelsPerMeter int32
// 	BiYPelsPerMeter int32
// 	BiClrUsed       uint32
// 	BiClrImportant  uint32
// }

// // http://msdn.microsoft.com/en-us/library/windows/desktop/dd162938.aspx
// type win_RGBQUAD struct {
// 	RgbBlue     byte
// 	RgbGreen    byte
// 	RgbRed      byte
// 	RgbReserved byte
// }

// // findWindow finds the handle to the window.
// func findWindow() (syscall.Handle, error) {
// 	var handle syscall.Handle

// 	// First look for the normal window
// 	ret, _, _ := procFindWindow.Call(
// 		0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("NTCM2"))))
// 	if ret == 0 {
// 		return handle, fmt.Errorf("App not found. Is it running?")
// 	}

// 	handle = syscall.Handle(ret)
// 	return handle, nil
// }

// // windowRect gets the dimensions for a Window handle.
// func windowRect(hwnd syscall.Handle) (image.Rectangle, error) {
// 	var rect win_RECT
// 	ret, _, err := procGetClientRect.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&rect)))
// 	if ret == 0 {
// 		return image.Rectangle{}, fmt.Errorf("Error getting window dimensions: %s", err)
// 	}

// 	return image.Rect(0, 0, int(rect.Right), int(rect.Bottom)), nil
// }

// // captureWindow captures the desired area from a Window and returns an image.
// func captureWindow(handle syscall.Handle, rect image.Rectangle) (image.Image, error) {
// 	// Get the device context for screenshotting
// 	dcSrc, _, err := procGetDC.Call(uintptr(handle))
// 	if dcSrc == 0 {
// 		return nil, fmt.Errorf("Error preparing screen capture: %s", err)
// 	}
// 	defer procReleaseDC.Call(0, dcSrc)

// 	// Grab a compatible DC for drawing
// 	dcDst, _, err := procCreateCompatibleDC.Call(dcSrc)
// 	if dcDst == 0 {
// 		return nil, fmt.Errorf("Error creating DC for drawing: %s", err)
// 	}
// 	defer procDeleteDC.Call(dcDst)

//   	// Determine the width/height of our capture
// 	width := rect.Dx();
// 	height := rect.Dy();

// 	// Get the bitmap we're going to draw onto
// 	var bitmapInfo win_BITMAPINFO
// 	bitmapInfo.BmiHeader = win_BITMAPINFOHEADER{
// 		BiSize:        uint32(reflect.TypeOf(bitmapInfo.BmiHeader).Size()),
// 		BiWidth:       int32(width),
// 		BiHeight:      int32(height),
// 		BiPlanes:      1,
// 		BiBitCount:    32,
// 		BiCompression: 0, // BI_RGB
// 	}
// 	bitmapData := unsafe.Pointer(uintptr(0))
// 	bitmap, _, err := procCreateDIBSection.Call(
// 		dcDst,
// 		uintptr(unsafe.Pointer(&bitmapInfo)),
// 		0,
// 		uintptr(unsafe.Pointer(&bitmapData)), 0, 0)
// 	if bitmap == 0 {
// 		return nil, fmt.Errorf("Error creating bitmap for screen capture: %s", err)
// 	}
// 	defer procDeleteObject.Call(bitmap)

// 	// Select the object and paint it
// 	procSelectObject.Call(dcDst, bitmap)
// 	ret, _, err := procBitBlt.Call(
// 		dcDst, 0, 0, uintptr(width), uintptr(height),
// 		dcSrc, uintptr(rect.Min.X), uintptr(rect.Min.Y), bitBlt_SRCCOPY)
// 	if ret == 0 {
// 		return nil, fmt.Errorf("Error capturing screen: %s", err)
// 	}

// 	// Convert the bitmap to an image.Image. We first start by directly
// 	// creating a slice. This is unsafe but we know the underlying structure
// 	// directly.
// 	var slice []byte
// 	sliceHdr := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
// 	sliceHdr.Data = uintptr(bitmapData)
// 	sliceHdr.Len = width * height * 4
	
// 	sliceHdr.Cap = sliceHdr.Len

// 	// Using the raw data, grab the RGBA data and transform it into an image.RGBA
// 	imageBytes := make([]byte, len(slice))
// 	for i := 0; i < len(imageBytes); i += 4 {
// 		imageBytes[i], imageBytes[i+2], imageBytes[i+1], imageBytes[i+3] = slice[i+2], slice[i], slice[i+1], 255
// 	}

//   	// The window gets screenshotted upside down and I don't know why.
// 	// Flip it.
// 	img := &image.RGBA{imageBytes, 4 * width, image.Rect(0, 0, width, height)}
// 	dst := image.NewRGBA(img.Bounds())
// 	gift.New(gift.FlipVertical()).Draw(dst, img)

// 	return dst, nil
// }
