package main

import(

)

func main() {
	cfg := config{
		commands: getCommands(),
	}
	
	startRepl(&cfg)
}