package generator

import (
	"fmt"
	"os"
)

func CreateFolder(name string) {
	dirName := "./test/" + name
	// exec.Command(fmt.Sprintf("pnpm create vite %s", dirName))
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("Folder exist")
		} else {
			fmt.Println("Error creating", err)
		}
	} else {
		os.Create(dirName + "/main.ipynb")
	}
}
