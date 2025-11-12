FindEulerianCircuit(G):
    if not IsConnected(G) then
        exit(graph not connected => No Eulerian circuit)
    if not NodesHaveEvenDegree(G) then 
        exit(∃ vertex with odd degree => No Eulerian circuit)
    
    circuit := empty list
    VisitNode(current):
        current.mark = true
        for all edges (current, other) in G:
            if not other.mark then
                VisitNode(other)
            circuit.push(current)

    startNode = any node in G
    VisitNode(startNode)
    return circuit

Time complexity is O(|V|+|E|) because we recurse over all vertices with "VisitNode" and during each visit we loop over the edges the current node. Resulting in looping over all edges.
Space complexity is the size of the 'circuit' variable. Which is O(|V|).