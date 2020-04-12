# minesweeper-API

This is an implementation of the game called Minesweeper.

Minesweeper has a very basic gameplay style. In its original form, mines are scattered throughout a board. This board is divided into cells, which have three states: uncovered, covered and flagged. A covered cell is blank and clickable, while an uncovered cell is exposed, either containing a number (the mines adjacent to it), or a mine. When a cell is uncovered by a player click, and if it bears a mine, the game ends. A flagged cell is similar to a covered one, in the way that mines are not triggered when a cell is flagged, and it is impossible to lose through the action of flagging a cell. However, flagging a cell implies that a player thinks there is a mine underneath, which causes the game to deduct an available mine from the display.

In order to win the game, players must logically deduce where mines exist through the use of the numbers given by uncovered cells. To win, all non-mine cells must be uncovered and all mine cells must be flagged. At this stage, the timer is stopped.

When a player left-clicks on a cell, the game will uncover it. If there are no mines adjacent to that particular cell, the mine will display a blank tile or a "0", and all adjacent cells will automatically be uncovered. Right-clicking on a cell will flag it, causing a flag to appear on it. Note that flagged cells are still covered, and a player can click on it to uncover it, like a normal covered cell.


In order to implement the solution we selected the following technologies:

- Golang 1.14.1

- Redis

- Docker

- dep  
     It is a dependency management tool for Go
- negroni library  
    It provide a web middleware functionality. 
- logrus library  
    It is a logger library for Go
- mux library  
    It implements a request router and dispatcher for matching incoming requests to their respective handler.

## Build local and standalone version
 
```
  $ dep ensure
  $ cd main
  $ go build main.go
```

## Run local and standalone version

```
   $ dep ensure
   $ cd main
   $ go build main.go
   $ ./main
```

## Run tests

```
  $ go test  $(go list ./... | grep -v /vendor/) -v
```


## Build Docker container

In order to generate a releasable artifact you need to run the following command:

```
    $ dep ensure
    $ cd scripts  
    $ ./build_images.sh
```

This command will generate a docker images ready to be executed called "minesweeper-api" 

## Run Docker container

Once you executed the steps described in the point called "Build steps", you need to created a new file called
"prod.env" using as template the filled called "example.env". In the file prod.env you need to 
provide the data related to the Redis server configuration. 
You need to execute the following steps

```
    $ cd script 
    $ cp example.env prod.env
    $ vi prod.env
    $ ./run_api.sh
```

## Create a New Game

```
  $ curl -i -X POST '127.0.0.1:8080/game' -d '{"name": "game1", "rows": 2, "cols": 3, "mines": 4}'
```

## Start the game

```
  $ curl -i -X POST '127.0.0.1:8080/game/game1/start'
```

## Play

```
  $ curl -i -X POST '127.0.0.1:8080/game/game1/click' -d '{"row": 1,"col":1}'
```
