package main

import "errors"

type BBrandFactory interface {
	makeShoe() BShoes
	makeShirt() BShirt
}

func GetBrandFactory(brand string) (BBrandFactory, error) {
	switch brand {
	case "puma":
		return PumaFactory{}, nil
	case "nike":
		return NikeFactory{}, nil
	}

	return nil, errors.New("invalid brand value")
}
