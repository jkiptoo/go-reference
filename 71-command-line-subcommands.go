// Using flag package to define subcommands each with their own flags
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	print := fmt.Println

	//Create a subcommand using 'NewFlagSet'then set flags specific to the function
	systemCommand := flag.NewFlagSet("SysConfig", flag.ExitOnError)
	systemEnable := systemCommand.Bool("Enable", false, "Enable")
	systemName := systemCommand.String("Name", "", "Name")

	//Create different subcommand with different flags
	SystemLimitCommand := flag.NewFlagSet("SysLimit", flag.ExitOnError)
	SystemLevel := SystemLimitCommand.Int("System Level", 0, "Level")
	
	//Pass subcommand as first argument to the program
	if len(os.Args) < 2 {
		print("System expected 'SysConfig' or 'SysLimit' subcommands")
		os.Exit(1)
	}

	//Check which subcommand has been invoked
	switch os.Args[1] {

	case "SysConfig":

		//For each subcommand, parse its flags and access trailing positional arguments
		systemCommand.Parse(os.Args[2:])
		print("Subcommand is 'SysConfig'")
		print(" Enable:", *systemEnable)
		print(" Name:", *systemName)
		print(" Tail:", systemCommand.Args())

	case "SystemLimit":
		SystemLimitCommand.Parse(os.Args[2:])
		print("Subcommand 'SysLimit'")
		print(" Level:", *SystemLevel)
		print(" Tail:", SystemLimitCommand.Args())
		
	default:
		print("Expected 'SysConfig' or 'SysLimit' subcommands")
		os.Exit(1)	
	}
}

//-------NOTES-----------//
// go build command-line-subcommands.go

//Invoke SysConfig subcommand
//'./command-line-subcommands SysConfig -enable -name=joe a1 a2'

//Next invoke System Limit subcommand
//'./command-line-subcommands SysLimit -level 7 a1'
 
//SysLimit can't accept SysConfig flags
//./command-line-subcommands bar -enable a1