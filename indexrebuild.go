package main

// Thanks to
// http://stackoverflow.com/questions/12657365/extracting-directory-hierarchy-using-go-language

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	tGlob0 := time.Now()

	fOut, err := os.Create("index.md")
	if err != nil {
		panic(err)
	}
	defer fOut.Close()

	// Define and implement what to do with path pieces
	visit := func(path string, info os.FileInfo, err error) error {
		// Split dir path from file name
		dir, file := filepath.Split(path)
		// Init list line string
		line := ""
		// Init and compute the level of indentation based on dit level
		indent := ""
		for i := 0; i < len(strings.Split(dir, "/"))-1; i++ {
			indent = indent + "    "
		}
		// If I reach a new dir
		if info.IsDir() {
			// Ignore local dir (exit function)
			if strings.Contains(path, ".") {
				return nil
			}
			// Create line for page with name == dirName
			line = indent + "* [" + strings.Title(filepath.Base(path)) + "](" + filepath.Join(path, filepath.Base(path)+".html)") + "\n"
			// If I found a new file
		} else {
			// Only consider md files
			if strings.HasSuffix(file, "md") {
				fileBase := strings.TrimSuffix(file, ".md")
				// Ignore files with name == dirName
				if fileBase == filepath.Base(dir) {
					return nil
				}
				// Create line
				linkName := strings.Replace(strings.Title(strings.Replace(fileBase, "-", "_", -1)), "_", " ", -1)
				line = indent + "* [" + linkName + "](" + filepath.Join(dir, strings.Replace(file, ".md", ".html", 1)) + ")\n"
			} else {
				return nil
			}
		}
		// Print line to file
		//         fmt.Print(line)
		if _, err = fOut.WriteString(line); err != nil {
			log.Fatal(err)
		}
		return nil
	}

	header := `<!-- 
.. link: 
.. description: 
.. tags: 
.. date: 2013/09/03 12:24:24
.. title: for future references summary
.. slug: index
-->

* [Blog](../index.html)
`

	if _, err = fOut.WriteString(header); err != nil {
		log.Fatal(err)
	}

	// Walk through folders
	err = filepath.Walk("./", visit)
	if err != nil {
		log.Fatal(err)
	}

	tGlob1 := time.Now()
	fmt.Println()
	log.Println("Recreated pages index in ", tGlob1.Sub(tGlob0))
} //END MAIN
