/**
 * @param {number[][]} edges
 * @return {number[]}
 */
var findRedundantConnection = function(edges) { // Brute force solution, build the graph with given coordinates and then traverse the graph
    let graph = {};
    let visited = [];
    for (let i = 0; i < edgets.length; i++) {
        const [x, y] = edges[i];
        if (!graph[x]) {
            graph[x] = []; 
        }
        if (!graph[y]) {
            graph[y] = [];
        }
        visited = [];
        if (graph[x].length && graph[y].length && dfs(x, y)) {
            return [x, y]; 
        }
        graph[x].push(y);
        graph[y].push(x);
    }
  
    function dfs(x, y) {
        if (visited.indexOf(x) === -1) {
            visited.push(x);
            if (x === y) {
                return true; 
            }
            return graph[x].map(ele => dfs(ele, y)).filter(ele => !!ele).length > 0;  // Recursively visit all nodes which x can reach and see if it can reach to y.
        }
    }
}
