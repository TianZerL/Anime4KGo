package main

import (
	"fmt"
	"time"

	"github.com/TianZerL/Anime4KGo/img"
	"github.com/TianZerL/Anime4KGo/options"
)

func main() {
	opt := options.NewOptions()
	if opt.Help == true {
		opt.Usage()
		return
	}
	//Load image
	m := img.LoadImg(opt.InputFile)
	//Show basic infomation of image and process
	m.ShowInfo(opt)
	//Start timing for processing
	s := time.Now()
	//main process
	m.Process(opt.Passes, opt.StrengthColor, opt.StrengthGradient, opt.FastMode)
	t := time.Since(s)
	fmt.Println("Total time for processing:", t)
	//Save image to disk
	m.SaveImg(opt.OutputFile)
}
