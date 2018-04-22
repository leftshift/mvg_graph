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
    for key, v := range result.Relations {
        if v.Meta.Tags["ref"] == "" {
            delete(result.Relations, key)
            continue
        }

        fmt.Printf("key:[%v], ref:[%s], rel:[%s]\n", v.Meta.ID, v.Meta.Tags["ref"], v.Meta.Tags["type"])
        for k, m := range v.Members {
            if m.Role == "stop" {
                fmt.Printf("\t%v: %v\n", k, m.Node.Meta.Tags["name"])
            }
        }
    }
}
