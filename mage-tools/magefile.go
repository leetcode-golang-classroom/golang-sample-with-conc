//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Default = Build

// clean the build binary
func Clean() error {
	return sh.Rm("bin")
}

// Creates the binary in the current directory.
func Build() error {
	mg.Deps(Clean)
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}
	err := sh.Run("go", "build", "-o", "./bin/simple-task", "./cmd/simple-task/main.go")
	if err != nil {
		return err
	}
	err = sh.Run("go", "build", "-o", "./bin/pool-sample", "./cmd/pool-sample/main.go")
	if err != nil {
		return err
	}
	err = sh.Run("go", "build", "-o", "./bin/aggregate-sample", "./cmd/aggregate-sample/main.go")
	if err != nil {
		return err
	}
	err = sh.Run("go", "build", "-o", "./bin/handle-errors", "./cmd/handle-errors/main.go")
	if err != nil {
		return err
	}
	return sh.Run("go", "build", "-o", "./bin/handle-context", "./cmd/handle-context/main.go")
}

// start the simple task
func LaunchSimpleTask() error {
	mg.Deps(Build)
	err := sh.RunV("./bin/simple-task")
	if err != nil {
		return err
	}
	return nil
}

// start the pool sample
func LaunchPoolSample() error {
	mg.Deps(Build)
	err := sh.RunV("./bin/pool-sample")
	if err != nil {
		return err
	}
	return nil
}

// start the aggregate sample
func LaunchAggregateSample() error {
	mg.Deps(Build)
	err := sh.RunV("./bin/aggregate-sample")
	if err != nil {
		return err
	}
	return nil
}

// start the handle errors sample
func LaunchHandleErrorsSample() error {
	mg.Deps(Build)
	err := sh.RunV("./bin/handle-errors")
	if err != nil {
		return err
	}
	return nil
}

// start the handle context sample
func LaunchHandleContextSample() error {
	mg.Deps(Build)
	err := sh.RunV("./bin/handle-context")
	if err != nil {
		return err
	}
	return nil
}

// run the test
func Test() error {
	err := sh.RunV("go", "test", "-v", "./...")
	if err != nil {
		return err
	}
	return nil
}
