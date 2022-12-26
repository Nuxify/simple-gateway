package main

// CastPluralToSingularForm convert common plural ending with s to singular equivalent for service routing
func CastPluralToSingularForm(pluralName string) string {
	// check if ends with `IES`, rename it with `Y`
	if pluralName[len(pluralName)-3:] == "IES" {
		pluralName = pluralName[:len(pluralName)-3]

		return pluralName + "Y"
	}

	// check if ends with `ES`, then remove ES
	if pluralName[len(pluralName)-2:] == "ES" {
		return pluralName[:len(pluralName)-2]
	}

	// check if ends with `SS`, return as is
	if pluralName[len(pluralName)-2:] == "SS" {
		return pluralName
	}

	// default, remove `S` in plural word
	return pluralName[:len(pluralName)-1]
}
