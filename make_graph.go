package main
import "fmt"
import "github.com/serjvanilla/go-overpass"

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
    result, _ := client.Query(query)
    //fmt.Println(result)
    for k, v := range result.Relations {
        fmt.Printf("key:[%s], ref:[%s]\n", k, v.Meta.Tags["ref"])
    }
}
