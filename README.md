# go-type more!

A blazingly fast TUI typing test built with Bubbletea & Lipgloss in golang to encourage you to type more!

## TODO:
	- Begin stats:
		- ntcharts: graphs based on speed throughout? total_time / 4 for quartiles of the graph
		- wpm
	- table of ranking and stats

## COMPLETED:
	- add author for credit
	- different sizes of strings
	- make instructions on how to use program
	- if character matches, change color of quote chracter to green, otherwise red
	- check all characters entered against the chosen quote
	- random quotes
	- tabs for diff pages
	- button to reset the quote and do it again
	- Once all characters are typed, go to stats screen
	- make tabs center of page and a different color
	- add tabs for stats - after finishing sentence

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


Day 1:
	- init project, got core functionality down
Day 2:
	- styles, got loop of gameplay done
	- set up github and fixed project structure
Day 3:
	- random quotes
	- random words

## Resources:
https://github.com/NimbleMarkets/ntcharts?tab=readme-ov-file
https://github.com/charmbracelet/bubbletea/blob/main/examples/tabs/main.go
https://github.com/charmbracelet/bubbletea/blob/main/examples/textinput/main.go
https://github.com/charmbracelet/lipgloss?tab=readme-ov-file

Inspiration:
https://github.com/lemnos/tt
