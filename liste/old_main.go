/* dichiara due puntatori a nodi (p1 e p2) e collega questi due nodi */
package liste
import "fmt"

func main(){
  var p1,p2,p3,p4 *Node

  p1 = new(Node)  //crea nello heap uno spazio per contenere un nodo e assegna a p1 riferimento a quello spazio
  p2 = new(Node)
  p3 = new(Node)
  p4 = new(Node)

  //fmt.Println(*p1)   => stampa nodo (che è una struttura: stringa vuota + valore iniziale del puntatore=nil)
  //fmt.Println(*p2)   => stampa nodo (che è una struttura: stringa vuota + valore iniziale del puntatore=nil)

  //riempire campi valore
  (*p1).Value="pippo"    //si può togliere * perchè go riconosce che è una deferenziazione del puntatore (ovvero il puntatore punta a una struct e noi con * ci riferiamo alla struct)
  (*p1).Next=p2         //attacco il secondo nodo al primo   (in teoria adesso p2 non serve più)

  (*p2).Value="pluto"
  (*p2).Next=p3

  (*p3).Value="camilla"
  (*p3).Next=p4

  (*p4).Value="bestemmia"


  fmt.Printf("p1=%p p2=%p p3=%p p4=%p\n",p1,p2,p3,p4)       //stampa valore p1 e p2
  fmt.Println(*p1)                                          //=> stampa nodo (che è una struttura: pippo + valore del puntatore=p2)
  fmt.Println(*p2)                                          //=> stampa nodo (che è una struttura: pluto  + valore iniziale del puntatore=nil)
  fmt.Println(*p3)
  fmt.Println(*p4)

  fmt.Printf("Lunghezza = %d\n",Length(p1))
  Print(p1)

}
