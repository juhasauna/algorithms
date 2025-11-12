// Note that we do not need G(V, E) as an argument.
CheckMSCT(T, givenEdge(u, v)):
    // Assume we're given the decreased cost edge 
    // because the problem explicitly 
    // states its vertices "{u, v}".

    path = DFSForPathBetweenNodes(T, u, v)
    if length(path) == 1 then 
        // 'givenEdge' already exists in T.
        // The same T is still a valid MCST.
        return T 
    
    maxEdgeInCycle := givenEdge 
    for all e in path:
        if e.cost > maxEdgeInCycle.cost then
            maxEdgeInCycle = e

    // If givenEdge is the most costly in the cycle then
    // it is simply added and immediately removed
    // resulting in T that is identical to the original.
    T.add(givenEdge)
    T.remove(maxEdgeInCycle)
    return T

// Finds a path between two nodes if there is one.
// If not, it returns an empty list.
DFSForPathBetweenNodes(T, start, endNode):
    path = [] // Records edges from start to endNode.

    FindPath(x):
        // Marking nodes prevents moving backwards
        // since we have an undirected graph/tree.
        mark x 
        if endNode.mark then return 

        for all edges (x, y) in T:
            if not y.mark AND not endNode.mark then
                path.push((x, y))
                if y == endNode then 
                    mark endNode
                    // Path is done. 
                    // We should exit 'DFSForPathBetweenNodes'.
                    return
                
                FindPath(y)
                if not endNode.mark then
                    // Remove last entry when we return 
                    // to a previous level of recursion
                    // from a dead-end.
                    path.pop() 

    FindPath(start)
    return path
    