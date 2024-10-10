package main

type PumaFactory struct{}

func (pf PumaFactory) makeShoe() BShoes {
	return PumaShoes{}
}

func (pf PumaFactory) makeShirt() BShirt {
	return PumaShirt{}
}
