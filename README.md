# Golang CLI Pomodoro Timer
## Introduction
The code in this repository is used to display a pomodoro timer in the terminal. The goal of this project was to familiarize myself with the fundamentals of the Go programming language. Throughout my studies in university, I heavily relied on the pomodoro study method often found in "study with me" videos on YouTube. When I thought of what could be a good project to learn go, I figured why not make my own? 

## Usage
First, clone the repo to your local machine, then you can run the binary included. This will start a pomodoro session with 5 pomodoros each with a length of 25 minutes with, 5 minute breaks for the first three, and a 15 minute break every fourth break. Alternatively, you can specify how many pomodoros, and their lengths with the optional flags. To change the number of pomodoros in a given session, use `-pomos=<num_pomos>`  to specify the desired number of pomodoros, and use `-length=<# minutes>` to specify the number of minutes each pomodoro should be (currently must be a multiple of 5).

## Notes
This is currently a work in progress, and improvements are made as they can be. Seeing that this is an educational project improvements will be made sparingly.
