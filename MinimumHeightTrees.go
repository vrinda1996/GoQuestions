func findMinHeightTrees(n int, edges [][]int) []int {
    if n == 1{
        return []int{0}
    }
    graph := make([][]int, n)
    inDegree := make([]int,n)
    
    for _, edge := range edges {
        graph[edge[0]] = append(graph[edge[0]], edge[1])
        graph[edge[1]] = append(graph[edge[1]], edge[0])
        inDegree[edge[0]]++
        inDegree[edge[1]]++
    }
    queue := make([]int,0)
    // find leaf nodes
    for i, deg := range inDegree{
        if deg == 1 {
            queue = append(queue, i)
        }
    }

    for n > 2{
        size := len(queue)
        n -= size
        for i:=0 ;i < size; i++ {
            inDegree[queue[i]]--

            for _, to := range graph[queue[i]]{
                inDegree[to]--
                
                if inDegree[to] ==1 {
                    queue = append(queue, to)
                }
            }
        }
        queue = queue[size:]
    }
    return queue 
}
