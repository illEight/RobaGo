package main

//Givene the pointer to the first node of a list, returns the number of elements => bisogna quinidi scandire la lista tramite un cursore curs (puntatore)

func Length(first *Node) int {
  var curs *Node
  count:=0
  for curs = first; curs!=nil;curs=(*curs).Next {
    //fmt.Println((*curs).Value )
    count++
  }

  return count
}
