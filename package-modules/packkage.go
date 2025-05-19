package main

import (
	"fmt"
	// Import the custom utils package
	// Make sure go.mod is created and the import path matches the module name
	"golang-foundation/package-modules/utils"
)

// Reuse: You can use utils functions in any file across the project
// Clean Code: Organize functions logically (auth, db, utils, etc.)
// Maintenance: Easier to manage large projects

// utils is one custom package.
// We create reusable functions in this package and import them where needed.

func main() {
	// Generate a user ID using GenerateNumber from the utils package
	userId := utils.GenerateNumber()
	fmt.Println(userId)
}

// go mod init <module-name>
// --------------------------
// Creates a go.mod file to manage dependencies in your Go project.
// It also defines the module path (used in imports).

// go mod tidy
// --------------------------
// What it does?
// - Adds missing dependencies
// - Removes unused ones
// - Updates go.sum
//
// When to use?
// - After adding/removing imports
// - After pulling new code

// go.sum
// --------------------------
// Automatically generated file.
// Stores version + hash of each dependency.
// Helps ensure dependency integrity (security).
// --------------------------

//PACKAGE
// package is cretae .go file
// its help to use the package to use other file

// MODULES
// module is cretae go.mod file
// it a golabal file
