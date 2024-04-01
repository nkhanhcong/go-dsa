This is some my thoughts about solving problems

</h1>Algorithm</h1>

- **DFS** is one of algorithms to travel graph or tree from a start node. Helping you travel all possible node of start node
- **Method to implement**:

    - You can use 2 options to solve problems: iterative loop ( using with stack) and recursion

- **Type of problems can be solve**:

    - Finding connected components
    - Finding path from one node to another node 
    - Detecting cycles in a graph
    - Topological sort

- **Tips**:

    - You can add some trick into "visited" variable, you can pull visited out and polute it with each time travel.
    - After doing EventualSafeNode, I realize when you use recursive it more clean code compare with using itterative
    - In some case when asking about minimum connected maked, we can base on attribute of graph: the graph n vertical need at least n-1 edges
    - In minimumFuelCost problem: I realize that when you meet the connected graph with non cycle, instead of using visitedMap to mark the node that is visited we can pass parent node to implement calculation
    