package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Item struct {
	Text     string
	Priority int
	position int
	Done     bool
}

type ByPri []Item

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}

}

func (i *Item) Ipretty() string {
	if i.Priority == 1 {
		return "[1]"
	}
	if i.Priority == 2 {
		return "[2]"
	}
	if i.Priority == 3 {
		return "[3]"
	}

	return ""
}

func (i *Item) Lable() string {
	return strconv.Itoa(i.position) + "."
}

func (i *Item) PrettyDone() string {
	if i.Done {
		return "X"
	}
	return ""
}

func (s ByPri) Len() int {
	return len(s)
}

func (s ByPri) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByPri) Less(i, j int) bool {
	if s[i].Done != s[j].Done {
		return s[i].Done
	}
	if s[i].Priority != s[j].Priority {
		return s[i].Priority < s[j].Priority
	}

	return s[i].position < s[j].position
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}
	fmt.Println(string(b))

	err = ioutil.WriteFile(filename, b, 644)
	if err != nil {
		return err
	}
	return nil
}

func ReadItems(filename string) ([]Item, error) {
	items := []Item{}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, nil
	}

	for i, _ := range items {
		items[i].position = i + 1
	}
	return items, nil
}
