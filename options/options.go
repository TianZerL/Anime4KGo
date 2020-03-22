package options

import (
	"flag"
)

//Opt defined options which can be accepted
type Opt struct {
	Passes           int
	FastMode         bool
	Help             bool
	InputFile        string
	OutputFile       string
	StrengthColor    float64
	StrengthGradient float64
}

var options Opt

func init() {
	flag.StringVar(&options.InputFile, "i", "./pic/p1.png", "File for loading")
	flag.StringVar(&options.OutputFile, "o", "out.png", "File for outputting")
	flag.BoolVar(&options.FastMode, "f", false, "Fast Mode but low quality")
	flag.BoolVar(&options.Help, "h", false, "Show help information")
	flag.BoolVar(&options.Help, "?", false, "Show help information")
	flag.IntVar(&options.Passes, "p", 2, "Passes for processing")
	flag.Float64Var(&options.StrengthColor, "sc", 1.0/3.0, "Strength for pushing color,range 0 to 1,higher for thinner")
	flag.Float64Var(&options.StrengthGradient, "sg", 1.0, "Strength for pushing gradient,range 0 to 1,higher for sharper")
	flag.Parse()
}

//NewOptions return a options object pointer
func NewOptions() *Opt {
	return &options
}

//Usage print help infomation
func (o *Opt) Usage() {
	flag.PrintDefaults()
}
