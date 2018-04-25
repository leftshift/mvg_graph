package main
import (
    "fmt"
    "log"
    "github.com/serjvanilla/go-overpass"
//    "gonum.org/v1/gonum/graph"
    "gonum.org/v1/gonum/graph/topo"
    "github.com/leftshift/mvg_graph/graph"
)

var query string = `
[out:json];
rel(7099055);
rel(r);
foreach(
  ._;
  rel(r)
  	->.route;
  (
  	node(r.route:"stop");
    node(r.route:"stop_exit_only");
  )->.stops;
  (
  	rel.route;
    node.stops;
  );
  out;
);`

type Line []*network_graph.Node

var g *network_graph.Graph = network_graph.NewGraph()
var lines = make(map[string]Line)

func addLinesAndStations(g *network_graph.Graph, result *overpass.Result) {
    for key, relation := range result.Relations {
        if relation.Meta.Tags["ref"] == "" {
            delete(result.Relations, key)
            continue
        }

        fmt.Printf("key:[%v], ref:[%s], rel:[%s]\n", relation.Meta.ID, relation.Meta.Tags["ref"], relation.Meta.Tags["type"])
        ref := relation.Meta.Tags["ref"]
        var add bool

        if _, ok := lines[ref]; !ok {
            lines[ref] = make([]*network_graph.Node, 0)
            add = true
        }
        var previous_station *network_graph.Node
        for k, member := range relation.Members {
            if member.Role == "stop" {
                name := member.Node.Meta.Tags["name"]
                fmt.Printf("\t%v: %v\n", k, name)
                if add {
                    station := g.NewNode()
                    station.Name = name
                    g.AddNode(&station)
                    lines[ref] = append(lines[ref], &station)

                    if previous_station != nil {
                        e := g.NewWeightedEdge(station, previous_station, 10)
                        f := g.NewWeightedEdge(previous_station, station, 10)
                        g.SetWeightedEdge(e)
                        g.SetWeightedEdge(f)
                    }
                    previous_station = &station
                }
            }
        }
    }
}

func main(){
    client := overpass.New()
    result, ok := client.Query(query)
    if ok != nil {
        log.Fatal(ok)
    } 

    addLinesAndStations(g, &result)

    fmt.Println(g.Edges())
    for ref, stations := range lines {
        fmt.Printf("%v:\n", ref)
        path1 := topo.PathExistsIn(g, stations[0], stations[len(stations)-1])
        path2 := topo.PathExistsIn(g, stations[len(stations)-1], stations[0])

        fmt.Printf("Path through line: %v", path1 && path2)
        for _, station := range stations {
            fmt.Printf("\t%v\n", station.Name)
        }
    }
}
