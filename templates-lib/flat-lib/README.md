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


```txt
$ tree --filesfirst
.
├── arrays.go
├── arrays_test.go
├── go.mod
├── go.sum
├── map.go
├── map_test.go
├── README.md
├── utils.go
├── utils_test.go
├── concurrent
│   ├── protect.go
│   ├── safemaps.go
│   └── wait.go
└── data
    ├── data.go
    ├── quicksort.go
    ├── quicksort_test.go
    ├── README.md
    ├── sort.go
    ├── container
    │   ├── container.go
    │   └── list
    │       ├── list.go
    │       ├── list_test.go
    │       ├── priority_queue.go
    │       ├── priority_queue_test.go
    │       ├── queue.go
    │       ├── queue_test.go
    │       ├── README.md
    │       ├── setlist.go
    │       ├── setlist_test.go
    │       ├── stack.go
    │       ├── stack_test.go
    │       ├── wrapper_safelist.go
    │       ├── wrapper_safelist_test.go
    │       ├── wrapper_setlist.go
    │       ├── wrapper_setlist_test.go
    │       ├── wrappers.go
    │       └── wrappers_test.go
    └── graph
        ├── bfs.go
        ├── bfs_test.go
        ├── errors.go
        ├── graph.go
        ├── hashmap_graph.go
        ├── hashmap_graph_impl.go
        ├── README.md
        ├── safe_graph.go
        ├── assets
        │   ├── directed_bfs_test_example.png
        │   └── example_flight_path
        │       ├── flight_graph.png
        │       └── main.go
        └── wgraph
            ├── dijkstra.go
            ├── dijkstra_impl.go
            ├── dijkstra_test.go
            ├── errors.go
            ├── hashmap_wgraph_impl.go
            ├── wedge.go
            ├── wgraph.go
            ├── wnode.go
            └── assets
                └── djikstra_test_graph.png
```