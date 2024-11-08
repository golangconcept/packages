# GoConvey

`GoConvey` is a testing framework for Go (Golang) that makes writing and running tests more intuitive and expressive.

It provides a rich set of features to help with writing tests in a behavior-driven development (`BDD`) style, where you can structure tests with clear and descriptive language.

GoConvey also includes a `web-based UI` to view test results, making it a useful tool for both test automation and interactive debugging.

### GoConvey Assertions

GoConvey includes a wide variety of assertions to check values, such as:

- `ShouldEqual`: Checks if two values are equal.
- `ShouldNotEqual`: Checks if two values are not equal.
- `ShouldBeNil`: Checks if a value is nil.
- `ShouldNotBeNil`: Checks if a value is not nil.
- `ShouldBeTrue`: Checks if a boolean value is true.
- `ShouldBeFalse`: Checks if a boolean value is false.
- `ShouldContain`: Checks if a string contains another string or a slice contains an element.
- `ShouldMatch`: Checks if a string matches a regular expression.
- `ShouldPanic`: Asserts that a function call results in a panic.
