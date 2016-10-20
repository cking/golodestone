package character

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/cking/golodestone"
	"github.com/cking/x/xhtml"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var (
	reID           = regexp.MustCompile(`(\d+)\/?$`)
	reRace         = regexp.MustCompile(`([^/]+) / ([^/]+) / ([♀♂])`)
	reGrandCompany = regexp.MustCompile(`([^/]+)/([^/]+)`)
)

// Character Stores all the Character Data from Lodestone
type Character struct {
	Simple          bool
	ID              string
	Name            string
	World           string
	FreeCompanyID   string
	FreeCompanyName string
	URL             string

	Avatar           string
	Portrait         string
	Title            string
	TitlePrefix      bool
	Race             string
	Clan             string
	Female           bool
	Gender           string
	Guardian         string
	GuardianIcon     string
	Nameday          string
	Realm            string
	RealmIcon        string
	GrandCompanyName string
	GrandCompanyRank string
	GrandCompanyIcon string
	CurrentClass     string

	Classes map[string]int
}

func (c *Character) String() string {
	return c.Name + "#" + c.World + " <" + c.ID + ">"
}

// FetchCharacter Fetch character details for a character id
func FetchCharacter(id string) (*Character, error) {
	root, err := golodestone.QueryLodestone(fmt.Sprintf("character/%v", id))
	if err != nil {
		return nil, err
	}

	rawChar, ok := scrape.Find(root, scrape.ById("character"))
	if !ok {
		return nil, errors.New("Character not found!")
	}

	char := &Character{Simple: false}

	subnode, _ := scrape.Find(rawChar, func(n *html.Node) bool {
		return n.DataAtom == atom.Img && scrape.Attr(n.Parent.Parent, "class") == "player_name_thumb"
	})
	char.Avatar = scrape.Attr(subnode, "src")

	subnode, _ = scrape.Find(rawChar, func(n *html.Node) bool {
		return n.DataAtom == atom.Img && scrape.Attr(n.Parent.Parent, "id") == "chara_img_area"
	})
	char.Portrait = scrape.Attr(subnode, "src")

	subnode, _ = scrape.Find(rawChar, func(n *html.Node) bool {
		return n.DataAtom == atom.A && n.Parent.DataAtom == atom.H2 && scrape.Attr(n.Parent.Parent, "class") == "area_footer player_name_txt"
	})
	char.ID = reID.FindStringSubmatch(scrape.Attr(subnode, "href"))[1]
	char.URL, _ = golodestone.BuildURL(scrape.Attr(subnode, "href"))
	char.Name = html.UnescapeString(scrape.Text(subnode))

	subnode, _ = scrape.Find(rawChar, func(n *html.Node) bool {
		return n.DataAtom == atom.Span && scrape.Attr(n.Parent.Parent, "class") == "area_footer player_name_txt"
	})
	char.World = scrape.Text(subnode)
	char.World = char.World[1 : len(char.World)-1]

	subnode, _ = scrape.Find(rawChar, scrape.ByClass("chara_title"))
	if subnode != nil {
		char.Title = scrape.Text(subnode)
		char.TitlePrefix = subnode == xhtml.MustFirstChild(subnode.Parent)
	}

	subnode, _ = scrape.Find(rawChar, scrape.ByClass("chara_profile_title"))
	race := reRace.FindStringSubmatch(scrape.Text(subnode))
	char.Race = race[1]
	char.Clan = race[2]
	char.Gender = race[3]
	char.Female = race[3] == "♀"

	subnodes := scrape.FindAll(rawChar, scrape.ByClass("chara_profile_box_info"))
	guardian := subnodes[0]
	subnode, _ = scrape.Find(guardian, scrape.ByClass("icon"))
	char.GuardianIcon = scrape.Attr(subnode.FirstChild, "src")
	subnode = xhtml.MustNthSibling(subnode, 2)
	char.Nameday = scrape.Text(subnode)
	subnode = xhtml.MustNthSibling(subnode, 2)
	char.Guardian = scrape.Text(subnode)

	for _, sne := range subnodes[1:] {
		typ, _ := scrape.Find(sne, scrape.ByClass("txt"))

		switch scrape.Text(typ) {
		case "City-state":
			char.Realm = scrape.Text(xhtml.MustNextSibling(typ))
			char.RealmIcon = scrape.Attr(xhtml.MustFirstChild(xhtml.MustNthSibling(typ, -1)), "src")

		case "Grand Company":
			gc := reGrandCompany.FindStringSubmatch(scrape.Text(xhtml.MustNextSibling(typ)))
			char.GrandCompanyName = gc[1]
			char.GrandCompanyRank = gc[2]
			char.GrandCompanyIcon = scrape.Attr(xhtml.MustFirstChild(xhtml.MustNthSibling(typ, -1)), "src")

		case "Free Company":
			fc := xhtml.MustFirstChild(xhtml.MustNextSibling(typ))
			char.FreeCompanyName = scrape.Text(fc)
			char.FreeCompanyID = reID.FindStringSubmatch(scrape.Attr(fc, "href"))[1]
		}
	}

	char.Classes = make(map[string]int)
	subnodes = scrape.FindAll(rawChar, func(n *html.Node) bool {
		return n.DataAtom == atom.Td && strings.Contains(scrape.Attr(n, "class"), "ic_class_wh24_box")
	})
	subnode, _ = scrape.Find(rawChar, func(n *html.Node) bool {
		return n.DataAtom == atom.Img && n.Parent.DataAtom == atom.Div && scrape.Attr(n.Parent, "class") == "ic_class_wh24_box"
	})
	for _, rawClass := range subnodes {
		if len(scrape.Text(rawClass)) == 0 {
			continue
		}
		level, _ := strconv.Atoi(scrape.Text(xhtml.MustNextSibling(rawClass)))
		char.Classes[strings.ToLower(scrape.Text(rawClass))] = level

		if scrape.Attr(subnode, "src") == scrape.Attr(xhtml.MustFirstChild(rawClass), "src") {
			char.CurrentClass = strings.ToLower(scrape.Text(rawClass))
		}
	}

	return char, nil
}

