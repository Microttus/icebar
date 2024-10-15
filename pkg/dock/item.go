package dock

type Item struct {
	Name string
	Exec string
	Icon string
}

func (item *Item) OnClick() error {
	//Launch application
	return nil
}
