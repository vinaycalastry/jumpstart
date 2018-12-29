# Jumpstart
jumpstart is a CLI tool to organize and save all of your quickstart code links.

*PS: Requires git and golang to be installed for this to work. Prebuilt binaries will be uploaded once project reaches a good development stage*

## Installation
Run `go get <<link>>` to install the program

    go get github.com/vinaycalastry/jumpstart
    
## Usage
Running `jumpstart` will display all available commands

 - new - *Downloads the quick-start code for the given project*
	 - Browse to required path: `cd <pathname>`
	 - Then run `jumpstart new <<project-name>>`
	 - Example: `jumpstart new electron` 

 - add - *adds a new projectname and its github link to storage*
	 - `jumpstart add <<project-name>> <<github-link>>`
	 - Example: `jumpstart add electron https://github.com/electron/electron-quick-start.git`

 - list - *lists all the stored projects and its links*
	 - `jumpstart list`
	 - Sample Output:
		```The following links to quickstart codes are stored:
		1: electron (https://github.com/electron/electron-quick-start.git)
		2: html5 (https://github.com/h5bp/html5-boilerplate.git)
	
 - delete - *deletes a project link from storage*
	 - `jumpstart delete <<project-name>>`
	 - Example: `jumpstart delete electron`

 - replace - *replaces an existing project's link with a provided one*
	 - `jumpstart replace <<project-name>> <<github-link>>`
	 - Example: `jumpstart replace electron https://github.com/electron/electron-quick-start.git`

