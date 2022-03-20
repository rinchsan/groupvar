package a

var _ int
var _, _, _ int
var _, _, _, _ int // want "grouped var declarations count should be up to 3"

var (
	_       int
	_, _    int
	_, _, _ int
)
var (
	_, _, _    int
	_, _, _, _ int // want "grouped var declarations count should be up to 3"
)

var _ = 1
var _, _, _ = 1, 2, 3
var _, _, _ = "foo", "bar", "baz"
var _, _, _, _ = 1, 2, 3, 4                    // want "grouped var declarations count should be up to 3"
var _, _, _, _ = "foo", "bar", "baz", "foobar" // want "grouped var declarations count should be up to 3"
var _, _, _, _ = 1, "foo", 2, "bar"            // want "grouped var declarations count should be up to 3"
var _, _ = 1, "foo"                            // want "grouped var declarations should be separated by types"

const _ = 1
const _, _, _ = 1, 2, 3
const _, _, _ = "foo", "bar", "baz"
const _, _, _, _ = 1, 2, 3, 4                    // want "grouped const declarations count should be up to 3"
const _, _, _, _ = "foo", "bar", "baz", "foobar" // want "grouped const declarations count should be up to 3"
const _, _, _, _ = 1, "foo", 2, "bar"            // want "grouped const declarations count should be up to 3"
const _, _ = 1, "foo"                            // want "grouped const declarations should be separated by types"

var (
	_, _, _ = 1, 2, 3
	_, _, _ = "foo", "bar", "baz"
)
var (
	_, _, _, _ = 1, 2, 3, 4                    // want "grouped var declarations count should be up to 3"
	_, _, _, _ = "foo", "bar", "baz", "foobar" // want "grouped var declarations count should be up to 3"
	_, _       = 1, "foo"                      // want "grouped var declarations should be separated by types"
)

const (
	_, _, _ = 1, 2, 3
	_, _, _ = "foo", "bar", "baz"
)
const (
	_, _, _, _ = 1, 2, 3, 4                    // want "grouped const declarations count should be up to 3"
	_, _, _, _ = "foo", "bar", "baz", "foobar" // want "grouped const declarations count should be up to 3"
	_, _       = 1, "foo"                      // want "grouped const declarations should be separated by types"
)

func f() {
	var _ int
	var _, _, _ int
	var _, _, _, _ int // want "grouped var declarations count should be up to 3"

	var (
		_       int
		_, _    int
		_, _, _ int
	)
	var (
		_, _, _    int
		_, _, _, _ int // want "grouped var declarations count should be up to 3"
	)

	var _ = 1
	var _, _, _ = 1, 2, 3
	var _, _, _ = "foo", "bar", "baz"
	var _, _, _, _ = 1, 2, 3, 4                    // want "grouped var declarations count should be up to 3"
	var _, _, _, _ = "foo", "bar", "baz", "foobar" // want "grouped var declarations count should be up to 3"
	var _, _, _, _ = 1, "foo", 2, "bar"            // want "grouped var declarations count should be up to 3"
	var _, _ = 1, "foo"                            // want "grouped var declarations should be separated by types"

	var (
		_, _, _ = 1, 2, 3
		_, _, _ = "foo", "bar", "baz"
	)
	var (
		_, _, _, _ = 1, 2, 3, 4                    // want "grouped var declarations count should be up to 3"
		_, _, _, _ = "foo", "bar", "baz", "foobar" // want "grouped var declarations count should be up to 3"
		_, _, _, _ = 1, "foo", 2, "bar"            // want "grouped var declarations count should be up to 3"
		_, _       = 1, "foo"                      // want "grouped var declarations should be separated by types"
	)
}
