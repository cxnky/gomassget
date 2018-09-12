package main

import (
	"flag"
	"github.com/cxnky/gomassget/utils"
	"fmt"
	"strconv"
)

/**

 * Created by cxnky on 12/09/2018 at 10:43
 * betterget
 * https://github.com/cxnky/
 
**/

func main() {

	var textFile string

	flag.StringVar(&textFile, "file", "packages.txt", "define the name of the text file to mass import packages from")
	flag.Parse()

	pwd := utils.GetPWD()

	fileLocation := pwd + "\\" + textFile

	count := utils.InstallPackagesFromFile(fileLocation)

	fmt.Println(strconv.Itoa(count) + " package(s) have been installed.")


}