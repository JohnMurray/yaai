# yaai

yet another actor implementation


## things

__shared objects__

User code built as a shared object. The runtime needs to read this shared
object and run some bits on it.

The Go stdlib has a [plugin](https://pkg.go.dev/plugin) package that does
some dlopen bits and runs a package's `init()` function. It doesn't support
_unloading_, but I think that modifying this package to add unloading could
definitely work. Using an init-based model might mean we can do more with
standard Go in the runtime.

__serialization__

All structs in user-code need to automatically be backed by something like
proto's for easy serialization.

## more things

we ought to have:

  + automatic serialization support for all sendable/receivable messages
  + elimination of concurrency primitives outside of actors
  + plugin hot-reloading 
  + automatic clustering
  + automatic actor ditribution across the cluster
  + seamless actor upgrades (on plugin update)
  + no "main" function, only actors
  + "plugins" should be isolated to individual actors and hot-reloading
    should be a per-actor thing
  + updates to plugins should propegate across the cluster

## compilation

__build phases__

(not in order)

0. AST parsing
1. proto generation for sendable messages (maybe don't need this for all structs)
2. interface generation
  + maybe some type of self-regitration system
  + maybe a generic handler that needs to be exported via C FFI
  + essentially this needs to offer a place to wire into the runtime and
    provide some way to query meta-data about actors

__build system__

It may be interesting to not only have a build _system_ that uses bazel,
but to actually use bazel to build an incremental compiler. All of the
caching that bazel does has the potential to make for a really fast
compiler runtime. :thinkies:

## target language / runtime

__go__

it has a great library and becomes pretty useful right out of the gate.
I just need to expose/wrap a bunch of stuff, which I think should be
pretty minimal. re-use as much as I can while avoiding or abstracting over
the stuff I don't want (channels, goroutines, etc).

Can transpile yaai over to Go by mimic'ing most of the syntax, which should
make life a bit easier.

Don't need to support things like generics. The 1.x backwards compatibility
should make code-gen rather stable.

ABI stability might be a problem. Probably gonna have to rely on C-Go style
things to make interop easier. Since we're not allowing any sort of
concurrency in yaai, I think that's probably not a big concern. So ABI
_could_ not be a problem.

__c / c++__

the library support is crap, and so I'd have to spend a lot of time
implementing things for it to be useful. Or I'd have to settle for
dumb examples doing very little. sounds boring.

the ABI bits would be _super_ stable, and so shared-object loading
would be really stable.

Since I wouldn't simply be transpiling, I'd have to care more about
actual compiler things. I don't want to really think about this.