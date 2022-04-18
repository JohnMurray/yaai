package test

// func TestParsePackageDecl(t *testing.T) {
// 	testInputs := [][]string{
// 		{"package Stonks\n", "Stonks"},             // title-case
// 		{"package lowercase\n", "lowercase"},       // lowercase
// 		{"package thing1thing2\n", "thing1thing2"}, // with numbers
// 		{"package thing\r", "thing"},               // \r line-ending
// 		{"package thing\n\n\r\n\r\n\n", "thing"},   // multiple line-endings
// 	}
// 	for _, test := range testInputs {
// 		t.Run(test[0], func(t *testing.T) {
// 			p, tl := createParser(test[0])
// 			antlr.ParseTreeWalkerDefault.Walk(tl, p.PackageClause())
// 			assert.Equal(t, test[1], tl.packageName)
// 			assert.Empty(t, tl.errors)
// 		})
// 	}
// }
//
// func TestParseBadPackageDecl(t *testing.T) {
// 	testInputs := []string{
// 		"package Stonks",         // no newline
// 		"stuff package Stonks\n", // leading crap
// 		"package 9things\n",      // can't start with number
// 		"packageThing\n",         // no space
// 		"Package Stonks\n",       // uppercase package
// 	}
// 	for _, test := range testInputs {
// 		t.Run(test, func(t *testing.T) {
// 			p, tl := createParser(test)
// 			antlr.ParseTreeWalkerDefault.Walk(tl, p.PackageClause())
// 			assert.NotEmpty(t, tl.errors)
// 		})
// 	}
// }
