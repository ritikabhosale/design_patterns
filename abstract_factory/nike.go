package main

type NikeFactory struct {
}

func (nf NikeFactory) makeShoe() BShoes {
	return NikeShoes{}
}

func (nf NikeFactory) makeShirt() BShirt {
	return NikeShirt{}
}
