package epick

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/gdamore/tcell"
	"gitlab.com/tslocum/cview"
)

const cellText string = "%.4s[black]____[white]%s "

// EmojiValues type used to hold emoji value when returned
type EmojiValues struct {
	Emoji       string
	Description string
}

var returnValues EmojiValues

var emojiData rawEmojiData

var app *cview.Application = cview.NewApplication()

var mainFlex *cview.Flex = cview.NewFlex()

var tables []*cview.Table

var pages *cview.Pages = cview.NewPages()
var pageIndex int

var searchInput *cview.InputField = cview.NewInputField()
var searchValues []string
var searchIndexes [][]int
var searchIndex int

// Application setup -----------------------------------------------------

func appSetup() {
	// Get the emoji data
	err := json.Unmarshal([]byte(emojiJson), &emojiData)
	if err != nil {
		log.Fatal(err)
	}

	// For each group of emojis, make a table based off of that group
	for i, group := range emojiData.LIST {
		flex := cview.NewFlex()

		pageSetup(flex, group)

		pages.AddPage(fmt.Sprintf("page-%d", i), flex, true, false)
	}

	// Switch to the first page
	pages.SwitchToPage("page-0")

	mainFlex.AddItem(pages, 0, 1, true)

	searchInputSetup()
}

// Set up an individual page
func pageSetup(flex *cview.Flex, emojiData emojiGroup) {
	flex.SetDirection(cview.FlexRow)

	title := cview.NewTextView()
	title.SetTextAlign(cview.AlignCenter)
	title.SetText(emojiData.NAME)
	flex.AddItem(title, 0, 1, false)

	table := cview.NewTable()
	table.Select(0, 0)
	table.SetSelectable(true, true)
	table.SetSelectedStyle(-1, tcell.ColorGray, tcell.AttrNone)
	table.SetSelectedFunc(tableSelectionFunc)
	tables = append(tables, table)

	// Loop through every emoji and assign it a table cell
	var x, y int
	for _, v := range emojiData.EMOJIS {
		cell := cview.NewTableCell(fmt.Sprintf(cellText, v.EMOJI, v.NAME))
		table.SetCell(x*2, y*2, cell)

		x++
		if x == 16 {
			x = 0
			y++
		}
	}

	flex.AddItem(table, 0, 10, true)
}

func searchInputSetup() {
	for _, group := range emojiData.LIST {
		for _, e := range group.EMOJIS {
			searchValues = append(searchValues, strings.ToLower(e.NAME))
		}
	}

	searchInput.SetLabel("Enter text to search for an emoji: ")
	searchInput.SetFieldWidth(120)

	searchInput.SetDoneFunc(searchInputDoneFunc)
	searchInput.SetAutocompleteFunc(searchInputAutocompleteFunc)
}

// Search Input functions ------------------------------------------------

func searchInputDoneFunc(key tcell.Key) {
	switch key {
	case tcell.KeyEscape:
		app.Stop()

	case tcell.KeyEnter:
		text := strings.ToLower(searchInput.GetText())

		// Get the possible locations of the search
		locations := getSearchLocations(text)
		searchIndexes = locations

		// Remove the search prompt
		mainFlex.RemoveItem(searchInput)
		app.SetFocus(pages)

		// If there are locations, go to the first instance
		if len(locations) > 0 {
			i := locations[0][0]
			row := locations[0][1]
			col := locations[0][2]

			pages.SwitchToPage(fmt.Sprintf("page-%d", i))
			pageIndex = i

			tables[i].Select(col, row)
		}
	}
}

func searchInputAutocompleteFunc(currentText string) []string {
	if len(currentText) == 0 {
		return nil
	}

	// Create a list of possible selections
	var entries []string
	for _, word := range searchValues {
		if strings.HasPrefix(strings.ToLower(word), strings.ToLower(currentText)) {
			entries = append(entries, word)
		}
	}

	// For better formatting/looks
	if len(entries) < 1 {
		entries = nil
	} else if len(entries) == 1 && entries[0] == currentText {
		entries = nil
	}

	return entries
}

