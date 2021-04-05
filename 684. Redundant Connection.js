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

var findRedundantConnection = function(edges) {  // Union-Find for graph path
    const parent = [];
    const rank = [];

    function find(x) {
        if (parent[x] !== x) {                  // Find the parent node, compress path to make all non-root node point to root node.
            parent[x] = parent[parent[x]];
            x = parent[x];
        }
        return x;
    }
    
    for (let i = 0; i < edges.length; i++) {
        const [x, y] = edges[i];
        if (!parent[x]) {
            parent[x] = x;   
        }
        if (!parent[y]) {
            parent[y] = y;
        }
        if (!rank[x]) {
            rank[x] = 1;
        }
        if (!rank[y]) {
            rank[y] = 1;   
        }
        const parentX = find(x);
        const parentY = find(y);
        
        if (parentX === parentY) {
            return [x, y];   
        }
        
        if (rank[parentY] > rank[parentX]) {           // Always merge smaller tree to bigger one.
            [parentY, parentX] = [parentX, parentY];   
        }
        
        parent[parentY] = parentX;
        rank[parentX] += rank[parentY];
    }
    return [];
}
