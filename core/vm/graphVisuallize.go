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
		// vid := vertex.Index
		nodes = append(nodes, opts.GraphNode{
			Name: string(byt),
		})
	}

	for source, targets := range g.Edges {
		vsource := g.Vertexes[source]
		vsbyt, _ := json.Marshal(vsource)
		for target := range targets {
			vtarget := g.Vertexes[target]
			vtbyt, _ := json.Marshal(vtarget)
			links = append(links, opts.GraphLink{
				Source: string(vsbyt),
				Target: string(vtbyt),
			})
		}
	}

	return
}

func newChart(g *DependencyGraph) *charts.Graph {
	graph := charts.NewGraph()
	nodes, links := getNodesAndLinks(g)
	graph.AddSeries("", nodes, links).SetSeriesOptions(
		charts.WithGraphChartOpts(opts.GraphChart{
			Layout:             "force",
			Force:              &opts.GraphForce{Repulsion: 100},
			Roam:               true,
			FocusNodeAdjacency: true,
			EdgeSymbol:         []string{"none", "arrow"},
		}),
		charts.WithEmphasisOpts(opts.Emphasis{
			Label: &opts.Label{
				Show:     true,
				Color:    "black",
				Position: "left",
			},
		}),
	)
	return graph
}

func VisualizeGraph(g *DependencyGraph) {
	page := components.NewPage()
	page.AddCharts(newChart(g))
	f, _ := os.Create(graphPath)
	page.Render(io.MultiWriter(f))
}
