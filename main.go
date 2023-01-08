package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

func main() {
	var isHelp *bool
	var inputFrom string
	var configFile string
	var parameter1, parameter2 *string

	isHelp = flag.Bool("h", false, "help")

	flag.StringVar(&inputFrom, "inputfrom", "cmd", "Where to get parameters. Options: \"cmd\", \"env\", \"config\"\nexample1: -inputfrom cmd -parameter1 asdf -parameter2 zxcv\nexample2: -inputfrom env\nexample3: -inputfrom config -config /tmp/para.conf\n")

	flag.StringVar(&configFile, "config", "", "The path of config file")

	parameter1 = flag.String("parameter1", "defaultP1", "test parameter1")
	parameter2 = flag.String("parameter2", "defaultP2", "test parameter2")

	// -inputfrom cmd -parameter1 asdf -parameter2 zxcv
	flag.Parse()

	// the command is "xxxx -h"
	if *isHelp {
		flag.Usage()
		return
	}
	var outPara1, outPara2 string

	switch inputFrom {
	case "env":
		outPara1, outPara2 = os.Getenv("PARAMETER1"), os.Getenv("PARAMETER2")
	case "config":
		viper.SetConfigFile(configFile)
		viper.SetConfigType("ini")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
		outPara1, outPara2 = viper.GetString("section.parameter1"), viper.GetString("section.parameter2")
	default:
		outPara1, outPara2 = *parameter1, *parameter2

	}

	fmt.Println("from "+inputFrom, outPara1, outPara2)

	// set an http server. http://localhost:25000/containerdemo
	http.HandleFunc("/containerdemo", func(w http.ResponseWriter, r *http.Request) {
		var output string
		output += fmt.Sprintf("Parameters from %s\n", inputFrom)
		output += fmt.Sprintf("parameter1 is %s\nparameter2 is %s\n", outPara1, outPara2)
		w.Write([]byte(output))
	})
	http.ListenAndServe(":25000", nil)
}
