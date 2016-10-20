package freecompany

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/cking/golodestone"
	"github.com/cking/x/xhtml"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var (
	reID        = regexp.MustCompile(`(\d+)\/?$`)
	reAttribute = regexp.MustCompile(`:\s*(.*)$`)
	reAddress   = regexp.MustCompile(`Plot (\d+), (\d+) Ward, ([^(]+) \(([^)]+)\)`)
)

// FreeCompany Stores all the Character Data from Lodestone
type FreeCompany struct {
	Simple      bool
	ID          string
	Name        string
	World       string
	Realm       string
	URL         string
	Members     int
	Active      string
	Recruitment bool

	Tag           string
	Rank          int
	Slogan        string
	House         string
	HouseAddress  *HouseAddress
	HouseGreeting string
}

// HouseAddress Address/Place of a house
type HouseAddress struct {
	Plot  int
	Ward  int
	Place string
	Size  string
}

func (ha *HouseAddress) String() string {
	return fmt.Sprintf("Plot %v of Ward %v in %v (%v)", ha.Plot, ha.Ward, ha.Place, ha.Size)
}

func (c *FreeCompany) String() string {
	return c.Name + "#" + c.World + " <" + c.ID + ">"
}

// FetchCompany Fetch free company details for a company id
func FetchCompany(id string) (*FreeCompany, error) {
	root, err := golodestone.QueryLodestone(fmt.Sprintf("freecompany/%v", id))
	if err != nil {
		return nil, err
	}

	rawComp, ok := scrape.Find(root, scrape.ById("myfreecompany"))
	if !ok {
		return nil, errors.New("Company not found!")
	}

	comp := &FreeCompany{Simple: false, ID: id}
	comp.URL, _ = golodestone.BuildURL(fmt.Sprintf("freecompany/%v", id))

	subnode, _ := scrape.Find(rawComp, scrape.ByClass("ic_freecompany_box"))
	subnode = xhtml.MustNextSibling(xhtml.MustFirstChild(subnode))
	comp.Realm = strings.TrimSpace(subnode.FirstChild.Data) // only the first child, aka the realm. `scrape.Text` returns all text nodes in sub tags as well
	subnode = xhtml.MustFirstChild(subnode)
	comp.Name = scrape.Text(xhtml.MustNthSibling(subnode, 2))
	comp.World = scrape.Text(xhtml.MustNthSibling(subnode, 3))
	comp.World = comp.World[1 : len(comp.World)-1]

	subnode, _ = scrape.Find(rawComp, scrape.ByClass("vm"))
	comp.Tag = strings.TrimSpace(subnode.LastChild.Data)
	comp.Tag = comp.Tag[utf8.RuneLen('«') : len(comp.Tag)-utf8.RuneLen('»')]

	subnodes := scrape.FindAll(rawComp, scrape.ByTag(atom.Tr))
	comp.Members, _ = strconv.Atoi(scrape.Text(xhtml.MustNextSibling(xhtml.MustFirstChild(subnodes[2]))))
	comp.Rank, _ = strconv.Atoi(scrape.Text(xhtml.MustNextSibling(xhtml.MustFirstChild(subnodes[3]))))
	comp.Slogan = scrape.Text(xhtml.MustNextSibling(xhtml.MustFirstChild(subnodes[5])))
	comp.Active = scrape.Text(xhtml.MustNextSibling(xhtml.MustFirstChild(subnodes[8])))
	comp.Recruitment = scrape.Text(xhtml.MustNextSibling(xhtml.MustFirstChild(subnodes[9]))) == "Open"

	subnode, _ = scrape.Find(subnodes[len(subnodes)-1], scrape.ByClass("txt_yellow"))
	if subnode != nil {
		comp.House = scrape.Text(subnode)
		addr := reAddress.FindStringSubmatch(scrape.Text(xhtml.MustNthSibling(subnode, 2)))
		plot, _ := strconv.Atoi(addr[1])
		ward, _ := strconv.Atoi(addr[2])
		comp.HouseAddress = &HouseAddress{plot, ward, addr[3], addr[4]}
		comp.HouseGreeting = scrape.Text(xhtml.MustNthSibling(subnode, 4))
	}

	return comp, nil
}

// SearchCompany Search for a free company on the FFXIV Lodestone
func SearchCompany(companyName string) ([]*FreeCompany, error) {
	qs := &url.Values{}
	qs.Add("q", companyName)

	return searchCompany(qs.Encode())
}

// SearchCompanyByWorld Search for a free company on the FFXIV Lodestone, limited by the world
func SearchCompanyByWorld(world string, companyName string) ([]*FreeCompany, error) {
	qs := &url.Values{}
	qs.Add("q", companyName)
	qs.Add("worldname", strings.Title(world))

	return searchCompany(qs.Encode())
}

func searchCompany(params string) ([]*FreeCompany, error) {
	root, err := golodestone.QueryLodestone(fmt.Sprintf("freecompany/?%v", params))
	if err != nil {
		return nil, err
	}

	resultTable, ok := scrape.Find(root, scrape.ByClass("table_black_border_bottom"))
	if !ok {
		return nil, errors.New("no list")
	}

	rawComps := scrape.FindAll(resultTable, scrape.ByTag(atom.Tr))
	compList := make([]*FreeCompany, cap(rawComps))
	for idx, rawComp := range rawComps {
		comp := &FreeCompany{Simple: true}

		// find the character name
		subnode, _ := scrape.Find(rawComp, scrape.ByClass("gc_name_box"))
		subnode = xhtml.MustFirstChild(subnode)
		comp.Realm = scrape.Text(subnode)
		subnode = xhtml.MustFirstChild(xhtml.MustNextSibling(subnode))
		comp.Name = scrape.Text(subnode)
		comp.ID = reID.FindStringSubmatch(scrape.Attr(subnode, "href"))[1]
		comp.URL, _ = golodestone.BuildURL(scrape.Attr(subnode, "href"))

		// find the world name
		subnode = xhtml.MustNthSibling(subnode, 2)
		comp.World = scrape.Text(subnode)
		comp.World = comp.World[1 : len(comp.World)-1]

		// find the character name
		subnodes := scrape.FindAll(rawComp, func(n *html.Node) bool {
			return n.DataAtom == atom.Li && scrape.Attr(n.Parent.Parent.Parent, "class") == "list_box"
		})
		attrMem := reAttribute.FindStringSubmatch(scrape.Text(subnodes[0]))
		attrActive := reAttribute.FindStringSubmatch(scrape.Text(subnodes[1]))
		attrRecruitment := reAttribute.FindStringSubmatch(scrape.Text(subnodes[2]))
		comp.Members, _ = strconv.Atoi(attrMem[1])
		comp.Active = attrActive[1]
		comp.Recruitment = attrRecruitment[1] == "Open"

		compList[idx] = comp
	}

	return compList, nil
}
