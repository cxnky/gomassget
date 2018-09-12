package utils

import (
	"runtime"
	"path/filepath"
	"os"
	"bufio"
	"fmt"
	"os/exec"
)

/**

 * Created by cxnky on 12/09/2018 at 10:46
 * utils
 * https://github.com/cxnky/
 
**/

func GetOSName() string {

	switch runtime.GOOS {

	case "darwin":
		return "osx"

	case "windows":
		return runtime.GOOS

	case "linux":
		return runtime.GOOS

	default:
		return runtime.GOOS

	}

}

func GetPWD() string {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {

		panic("could not navigate to working directory")

	}

	return dir

}

func InstallPackagesFromFile(filePath string) int {

	file, err := os.Open(filePath)

	if err != nil {

		panic("unable to open package file: " + err.Error())

	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0

	for scanner.Scan() {

		fmt.Print("Installing ", scanner.Text(), "...")

		cmd := exec.Command("go", "get", scanner.Text())

		if err := cmd.Run(); err != nil {

			fmt.Print("error")

		} else {

			fmt.Println("done!")

		}

		i++
	}

	if err := scanner.Err(); err != nil {

		panic("unable to read from file: " + err.Error())

	}

	return i

}