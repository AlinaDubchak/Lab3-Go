# Lab 3 Go
## Topic: 
- Event Loop

## Purpose: 
- Acquiring skills of system implementation using event loop. Its usage in UI frameworks

## This project create by:

> [Мартинюк Марія](https://github.com/mmarty12) <br>
> [Котенко Ярослав](https://github.com/yarikkot04) <br>
> [Дубчак Аліна](https://github.com/AlinaDubchak) <br>

## Description
This project contains the implementation of minimal GUI using event loop, realisation of another event loop to control the image that was drawn by the previous component, 
the digram of dependencies, and bash scripts. Files loop.go and parser.go are covered with tests (see folder [test](https://github.com/AlinaDubchak/Lab3-Go/tree/main/test))

## Description of bash scripts
**1. changeBackgroundColor**

This script fills the window with a certain color.
```
changeBackgroundColor.sh black   //fill with black
changeBackgroundColor.sh white   //fill with white
changeBackgroundColor.sh green   //fill with green
```

**2. createDiagonal**

By running this script the cross will start moving from top left to bottom right conner and backwards.
```
createDiagonal.sh 0   //black background
createDiagonal.sh 1   //white background
createDiagonal.sh 2   //green background
```
**3. createFigure**

This script creates a figure on a background you choose.
```
createFigure.sh black     //black background
createFigure.sh white     //white background
createFigure.sh green     //green background
```
**4. createGreenFrame**

This script creates a green frame on a background you choose.
```
createGreenFrame.sh 0   //black background
createGreenFrame.sh 1   //white backgroun
createGreenFrame.sh 2   //green backgroun
```

**5. createSquare**

By running this script the cross will start moving along the frame of the window creating a square.
```
createSquare.sh 0   //black background
createSquare.sh 1   //white background
createSquare.sh 2   //green background
```

**6. resetProgram**

This script resets the window and fills it with certain color.
```
resetProgram.sh  OR  resetProgram.sh black   //black background
resetProgram.sh white   //white background
resetProgram.sh green   //green background
```

## How to run project

1. Clone the repository to your local machine

2. Navigate to the project directory in your terminal or command prompt

- To run the project
```
go run cmd/painter/main.go 
```
(you will see a window with black background and yellow cross, that you can move to different position by clicking the left mouse button)

- To run bash scripts

In first terminal run the main file:
```
go run cmd/painter/main.go 
```
Open the second terminal and move to 'scripts' folder:
```
cd scripts
```
Next run the bashscript by simply entering its name and additional argument (see **Description of bash scripts**). 

---
## References

[Github Actions](https://github.com/AlinaDubchak/Lab3-Go/actions)

[Dependancy Diagram](https://github.com/AlinaDubchak/Lab3-Go/blob/main/components.pdf)
