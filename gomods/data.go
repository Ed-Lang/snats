package snats

import (
	"fyne.io/fyne/v2"
)

// Tutorial defines the data structure for a tutorial
type Tutorial struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
	SupportWeb   bool
}

var (
	// Tutorials defines the metadata for each tutorial
	Tutorials = map[string]Tutorial{
		"welcome": {"Welcome", "", welcomeScreen, true},
		"canvas": {"Canvas",
			"See the canvas capabilities.",
			canvasScreen,
			true,
		},
		"animations": {"Animations",
			"See how to animate components.",
			makeAnimationScreen,
			true,
		},
		"icons": {"Theme Icons",
			"Browse the embedded icons.",
			iconScreen,
			true,
		},
		"containers": {"Containers",
			"Containers group other widgets and canvas objects, organising according to their layout.\n" +
				"Standard containers are illustrated in this section, but developers can also provide custom " +
				"layouts using the fyne.NewContainerWithLayout() constructor.",
			containerScreen,
			true,
		},
		"apptabs": {"AppTabs",
			"A container to help divide up an application into functional areas.",
			makeAppTabsTab,
			true,
		},
		"border": {"Border",
			"A container that positions items around a central content.",
			makeBorderLayout,
			true,
		},
		"box": {"Box",
			"A container arranges items in horizontal or vertical list.",
			makeBoxLayout,
			true,
		},
		"center": {"Center",
			"A container to that centers child elements.",
			makeCenterLayout,
			true,
		},
		"doctabs": {"DocTabs",
			"A container to display a single document from a set of many.",
			makeDocTabsTab,
			true,
		},
		"grid": {"Grid",
			"A container that arranges all items in a grid.",
			makeGridLayout,
			true,
		},
		"split": {"Split",
			"A split container divides the container in two pieces that the user can resize.",
			makeSplitTab,
			true,
		},
		"scroll": {"Scroll",
			"A container that provides scrolling for it's content.",
			makeScrollTab,
			true,
		},
		"widgets": {"Widgets",
			"In this section you can see the features available in the toolkit widget set.\n" +
				"Expand the tree on the left to browse the individual tutorial elements.",
			widgetScreen,
			true,
		},
		"accordion": {"Accordion",
			"Expand or collapse content panels.",
			makeAccordionTab,
			true,
		},
		"button": {"Button",
			"Simple widget for user tap handling.",
			makeButtonTab,
			true,
		},
		"card": {"Card",
			"Group content and widgets.",
			makeCardTab,
			true,
		},
		"entry": {"Entry",
			"Different ways to use the entry widget.",
			makeEntryTab,
			true,
		},
		"form": {"Form",
			"Gathering input widgets for data submission.",
			makeFormTab,
			true,
		},
		"input": {"Input",
			"A collection of widgets for user input.",
			makeInputTab,
			true,
		},
		"text": {"Text",
			"Text handling widgets.",
			makeTextTab,
			true,
		},
		"toolbar": {"Toolbar",
			"A row of shortcut icons for common tasks.",
			makeToolbarTab,
			true,
		},
		"progress": {"Progress",
			"Show duration or the need to wait for a task.",
			makeProgressTab,
			true,
		},
		"collections": {"Collections",
			"Collection widgets provide an efficient way to present lots of content.\n" +
				"The List, Table, and Tree provide a cache and re-use mechanism that make it possible to scroll through thousands of elements.\n" +
				"Use this for large data sets or for collections that can expand as users scroll.",
			collectionScreen,
			true,
		},
		"list": {"List",
			"A vertical arrangement of cached elements with the same styling.",
			makeListTab,
			true,
		},
		"table": {"Table",
			"A two dimensional cached collection of cells.",
			makeTableTab,
			true,
		},
		"tree": {"Tree",
			"A tree based arrangement of cached elements with the same styling.",
			makeTreeTab,
			true,
		},
		"dialogs": {"Dialogs",
			"Work with dialogs.",
			dialogScreen,
			true,
		},
		"windows": {"Windows",
			"Window function demo.",
			windowScreen,
			false,
		},
		"binding": {"Data Binding",
			"Connecting widgets to a data source.",
			bindingScreen,
			true,
		},
		"advanced": {"Advanced",
			"Debug and advanced information.",
			advancedScreen,
			true,
		},
	}

	// TutorialIndex  defines how our tutorials should be laid out in the index tree
	TutorialIndex = map[string][]string{
		"":            {"welcome", "canvas", "animations", "icons", "widgets", "collections", "containers", "dialogs", "windows", "binding", "advanced"},
		"collections": {"list", "table", "tree"},
		"containers":  {"apptabs", "border", "box", "center", "doctabs", "grid", "scroll", "split"},
		"widgets":     {"accordion", "button", "card", "entry", "form", "input", "progress", "text", "toolbar"},
	}
)
