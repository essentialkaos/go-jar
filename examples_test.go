package jar

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleReadFile() {
	manifest, err := ReadFile("wagon-file-3.3.2.jar")

	fmt.Printf("Error: %v\n", err)
	fmt.Printf("Specification-Vendor: %s\n", manifest["Specification-Vendor"])
	fmt.Printf("Implementation-Vendor-Id: %s\n", manifest["Implementation-Vendor-Id"])
	fmt.Printf("Specification-Version: %s\n", manifest["Specification-Version"])
}
