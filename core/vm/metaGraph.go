package vm

type DependencyGraph struct {
	Vertexes map[int]Metadata
	Edges    map[int]map[int]struct{}
}

func NewDependencyGraph() *DependencyGraph {
	v := make(map[int]Metadata)
	v[-1] = SourceMeta

	return &DependencyGraph{
		Vertexes: v,
		Edges:    make(map[int]map[int]struct{}), // 类似于邻接矩阵吧
	}
}

func (g *DependencyGraph) AddDependency(sources []Metadata, target Metadata) {
	// Add vertexes if not exist
	for _, meta := range sources {
		if _, ok := g.Vertexes[meta.Index]; !ok {
			g.Vertexes[meta.Index] = meta
		}
	}
	if _, ok := g.Vertexes[target.Index]; !ok {
		g.Vertexes[target.Index] = target
	}

	// Add edges
	for _, meta := range sources {
		if _, ok := g.Edges[meta.Index]; !ok {
			g.Edges[meta.Index] = make(map[int]struct{})
		}
		g.Edges[meta.Index][target.Index] = struct{}{}
	}
}
