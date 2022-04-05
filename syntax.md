# syntax

The syntax is _roughly_ the same as Go. The major exceptions of course are
that the concurrency primitives are given their own keywords and structures.

A major _removal_ of syntax from Go is the `go` keyword (and the concept of
GoRoutines in general).

## actor block

```scala
// named actor (creatable via an ACTOR_NAME)
actor ACTOR_NAME {
    ... (actor body) ...
}

// anonymous actor (must be assigned or used)
actor {
    ... (actor body) ...
}
```


## receive methods

`receive` blocks are typed message receivers for an `actor`. An `actor`
may define multiple `receive` blocks. However, these are not named, so
`receive` statements cannot be duplicated for a given type.

```scala
receive (VAR_NAME VAR_TYPE) {
    ... (method body) ...
}
```

The culmination of all `receive` blocks determines the handling capabilities
of an actor. Unhandled message types can be explicitly handled and introspected
with an `unhandled` block.

```scala
unhandled (VAR_NAME TypeInfo, VAR_NAME interface{}) {
    ... (method body) ...
}
```

`TypeInfo` represents a descriptor for the unknown type. The value is provided
as an `interface{}` object.