package main

import (
	"day02/internal/find"
	"flag"
	"log"
	"os"
)

const usageMsg = "Usage: ./myFind [-d][-f [-ext]][-sl] foldername"

func main() {
	directoryFlag := flag.Bool("d", false, "usage: [-d] for print directories")
	filesFlag := flag.Bool("f", false, "usage: [-f] for print files")
	symbolLinksFlag := flag.Bool("sl", false, "usage: [-sl] for print symbol links")
	extensionFlag := flag.String("ext", "", "usage: [-f[-ext]] for specify file's extension")

	flag.Parse()

	noFlags := !*directoryFlag && !*filesFlag && !*symbolLinksFlag
	correctExtensionFlagUsage := !(*extensionFlag != "" && !*filesFlag)

	if len(flag.Args()) != 1 {
		log.Fatal("wrong number of CLI arguments. ", usageMsg)
	}

	if !correctExtensionFlagUsage {
		log.Fatal("wrong usage of [-ext] flag. ", usageMsg)
	}

	if noFlags {
		*directoryFlag = true
		*filesFlag = true
		*symbolLinksFlag = true
	}

	filepath := flag.Arg(0)

	directoryContent, err := find.ReadContent(filepath)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	if *directoryFlag {
		directoryContent.PrintFolders(os.Stdout)
	}
	if *filesFlag {
		directoryContent.PrintFiles(os.Stdout, *extensionFlag)
	}
	if *symbolLinksFlag {
		directoryContent.PrintSymbolLinks(os.Stdout)
	}
}