// Get all cells from tables that contain the text string
func getSearchLocations(text string) [][]int {
	var locations [][]int
	for i, group := range emojiData.LIST {
		for x := 0; x < int(math.Ceil(float64(len(group.EMOJIS)/16))); x++ {
			for y := 0; y < 16; y++ {
				desc := strings.ToLower(group.EMOJIS[x*16+y].NAME)
				if strings.Contains(desc, text) {
					var location = []int{i, x * 2, y * 2}

					locations = append(locations, location)
				}
			}
		}
	}

	return locations
}

func inputHandler(event *tcell.EventKey) *tcell.EventKey {
	if searchInput.HasFocus() {
		return event
	}

	switch {
	case event.Rune() == 'q' || event.Key() == tcell.KeyEscape:
		app.Stop()

	case event.Rune() == '`':
		//showHelp()

	case event.Rune() == '?':
		mainFlex.AddItem(searchInput, 0, 1000, true)
		searchInput.SetText("")
		app.SetFocus(searchInput)

	// Switch pages
	case event.Rune() == 'C':
		if pageIndex < len(tables)-1 {
			pageIndex++

			pages.SwitchToPage(fmt.Sprintf("page-%d", pageIndex))
		}

	case event.Rune() == 'c':
		if pageIndex > 0 {
			pageIndex--

			pages.SwitchToPage(fmt.Sprintf("page-%d", pageIndex))
		}

	case event.Rune() == 'N':
		if len(searchIndexes) <= 1 {
			return nil
		}

		if searchIndex == 0 {
			searchIndex = len(searchIndexes) - 1
		}
		searchIndex--

		pages.SwitchToPage(fmt.Sprintf("page-%d", searchIndexes[searchIndex][0]))
		tables[searchIndexes[searchIndex][0]].Select(searchIndexes[searchIndex][1], searchIndexes[searchIndex][2])

	case event.Rune() <= 'n':
		if len(searchIndexes) <= 1 {
			return nil
		}

		if searchIndex == len(searchIndexes)-1 {
			searchIndex = 0
		}
		searchIndex++

		pages.SwitchToPage(fmt.Sprintf("page-%d", searchIndexes[searchIndex][0]))
		tables[searchIndexes[searchIndex][0]].Select(searchIndexes[searchIndex][2], searchIndexes[searchIndex][1])
	}

	event = movementHandler(event)

	return event
}

func movementHandler(event *tcell.EventKey) *tcell.EventKey {
	switch {
	case event.Rune() == 'l' || event.Key() == tcell.KeyRight:
		row, col := tables[pageIndex].GetSelection()
		if col < tables[pageIndex].GetColumnCount()-2 && tables[pageIndex].GetCell(row, col+2).Text != "" {
			tables[pageIndex].Select(row, col+2)
		}
		return nil

	case event.Rune() == 'h' || event.Key() == tcell.KeyRight:
		row, col := tables[pageIndex].GetSelection()
		if col > 0 {
			tables[pageIndex].Select(row, col-2)
		}
		return nil

	case event.Rune() == 'k' || event.Key() == tcell.KeyRight:
		row, col := tables[pageIndex].GetSelection()
		if row > 0 {
			tables[pageIndex].Select(row-2, col)
		}
		return nil

	case event.Rune() == 'j' || event.Key() == tcell.KeyRight:
		row, col := tables[pageIndex].GetSelection()
		if row < tables[pageIndex].GetRowCount()-2 && tables[pageIndex].GetCell(row+2, col).Text != "" {
			tables[pageIndex].Select(row+2, col)
		}
		return nil

	case event.Rune() == 'G':
		row := tables[pageIndex].GetRowCount() - 1
		col := tables[pageIndex].GetColumnCount() - 1
		for true {
			cell := tables[pageIndex].GetCell(row, col)
			if cell.Text != "" {
				break
			} else {
				row -= 2
			}
		}
		tables[pageIndex].Select(row, col)
		return nil
	}

	return event
}

// Function that runs when an emoji is selected
func tableSelectionFunc(row int, column int) {
	cell := tables[pageIndex].GetCell(row, column)

	text := strings.Split(cell.Text, "[black]____[white]")

	returnValues = EmojiValues{text[0], text[1]}

	app.Stop()
}

// Start function starts the epick application
func Start(testing bool) (EmojiValues, error) {
	appSetup()

	app.SetInputCapture(inputHandler)

	if err := app.SetRoot(mainFlex, true).Run(); err != nil {
		log.Fatal(err)
		panic(err)
	}

	return returnValues, nil
}
