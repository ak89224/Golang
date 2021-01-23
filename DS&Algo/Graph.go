package main

import "fmt"

type Vertex struct {
	key int
	adjVertices map[int]*Vertex
}

type Graph struct {
	 edges map[int]*Vertex
	 directed bool
}

func NewVertex(key int) *Vertex {
	return &Vertex{
		key:      key,
		adjVertices: map[int]*Vertex{},
	}
}

func NewDirectedGraph() *Graph {
	return &Graph{
		edges: map[int]*Vertex{},
		directed: true,
	}
}

func NewUndirectedGraph() *Graph {
	return &Graph{
		edges: map[int]*Vertex{},
	}
}

func (g *Graph) AddVertex(key int) {
	if _, ok := g.edges[key]; ok {
		return
	}
	v := NewVertex(key)
	g.edges[v.key] = v
}

func (g *Graph) AddEdge(src int, dest int) {
	if _, ok := g.edges[src]; !ok {
		g.AddVertex(src)
	}
	if _, ok := g.edges[dest]; !ok {
		g.AddVertex(dest)
	}

	if _, ok := g.edges[src].adjVertices[dest]; !ok {
		g.edges[src].adjVertices[dest] = g.edges[dest]
	}

	if !g.directed {
		if _, ok := g.edges[dest].adjVertices[src]; !ok {
			g.edges[dest].adjVertices[src] = g.edges[src]
		}
	}
}

func (g *Graph) DSF(startVertex *Vertex, visited map[int]bool)  {
	if startVertex != nil {
		visited[startVertex.key] = true
		fmt.Print(startVertex.key, "-->")
		for _, v := range startVertex.adjVertices {
			if visited[v.key] {
				continue
			}
			g.DSF(v, visited)
		}
	}
	return
}

func (g *Graph) FindPathDFS(src *Vertex, dest *Vertex, visited map[int]bool) bool {
	if src == dest{
		return true
	}

	if src != nil {
		visited[src.key] = true
		fmt.Print(src.key, "-->")
		for _, v := range src.adjVertices {
			if visited[v.key] {
				continue
			}
			isCorrectPath := g.FindPathDFS(v, dest, visited)
			if isCorrectPath {
				return true
			}
		}
	}
	return false
}

func (g *Graph) FindPathDFSStack(src *Vertex, dest *Vertex, visited map[int]bool) bool {
	if src == dest {
		return true
	}
	visited[src.key] = true

	var stack []*Vertex
	stack = append(stack, src)

	for len(stack) != 0 {
		l := len(stack)
		curr_vertex := stack[l-1]
		stack = stack[:l-1]
		if curr_vertex == dest {
			return true
		}
		for _, v := range curr_vertex.adjVertices {
			if visited[v.key] {
				continue
			}
			visited[v.key] = true
			stack = append(stack, v)
		}

	}
	return false
}

func (g *Graph) BFS(src, dest *Vertex, visited map[int]bool) int {
	if src == dest {
		return 1
	}
	visited[src.key] = true
	var distance int
	parent := make([]int, len(g.edges))

	var queue []*Vertex
	queue = append(queue, src)
	parent[src.key] = -1

	for len(queue) != 0 {
		curr_vertex := queue[0]
		queue = queue[1:]
		if curr_vertex == dest {
			break
		}
		for _, v := range curr_vertex.adjVertices {
			if visited[v.key] {
				continue
			}
			visited[v.key] = true
			queue = append(queue, v)
			parent[v.key] = curr_vertex.key
		}

	}
	temp := dest.key
	for parent[temp] != -1{
		fmt.Print(temp, "-->")
		distance += 1
		temp = parent[temp]
	}
	fmt.Println()
	return distance
}

func (g *Graph) isCyclic(src *Vertex, visited map[int]bool, parent int) bool {
		visited[src.key] = true

		for _, v := range src.adjVertices {
			if visited[v.key] {
				if v.key != parent {
					return true
				}
				continue
			}

			if g.isCyclic(v, visited, src.key) {
				return true
			}

		}
	return false
}

func TopologicalSortUtil(src *Vertex, visited map[int]bool, stack *[]int){
	visited[src.key] = true

	for _, v := range src.adjVertices {
		if visited[v.key] {
			continue
		}
		TopologicalSortUtil(v, visited, stack)
	}
	*stack = append(*stack, src.key)
}

func (g *Graph) TopologicalSort(){
	var src *Vertex = g.edges[0]
	visited := map[int]bool{}
	stack := []int{}

	TopologicalSortUtil(src, visited, &stack)

	for i := range stack {
		fmt.Print(stack[len(stack)-1 - i])
	}
	fmt.Println()
}

func (g *Graph) Print()  {
	for k, v := range g.edges {
		fmt.Println(k,"-->",v.adjVertices)
	}
}


func main() {
	g := NewUndirectedGraph()
	g.AddVertex(0)
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)

	g.AddEdge(0, 4)
	g.AddEdge(4, 3)
	g.AddEdge(3, 2)
	g.AddEdge(2, 1)
	g.AddEdge(1, 2)
	g.AddEdge(5, 1)
	g.AddEdge(7, 6)
	g.AddEdge(7, 1)
	g.AddEdge(7, 7)
	g.AddEdge(9, 8)

	g.Print()
	visited := map[int]bool{}
	g.DSF(g.edges[1], visited)
	fmt.Println()
	visited = map[int]bool{}
	fmt.Println(g.FindPathDFS(g.edges[0], g.edges[7], visited))
	visited = map[int]bool{}
	fmt.Println(g.FindPathDFSStack(g.edges[0], g.edges[3], visited))
	visited = map[int]bool{}
	fmt.Println(g.BFS(g.edges[0], g.edges[3], visited))
	visited = map[int]bool{}
	fmt.Println(g.isCyclic(g.edges[0], visited, -1))
	visited = map[int]bool{}
	g.TopologicalSort()
	//fmt.Println(stack)
	fmt.Println()


}

