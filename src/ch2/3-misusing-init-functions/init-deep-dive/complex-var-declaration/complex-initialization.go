// complex_initialization.go
package main

import (
	"fmt"
)

// =========================================================================
// Use Case 1: Complex Initialization of Variables
// =========================================================================
// The init() function is ideal for complex variable initialization that
// cannot be expressed as simple declarations. This includes:
// - Building complex data structures requiring multiple steps
// - Performing calculations to derive values
// - Loading and transforming data

// Example 1.1: Complex graph data structure initialization
// This example demonstrates creating and configuring a graph with nodes,
// edges, and calculated weights - something too complex for a single declaration.

type Node struct {
	ID    string
	Value int
}

type Graph struct {
	nodes map[string]*Node
	edges map[string][]string
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string]*Node),
		edges: make(map[string][]string),
	}
}

func (g *Graph) AddNode(id string) *Node {
	node := &Node{ID: id}
	g.nodes[id] = node
	return node
}

func (g *Graph) AddEdge(from, to string) {
	g.edges[from] = append(g.edges[from], to)
}

func (g *Graph) Nodes() []*Node {
	nodes := make([]*Node, 0, len(g.nodes))
	for _, node := range g.nodes {
		nodes = append(nodes, node)
	}
	return nodes
}

// Package-level variables to be initialized in init()
var graph *Graph
var nodeWeights map[string]int

// Helper function for initialization
func calculateInitialWeight(node *Node, g *Graph) int {
	// Simplified weight calculation - counts number of connections
	connections := len(g.edges[node.ID])
	return connections * 10
}

// First init function - builds the graph structure and calculates node weights
// This is a perfect use case for init() because:
// - We need to create the graph first
// - Then add nodes
// - Then add edges
// - Then calculate derived values (weights) based on the structure
func init() {
	fmt.Println("Running init for graph data structure...")

	// Create graph
	graph = NewGraph()

	// Add nodes
	graph.AddNode("A")
	graph.AddNode("B")
	graph.AddNode("C")
	graph.AddNode("D")

	// Add edges
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "D")

	// Initialize weights based on graph structure
	nodeWeights = make(map[string]int)
	for _, node := range graph.Nodes() {
		nodeWeights[node.ID] = calculateInitialWeight(node, graph)
	}
}

// Example 1.2: Generating a lookup map with computed values
// This example shows how to generate a lookup table with both
// direct assignments and programmatically derived values.

// Package-level map to be initialized
var romanNumerals map[int]string

// Second init function - populates the roman numerals map
// This demonstrates:
// - Making the map
// - Setting base values
// - Deriving additional values programmatically
func init() {
	fmt.Println("Running init for roman numerals map...")

	romanNumerals = make(map[int]string)

	// Base numerals (these could be simple declarations,
	// but we want to compute additional values)
	romanNumerals[1] = "I"
	romanNumerals[5] = "V"
	romanNumerals[10] = "X"
	romanNumerals[50] = "L"
	romanNumerals[100] = "C"
	romanNumerals[500] = "D"
	romanNumerals[1000] = "M"

	// Derive additional entries programmatically
	// This demonstrates why an init function is useful -
	// we can write logic to generate values
	for _, base := range []int{1, 10, 100} {
		romanNumerals[base*4] = romanNumerals[base] + romanNumerals[base*5]
		romanNumerals[base*9] = romanNumerals[base] + romanNumerals[base*10]
	}
}

// Display the initialized structures when imported
func DisplayComplexData() {
	fmt.Println("\n=== Complex Initialization Results ===")

	fmt.Println("\nGraph node weights:")
	for nodeID, weight := range nodeWeights {
		fmt.Printf("Node %s: %d\n", nodeID, weight)
	}

	fmt.Println("\nRoman numerals map:")
	for _, num := range []int{4, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000} {
		fmt.Printf("%d = %s\n", num, romanNumerals[num])
	}
}

// This main function only executes when running this file directly
func main() {
	fmt.Println("Complex Initialization Examples")
	DisplayComplexData()
}
