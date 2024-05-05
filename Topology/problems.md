Sequence reconstruction:
    - the problems is related issues: how we can detect whether it possilble to construct only one sequence from the given sequences
    - **Solution**: 
        - 1. Construct a graph from the given sequences
        - 2. Topological sort the graph using BFS 
        - 3. Check if the topological sort when add queue that have only one element 