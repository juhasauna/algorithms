Algorithm Print_DFS_Edges(G(V, E), v);
    // The edges that DFS traverses form a tree. 
    // In a tree, apart from the root, 
    // all vertices have exactly one parent.
    // We save a pointer to this parent for each vertex.
	v.parent = nil
    push v to Stack; 
    
    while Stack is not empty do
        pop w from Stack;
        if w is unmarked then

            // We print out only the edge 
            // that corresponds to the last time
            // that a given vertex was pushed to the stack.
            // Marking w and the LIFO nature of the stack ensures this.
            // The printed edge goes from w.parent to w.

            if w.parent not nil then PRINT(parent "->" w)

            mark w;
            for all edges (w, x) such that x is unmarked do
                // It is possible for the same vertex to be
                // pushed onto the stack multiple times.
                // But only the last 'parent -> x' edge gets printed.
                x.parent = w

                push x to Stack