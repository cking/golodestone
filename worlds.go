package golodestone

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"

	"github.com/cking/x/xhtml"
	"github.com/yhat/scrape"
)

// DataCenterMap Map of all data centers
type DataCenterMap map[string]WorldStatusMap

// WorldStatusMap Map of all worlds
type WorldStatusMap map[string]StatusFlag

// StatusFlag Status flag of a World
type StatusFlag int

const (
	// StatusOnline The world is online
	StatusOnline = iota
	// StatusMaintenance The world is under maintenance and may have connection issues
	StatusMaintenance = iota
	// StatusOffline The world is offline
	StatusOffline = iota
)

var (
	// DataCenter Map of all Worlds and their status separated by DataCenters
	DataCenter DataCenterMap
	// Worlds Map of all Worlds and their status
	Worlds WorldStatusMap

	reWorldStatus = regexp.MustCompile(`worldstatus_(\d)`)
)

func init() {
	DataCenter, _ = WorldStatus()
}

// WorldStatus Query the world status
func WorldStatus() (DataCenterMap, error) {
	if DataCenter == nil {
		DataCenter = make(DataCenterMap)
	}

	node, err := QueryLodestone("worldstatus")
	if err != nil {
		return nil, err
	}

	elDatacenters := scrape.FindAll(node, scrape.ByClass("text-headline"))
	for _, elDc := range elDatacenters {
		dc := strings.ToLower(strings.TrimSpace(elDc.LastChild.Data))
		DataCenter[dc] = make(WorldStatusMap)
		worldNode, _ := scrape.FindParent(elDc, scrape.ByClass("area_body"))
		worldNode = xhtml.MustNextSibling(worldNode)

		worlds := scrape.FindAll(worldNode, func(n *html.Node) bool {
			return strings.Contains(scrape.Attr(n, "class"), "worldstatus_")
		})
		for _, w := range worlds {
			classList := scrape.Attr(w, "class")
			statusFlag, _ := strconv.Atoi(reWorldStatus.FindStringSubmatch(classList)[1])
			worldName, _ := xhtml.FirstChild(w)
			DataCenter[dc][strings.ToLower(scrape.Text(worldName))] = StatusFlag(statusFlag)
		}
	}

	Worlds = regenerateWorldMap()
	return DataCenter, nil
}

func regenerateWorldMap() WorldStatusMap {
	worlds := make(WorldStatusMap)
	for _, dc := range DataCenter {
		for name, status := range dc {
			worlds[name] = status
		}
	}

	return worlds
}

// IsValidWorld Checks if name is a valid world
func IsValidWorld(name string) bool {
	name = strings.ToLower(name)

	for world := range Worlds {
		if name == world {
			return true
		}
	}

	return false
}

// GuessWorldName Try to autocomplete world name
func GuessWorldName(name string) (string, error) {
	var guesses []string
	name = strings.ToLower(name)

	for world := range Worlds {
		if strings.Contains(world, name) {
			guesses = append(guesses, world)
		}
	}

	if len(guesses) == 1 {
		return guesses[0], nil
	}

	return "", errors.New("World not found")
}
