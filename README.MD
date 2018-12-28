# Jumpstart
jumpstart is a CLI tool to organize and save all of your quickstart code links.

## Installation
Run go get <<link>> to install the program

    go get https://github.com/vinaycalastry/jumpstart.git
    
## Usage
Running `jumpstart` will display all available commands

 - add - *adds a new projectname and its github link to storage*
	 - `jumpstart add <<project-name>> <<github-link>>`
	 - Example: `jumpstart add electron https://github.com/electron/electron-quick-start.git`
 - list - *lists all the stored projects and its links*
	 - `jumpstart list`
	 - Sample Output:
		 The following links to quickstart codes are stored:
		1: electron (https://github.com/electron/electron-quick-start.git)
		2: html5 (https://github.com/h5bp/html5-boilerplate.git)
 - new - *Downloads the quick-start code for the given project*
	 - Browse to path
	 - `cd <pathname>`
	 - Then, run `jumpstart new <<project-name>>`
	 - Example: `jumpstart new electron` 

