# go-type more!

A TUI typing test built with Bubbletea & Lipgloss in golang

## TODO:
	- Begin stats:
		- wpm
		- accuracy
		- ntcharts: graphs based on speed throughout? total_time / 4 for quartiles of the graph
	- Once all characters are typed, go to stats screen
	- add tabs for stats - after finishing sentence, give complete message and move to stats page
	- maybe a table of your own ranking and stats from the last several attempts that are stored locally
	- different types of quotes chosen by the user
	- make tabs center of page and a different color
	- make instructions on how to use program

## COMPLETED:
	- if character matches, change color of quote chracter to green, otherwise red
	- check all characters entered against the chosen quote
	- random quotes
	- tabs for diff pages
	- button to reset the quote and do it again

## Progression:
1. Initialized a bubble tea program
	- make it a full sized application
	- put a box on the screen with my content
2. Add some styles using lipgloss, center in screen, get terminal size
3. Get user input and show it on the screen
4. Compare user input to a target string and change colors based on correctness
5. add different tabs for stats, typing test, and info/about
6. game functionality
	- once a user finishes typing, go to stats page
	- calculate stats, wpm, accuracy, etc.
	- show charts with ntcharts
	- clean up some of the styles

## Resources:
https://github.com/NimbleMarkets/ntcharts?tab=readme-ov-file
https://github.com/charmbracelet/bubbletea/blob/main/examples/tabs/main.go
https://github.com/charmbracelet/bubbletea/blob/main/examples/textinput/main.go
https://github.com/charmbracelet/lipgloss?tab=readme-ov-file

Inspiration:
https://github.com/lemnos/tt
