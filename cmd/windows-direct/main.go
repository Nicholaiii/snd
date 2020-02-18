package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/BigJk/snd/printing/cups"
	"github.com/BigJk/snd/printing/remote"
	"github.com/BigJk/snd/printing/windows"

	"github.com/BigJk/snd"
	"github.com/asticode/go-astilectron"
)

func main() {
	debug := flag.Bool("debug", false, "")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	s, err := snd.NewServer("./data.db", snd.WithPrinter(&cups.CUPS{}), snd.WithPrinter(&remote.Remote{}), snd.WithPrinter(&windows.Direct{}))
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	go func() {
		_ = s.Start(":7123")
	}()

	var targetWriter io.Writer
	if !*debug {
		targetWriter = ioutil.Discard
	} else {
		targetWriter = os.Stdout
	}

	_ = os.Mkdir("./data", 0666)

	var a, _ = astilectron.New(log.New(targetWriter, "", 0), astilectron.Options{
		AppName:            "Sales & Dungeons",
		BaseDirectoryPath:  "./data",
		DataDirectoryPath:  "./data",
		AppIconDefaultPath: "icon.png",
		VersionElectron:    "8.0.1",
	})
	defer a.Close()
	a.Start()

	var w, _ = a.NewWindow("http://127.0.0.1:7123", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(720),
		Width:  astikit.IntPtr(1280),
	})
	w.Create()

	a.Wait()
}
