Tool-Chain
---

An easily expandable command line app written in Golang. Currently, Tool-Chain is capable of maintaining command line 'profiles'. 
profiles are simply a set of environment exports. This is useful for swapping cloud profiles, such as AWS access keys, or GCP service accounts. 
It can also be used to save build configurations/environments. For example, exports may differ between C, C++, Java or Go projects. 
This tool aims to make organization of these exports easier and less prone to human error. 



### Usage

Clone the repository and build the binary using the command `go build -o tool-chain`. Afterwards, add the compiled binary to your path. 

To setup Tool-Chain for the first type simply type the command without any arguments. Once you complete the setup process you will have 
a profile defined and the exports file linked within your `zshrc` or `bashrc`.



### Commands

`tool-chain profile new profileName` will create a new profile named `profileName`. Adding a profile will prompt you for all environment variables you wish to store in the profile.  

`tool-chain profile activate profileName` will uncomment all exports contained within the profile, and comment out all other profiles. You must manually source your rc file for the changes to take effect. 

`tool-chain profile list` will list all configured profiles.

`tool-chain profile remove profileName` will remove the desired profile from the exports file if it exists.  

`tool-chain net scan` will print the IP address and hostname of all hosts on the provided subnet, so long as it is accessible. 


#### Why does Tool-Chain ask for the path of my `zshrc`/`bashrc` file?
In order to not interfere with the format of your `zshrc` or `bashrc`, and to make the management of profiles easier, a 
separate `exports` file is created, and a single line is added to your rc file. This command is the following,

`source ${PATH_TO_TOOL_CHAIN_BINARY}/exports`   

This export will only ever be added to your RC file at the time of setup, however subsequent setups will not add additional lines to your rc file.
You may add additional exports manually, as long as you maintain the header and footer format.  


