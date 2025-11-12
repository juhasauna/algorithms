IsDFSTree(G, v, T):
    // Global bookkeeping for time. 
    // Results in unique in/out times for each vertix in T.
    // Ancestor (a) of a vertex (b) will always have 
    // a.t_in < b.t_in AND b.t_out < a.t_out
    time := 0 

    DFS_ON_T(w):
        time++
        mark w
        w.t_in = time
        for each (w, u) in T:
            if not u.mark then
                DFS_ON_T(u)
        time++
        w.t_out := time

    DFS_ON_T(v)

    IsAncestor(w, u):
        return w.t_in < u.t_in AND u.t_out < w.t_out

    for each edge (w, u) in G: // Check if Lemma 3 is satisfied.
        if (w, u) not in T then 
            isBackEdge := IsAncestor(w, u) OR IsAncestor(u, w)
            if not isBackEdge then 
                return false

    return true