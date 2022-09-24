# Bellman-Ford Shortest Path Algorithm

Bellman-Ford algorithm is a single-source shortest path algorithm, so when you have negative edge weight then it can detect negative cycles in a graph.

The only difference between the two is that Bellman-Ford is also capable of handling negative weights whereas Dijkstra Algorithm can only handle positives.

Dijkstra is however generally considered better in the absence of negative weight edges, as a typical binary heap priority queue implementation has `O((|E|+|V|)log|V|)` time complexity (A Fibonacci heap priority queue gives `O(|V|log|V|+|E|)`), while the Bellman-Ford algorithm has `O(|V||E|)` complexity.[[1](https://stackoverflow.com/questions/19482317/bellman-ford-vs-dijkstra-under-what-circumstances-is-bellman-ford-better)]

File shortestpath.cpp contain C++ code for Bellman-Ford algorithm for a graph with 13 nodes represented in adjacency matrix. You can compile the code using the following command

```
$ g++ shortestpath.cpp
```

Then you can try to execute the program

```
$ ./a.out
Shortest distance of node A from A is 0
Shortest distance of node B from A is 10
Shortest distance of node C from A is 9
Shortest distance of node D from A is 5
Shortest distance of node E from A is 17
Shortest distance of node F from A is 13
Shortest distance of node G from A is 14
Shortest distance of node H from A is 22
Shortest distance of node I from A is 20
Shortest distance of node J from A is 6
Shortest distance of node K from A is 23
Shortest distance of node L from A is 20
Shortest distance of node M from A is 13
```
