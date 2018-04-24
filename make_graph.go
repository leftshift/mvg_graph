package main
import (
    "fmt"
    "log"
    "github.com/serjvanilla/go-overpass"
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

func main(){
    client := overpass.New()
    result, ok := client.Query(query)
    if ok != nil {
        log.Fatal(ok)
    } 
    //fmt.Println(result)
    for key, relation := range result.Relations {
        if relation.Meta.Tags["ref"] == "" {
            delete(result.Relations, key)
            continue
        }

        fmt.Printf("key:[%v], ref:[%s], rel:[%s]\n", relation.Meta.ID, relation.Meta.Tags["ref"], relation.Meta.Tags["type"])
        for k, member := range relation.Members {
            if member.Role == "stop" {
                fmt.Printf("\t%v: %v\n", k, member.Node.Meta.Tags["name"])
            } else {
                // delete(relation.Members, k)
            }
        }
    }
}
