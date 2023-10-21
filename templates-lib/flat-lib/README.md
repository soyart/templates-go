# flat-lib

Go libraries are usually structured very flatly.

Most of the times, code is only placed in the module root.
The central/main symbols of a library is always placed in its root.
Other secondary symbols can be placed in their own modules.

## flat-lib structure

flat-lib is modeled after [gsl](https://github.com/soyart/gsl), which is
a general purpose Go library containing 3 main categories of code:

- Batteries-included (general utils)

    The code in this category is placed at the root of the library

- Concurrent programming utility

    The code in this category is placed at `./concurrent`. All files
    are just direct children of the package.

- Data structure code

    The code in this category is placed at `./data`. The general types
    and functions of this category is placed at root `./data/`, e.g.
    `./data/data.go` and `./data/quicksort.go`.

    Sub-categories like graphs and containers are placed inside their own
    modules `./data.graph` `./data/container`.