// SearchCharacter Search for a character on the FFXIV Lodestone
func SearchCharacter(characterName string) ([]*Character, error) {
	qs := &url.Values{}
	qs.Add("q", characterName)

	return searchCharacter(qs.Encode())
}

// SearchCharacterByWorld Search for a character on the FFXIV Lodestone, limited by the world
func SearchCharacterByWorld(world string, characterName string) ([]*Character, error) {
	qs := &url.Values{}
	qs.Add("q", characterName)
	qs.Add("worldname", strings.Title(world))

	return searchCharacter(qs.Encode())
}

// SearchCharacter Search for a character on the FFXIV Lodestone
func searchCharacter(params string) ([]*Character, error) {
	root, err := golodestone.QueryLodestone(fmt.Sprintf("character/?%v", params))
	if err != nil {
		return nil, err
	}

	resultTable, ok := scrape.Find(root, scrape.ByClass("table_black_border_bottom"))
	if !ok {
		return nil, errors.New("no list")
	}

	rawChars := scrape.FindAll(resultTable, scrape.ByTag(atom.Tr))
	charList := make([]*Character, cap(rawChars))
	for idx, rawChar := range rawChars {
		char := &Character{Simple: true}

		// find the character name
		subnode, _ := scrape.Find(rawChar, func(n *html.Node) bool {
			return n.DataAtom == atom.A && n.Parent.DataAtom == atom.H4
		})
		char.ID = reID.FindStringSubmatch(scrape.Attr(subnode, "href"))[1]
		char.URL, _ = golodestone.BuildURL(scrape.Attr(subnode, "href"))
		char.Name = html.UnescapeString(scrape.Text(subnode))

		// find the world name
		subnode, _ = scrape.Find(rawChar, func(n *html.Node) bool {
			return n.DataAtom == atom.Span && n.Parent.DataAtom == atom.H4
		})
		char.World = scrape.Text(subnode)
		char.World = char.World[1 : len(char.World)-1]

		// find the character name
		subnode, _ = scrape.Find(rawChar, func(n *html.Node) bool {
			return n.DataAtom == atom.A && n.Parent.DataAtom == atom.Div && scrape.Attr(n.Parent, "class") == "txt_gc"
		})
		if subnode != nil {
			char.FreeCompanyID = reID.FindStringSubmatch(scrape.Attr(subnode, "href"))[1]
			char.FreeCompanyName = scrape.Text(subnode)
		}

		charList[idx] = char
	}

	return charList, nil
}
