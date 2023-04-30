#!/bin/bash
    curl -d "reset" http://localhost:17000
    curl -d "$1" http://localhost:17000 
    curl -d "update" http://localhost:17000 
