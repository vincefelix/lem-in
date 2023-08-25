# Lem in  


## Table of Contents
1. [Description](#description)
2. [Authors](#authors)
3. [Usage:](#usage)
4. [File system:](#file-system)
5. [How to use the program:](#how-to-use-the-program)
6. [Implementation details: algorithm](#implementation-details-algorithm)

### Description:
***
***Hi Talent* !**   
Lem in is program that will read from a file (describing the ants and the colony) given in the arguments.  
Upon successfully finding the quickest path, lem-in will display the content of the file passed as argument and each move the ants make from room to room.  
This project is written using GO.

### Go,Html,Css:
***
- **GO**, also called Golang or Go language, is an open source programming language that Google developed.  
formore details check their website : [https://golang.org]

### Authors:
***
+ Seynabou Niang (*sniang*)  -  [https://learn.zone01dakar.sn/git/sniang]
* Masseck Thiaw (*mthiaw*) - **captain** - [https://learn.zone01dakar.sn/git/mthiaw]
- Vincent Félix Ndour (*vindour*) - [https://learn.zone01dakar.sn/git/vindour]

### Usage:
***
A little intro about how to install:
```
$ git clone https://learn.zone01dakar.sn/git/mthiaw/lem-in
$ cd lem in
```

## file-system  

The file system  looks like this:
```  
.
|
|____📁audit_files
|    |-----------📄badexample00.txt
|    |-----------📄badexample01.txt
|    |-----------📄example00.txt
|    |-----------📄example01.txt
|    |-----------📄example02.txt
|    |-----------📄example03.txt
|    |-----------📄example04.txt
|    |-----------📄example05.txt
|    |-----------📄example06.txt
|    |-----------📄example07.txt
|
|____📁tools
|    |--------------📄ant.go
|    |--------------📄functions.go
|    |--------------📄path.go
|    |--------------📄prog.go
|    |--------------📄roomstruct.go
|    |--------------📄validity.go
|
|____⚙go.mod
|
|____📝main.go
|
|____📜README.md
```
## How-to-use-the-program    
To use the program, you need to provide the path to the input
 file as an argument.  
**_PS: if the file is in a directory, provide the file path_**      
    
<u>**Example :**</u>  

    $go run main.go audit_file/example00.txt

<u>**Output :**</u>

```
4
##start
0 0 3
3 4 0
##end
2 2 5
1 8 3
0-2
2-3
3-1 

L1-3 L2-2
L1-2
L1-1 L2-3 L3-2
L2-1 L3-3 L4-2
L3-1 L4-3
L4-1
$

```

## Implementation-details-algorithm:
The program is divided into several files:

* `main.go`: This is the main file that imports the necessary packages and starts the program.
* `tools/functions.go`: This file contains various functions that are used throughout the program.
* `tools/path.go`: This file contains functions that are used to find paths between rooms.
* `tools/prog.go`: This file contains the main logic of the program.
* `tools/roomstruct.go`: This file defines the `Room` and `Link` structs.
* `tools/validity.go`: This file contains functions that are used to check the validity of the input file.  


in order to solve this enigma, the program go through 3 steps:
* #### `  Step one - check the validity of the file :`
    In this step we read each line from the input file and verify its format (number of fields). If it's not valid then exit with error message     
 
* #### `Step two - Find all possible paths from start room to end room :`
    
    Here we range the whole anthill find rooms connected from start to end .

* #### `Step three - Filter paths :`
    We filter out those paths which has collision with other paths. That means there is a repeated room. 

* #### `Step four - assign each ant to a path :`
    Thanks to Edmonds Karp's algorithm we determines the number of ant per path and the quickest for all the colony to move in few steps  
for more details about the algorithm check out this website : [https://medium.com/@jamierobertdawson/lem-in-finding-all-the-paths-and-deciding-which-are-worth-it-2503dffb893]

* #### `Step four - move the ants :`   
    in this part, we give each ant an entrance index , the position according how it will get into the paths.  

* #### `Step five - Simulating the movements : `
    In this step, we simulate the movement of the ants in the anthill by moving them along their assigned paths and their entrance index. We move one ant per room in each step.




    *@Licensed by team VMS*