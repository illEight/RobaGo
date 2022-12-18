
package liste
import "fmt"

func main(){
  var first,second *Node

  /* viene creata la lista a partire dall'ultimo nodo
  first = AddFront(first, "camilla")
  first = AddFront(first, "è")
  first = AddFront(first, "stata")
  first = AddFront(first, "ferita")

  temp=first
  first = AddFront(first, "ma")
  first = AddFront(first, "anna")
  first = AddFront(first, "non")
  first = AddFront(first, "capisce") */

  first = AddLast(first, "camilla")
  first = AddLast(first, "è")
  first = AddLast(first, "stata")
  first = AddLast(first, "ferita")
  first = AddLast(first, "ma")
  first = AddLast(first, "anna")
  first = AddLast(first, "non")
  first = AddLast(first, "capisce")
  //first = AddFront(first, "povera")
  //first = AddFront(first, "la")

  /* first = AddInOrder(first, "camilla")
  first = AddInOrder(first, "è")
  first = AddInOrder(first, "stata")
  first = AddInOrder(first, "ferita")
  first = AddInOrder(first, "ma")
  first = AddInOrder(first, "anna")
  first = AddInOrder(first, "non")
  first = AddInOrder(first, "capisce") */


  fmt.Printf("Lunghezza = %d\n",Length(first))
  Print(first)
  //Print(temp)   //per stampare sottolista


    second = AddLast(second, "chissà")
    second = AddLast(second, "se")
    second = AddLast(second, "mi")
    second = AddLast(second, "scriverà")
    second = AddLast(second, "mercoledì")
    second = AddLast(second, "o")
    second = AddLast(second, "giovedì")

    ok,newFirst:=Concatenate(first,second)

    if ok {
      Print(newFirst)
    } else {
      fmt.Println("Una delle due liste ( o entrambe) è vuota.")
    }
}
