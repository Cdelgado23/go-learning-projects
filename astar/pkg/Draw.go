package astar

const (
	Normal = "\033[37m*\033[0m"
	Path   = "\033[32mx\033[0m"
	Start  = "\033[34mO\033[0m"
	End    = "\033[35mO\033[0m"
)

// Draw a map and the connections with its neighbors, using * for a node and - or | for a connection
func Draw(nodes [][]*Node, path []*Node) {
	for i := range nodes {
		for j := range nodes[i] {
			printNode(nodes[i][j], path)
		}
		println()
	}
}

func printNode(node *Node, path []*Node) {
	symbol := getSymbol(node, path)
	print(symbol)
}

// getSymbol returns the symbol to use for a node
func getSymbol(node *Node, path []*Node) string {
	for _, n := range path {
		if n == node {
			return getPathSymbol(node, path)
		}
	}
	return Normal
}

func getPathSymbol(node *Node, path []*Node) string {
	if node == path[0] {
		return Start
	}
	if node == path[len(path)-1] {
		return End
	}
	return Path
}
