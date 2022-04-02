# yaai

yet another actor implementation


## things

__shared objects__

User code built as a shared object. The runtime needs to read this shared
object and run some bits on it.

__serialization__

All structs in user-code need to automatically be backed by something like
proto's for easy serialization.

## target language

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