#!/bin/bash

curl -d "changeBackground $1" http://localhost:17000

startWinPositon=100
finishWinPositon=700

windowSize=800
figureSize=100

maxDistance=600
jmp=10
timeInterval=0.007

currPosX=$startWinPositon
currPosY=$startWinPositon

distanceX=0
distanceY=0

curl -d "figure $currPosX $currPosY" http://localhost:17000

sleep $timeInterval

while true; do
    curl -d "update" http://localhost:17000
    while (( currPosX < finishWinPositon )); do
        if (( distanceX >= maxDistance )); then
            currPosX=$finishWinPositon
            distanceX=$maxDistance
            distanceY=$(( distanceY + jmp ))
        else
            curl -d "move $jmp 0" http://localhost:17000
            currPosX=$(( currPosX + jmp ))
            distanceX=$(( distanceX + jmp ))
        fi
        curl -d "update" http://localhost:17000
        sleep $timeInterval
    done

    while (( currPosY < finishWinPositon )); do
        if (( distanceY >= maxDistance )); then
            currPosY=$finishWinPositon
            distanceY=$maxDistance
            distanceX=$(( distanceX - jmp ))
        else
            curl -d "move 0 $jmp" http://localhost:17000
            currPosY=$(( currPosY + jmp ))
            distanceY=$(( distanceY + jmp ))
        fi
        curl -d "update" http://localhost:17000
        sleep $timeInterval
    done

    while (( currPosX > startWinPositon )); do
        if (( distanceX <= 0 )); then
            currPosX=$startWinPositon
            distanceX=0
            distanceY=$(( distanceY - jmp ))
        else
            curl -d "move $((-jmp)) 0" http://localhost:17000
            currPosX=$(( currPosX - jmp ))
            distanceX=$(( distanceX - jmp ))
        fi
        curl -d "update" http://localhost:17000
        sleep $timeInterval
    done

    while (( currPosY > startWinPositon )); do
        if (( distanceY <= 0 )); then
            currPosY=$startWinPositon
            distanceY=0
            distanceX=$(( distanceX + jmp ))
        else
            curl -d "move 0 $((-jmp))" http://localhost:17000
            currPosY=$(( currPosY - jmp ))
            distanceY=$(( distanceY - jmp ))
        fi
        curl -d "update" http://localhost:17000
        sleep $timeInterval
    done
done