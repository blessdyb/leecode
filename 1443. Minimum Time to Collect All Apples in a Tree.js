/**
 * @param {number} n
 * @param {number[][]} edges
 * @param {boolean[]} hasApple
 * @return {number}
 */
var minTime = function(n, edges, hasApple) {
    // Given edges, intuitively we build a new undirect / direct graph (based on the given condition that if all edges are start from root to leaf node), so for this one it would be a undirectional graph
    const graph = Array(n).fill(0).map((i, id) => []);
    edges.forEach(edge => {
        const [x, y] = edge;
        graph[x].push(y);
        graph[y].push(x);
    });
    let visited = {};  // For undirectional graph, we need to avoid the infinite loop
    return (function dfs(node) {   // Since any node cost depends on it's child node, so we can use DFS postOrder to accumulate all node's cost based on its child nodes    
        let total = 0;
        visited[node] = true;
        graph[node].forEach(c => {
            if (!visited[c]) {            
                const childCost = dfs(c); 
                if (childCost > 0 || hasApple[c]) {
                    total += childCost + 2; 
                }
            }
        });
        return total;
    })(0);
};
