package theme

import "fmt"

var VerySmallPadding = CssUnit{6.4, Px} // = 0.4rem
var SmallPadding = CssUnit{12.8, Px}    // = 0.8rem
var NormalPadding = CssUnit{19.2, Px}   // = 1.2rem
var BigPadding = ContentBlockMargin

var ButtonPadding = SmallPadding
var InputPadding = SmallPadding
var ItemPadding = CssUnit{16, Px} // = 1rem

var TotalBlockItemPadding = ItemPadding.Mult(2)
var CompletePadding = fmt.Sprintf("%s %s %s %s", ButtonPadding.String(), "32px", ButtonPadding.String(), "32px")
