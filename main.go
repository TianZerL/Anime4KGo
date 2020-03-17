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
	m := img.LoadImg(opt.InputFile)
	m.ShowInfo(opt)
	s := time.Now()
	m.Process(opt.Passes, opt.StrengthColor, opt.StrengthGradient, opt.FastMode)
	t := time.Since(s)
	fmt.Println("Total time for processing:", t)
	m.SaveImg(opt.OutputFile)
}
