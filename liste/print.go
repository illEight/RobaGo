package liste
import "fmt"

//funzione per stampare valore nodi della lista

func Print(first *Node) {
  var curs *Node

  for curs = first; curs!=nil;curs=(*curs).Next {
    if (*curs).Next == nil {
      fmt.Printf("%s\n", (*curs).Value )
    } else {
      fmt.Printf("%s -> ", (*curs).Value )
    }
  }

}
