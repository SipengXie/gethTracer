package vm

import (
	"encoding/json"
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var graphPath = "./graph.html"

func getNodesAndLinks(g *DependencyGraph) (nodes []opts.GraphNode, links []opts.GraphLink) {
	nodes = make([]opts.GraphNode, 0)
	links = make([]opts.GraphLink, 0)

	for _, vertex := range g.Vertexes {
		byt, _ := json.Marshal(vertex)
		nodes = append(nodes, opts.GraphNode{
			Name: string(byt),
		})
	}

	for source, targets := range g.Edges {
		sourceByt, _ := json.Marshal(g.Vertexes[source])
		for target := range targets {
			targetByt, _ := json.Marshal(g.Vertexes[target])
			links = append(links, opts.GraphLink{
				Source: string(sourceByt),
				Target: string(targetByt),
			})
		}
	}

	return
}

func newChart(g *DependencyGraph) *charts.Graph {
	graph := charts.NewGraph()
	nodes, links := getNodesAndLinks(g)
	graph.AddSeries("", nodes, links, charts.WithGraphChartOpts(
		opts.GraphChart{
			Layout:     "none",
			EdgeSymbol: []string{"none", "arrow"},
		},
	))
	return graph
}

func VisualizeGraph(g *DependencyGraph) {
	page := components.NewPage()
	page.AddCharts(newChart(g))
	f, _ := os.Create(graphPath)
	page.Render(io.MultiWriter(f))
}
