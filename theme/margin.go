package theme

import "fmt"

var ContentInlineMargin = CssUnit{20, Px}
var ContentTotalInlineMargin = ContentInlineMargin.Mult(2)

var ContentVeryVerySmallBlockMargin = CssUnit{8, Px}
var ContentVerySmallBlockMargin = CssUnit{20, Px}
var ContentSmallBlockMargin = CssUnit{25, Px}
var ContentBlockMargin = CssUnit{30, Px}
var ContentTotalBlockMargin = ContentBlockMargin.Mult(2)

var ContentTopMargin = ContentBlockMargin.Add(10)
var ContentRightMargin = ContentInlineMargin.Add(30)
var ContentBottomMargin = ContentBlockMargin.Add(10)
var ContentLeftMargin = ContentInlineMargin.Add(30)
var ContentMargin = fmt.Sprintf("%s %s %s %s", ContentTopMargin.String(), ContentRightMargin.String(), ContentBottomMargin.String(), ContentLeftMargin.String())

var ContentTopSafeMargin = ContentBlockMargin
var ContentRightSafeMargin = ContentRightMargin
var ContentBottomSafeMargin = ContentBlockMargin
var ContentLeftSafeMargin = ContentLeftMargin
var ContentSafeMargin = fmt.Sprintf("%s %s %s %s", ContentTopSafeMargin.String(), ContentRightSafeMargin.String(), ContentBottomSafeMargin.String(), ContentLeftSafeMargin.String())

var ContentMainMargin = fmt.Sprintf("%s %s %s %s", ContentTopMargin.String(), ContentRightSafeMargin.String(), ContentBottomMargin.String(), ContentLeftSafeMargin.String())
var ContentMainBottomSafeMargin = fmt.Sprintf("%s %s %s %s", ContentTopMargin.String(), ContentRightSafeMargin.String(), ContentBottomSafeMargin.String(), ContentLeftSafeMargin.String())

var ContentTopBottomMargin = ContentTopMargin.AddUnit(ContentBottomMargin, WithUnit(Px))
var ContentTopBottomSafeMargin = ContentTopSafeMargin.AddUnit(ContentBottomSafeMargin)
