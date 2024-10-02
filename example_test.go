package zip7

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleAdd() {
	input := "log.txt"
	output := "log.7z"

	out, err := Add(
		Props{
			File:        output,
			Compression: 6,
			Password:    "mYSuppaPAssWORD",
			Threads:     4,
			Delete:      true,
		}, input)

	fmt.Printf("p7zip output: %s\n", out)
	fmt.Printf("Error: %v\n", err)
}

func ExampleAddList() {
	input1 := "log1.txt"
	input2 := "log2.txt"
	input3 := "log3.txt"
	output := "log.7z"

	out, err := AddList(Props{File: output}, []string{input1, input2, input3})

	fmt.Printf("p7zip output: %s\n", out)
	fmt.Printf("error: %v\n", err)
}

func ExampleExtract() {
	out, err := Extract(Props{File: "log.7z"})

	fmt.Printf("p7zip output: %s\n", out)
	fmt.Printf("Error: %v\n", err)
}

func ExampleList() {
	info, err := List(Props{File: "log.7z"})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Path: %s\n", info.Path)
	fmt.Printf("Type: %s\n", info.Type)
	fmt.Printf("Solid: %t\n", info.Solid)
	fmt.Printf("Blocks: %d\n", info.Blocks)
	fmt.Printf("PhysicalSize: %d\n", info.PhysicalSize)
	fmt.Printf("HeadersSize: %d\n", info.HeadersSize)

	fmt.Println("Files:")

	for _, file := range info.Files {
		fmt.Printf("  %s (size: %d)\n", file.Path, file.Size)
	}
}

func ExampleCheck() {
	ok, err := Check(Props{File: "log.7z"})

	fmt.Printf("File ok: %t\n", ok)
	fmt.Printf("Error: %v\n", err)
}

func ExampleDelete() {
	out, err := Delete(Props{File: "log.7z"}, "dir1/file1", "dir2/file2")

	fmt.Printf("p7zip output: %s\n", out)
	fmt.Printf("Error: %v\n", err)
}
