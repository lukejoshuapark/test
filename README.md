![](icon.png)

# test

An easy-to-use testing module with pleasing syntax and support for custom
assertions.

## Usage Example

```go
func TestTheFabricOfTheUniverse(t *testing.T) {
	test.That(t, 2 + 2, is.EqualTo(4))
}
```

## Assertions

The core function in this module, `That`, accepts an `Assertion` as its last
argument.  `Assertion` is simply an interface, so can be implemented to provide
custom assertions for specific use cases.

The `does`, `has` and `is` packages contained in this module provide a number of
useful assertions.  Some of the more commonly used assertions are:

- `is.EqualTo`
- `is.Nil`
- `is.True`
- `is.False`
- `has.Length`
- `does.Panic`

You can also create assertions that use other assertions internally.  For
example, the `does.PanicWithValueThat` assertion:

```go
f := func () { panic("ahh!") }
test.That(t, f, does.PanicWithValueThat(is.EqualTo("ahh!")))
```

or the `has.ValueThat` assertion:

```go
m := map[string]int{"a": 42}
test.That(t, m, has.ValueThat(is.EqualTo(42)))
```

## Nil and Interfaces

A well known "gotcha" in the Go language is the comparison behavior when
comparing `nil` with an interface that has a type but no value.

For example:

```go
fmt.Printf("%v", io.Writer((*os.File)(nil)) == nil) // false
```

Which is somewhat unintuitive.  This package respects that behavior for all
comparison-based assertions **except** for `is.Nil` and `is.NotNil` which
consider an interface, map, slice, func or chan with a type, but a nil value, to
still be nil.

For example, this test will pass:

```go
test.That(t, io.Writer((*os.File)(nil)), is.Nil)
```

but this test will fail:

```go
test.That(t, io.Writer((*os.File)(nil)), is.EqualTo(nil))
```

---

Icons made by [iconixar](https://www.flaticon.com/authors/iconixar) from
[www.flaticon.com](https://www.flaticon.com/).
