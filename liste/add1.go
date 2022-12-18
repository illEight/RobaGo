package liste

/* dato il primo elemento di una lista e una nuova stringa aggiunge quella stringa all'inizio di quella lista
 deve quindi restituire il nuovo puntatore che punta a quel nuov elemento aggiunto in cima alla lista */

func AddFront(first *Node, x string) (newFirst *Node){
  newFirst=new(Node)
  (*newFirst).Value=x
  (*newFirst).Next=first
  return
}

//aggiunge elemento come ultimo elemento della lista e restituisce il nuovo primo elemento (necessario nel caso in cui)
//la lista sia vuota (first=nil) (in altri casi sarà uguale al vecchio primo elemento)
func AddLast(first *Node, x string ) (newFirst *Node){
  var curs *Node
  newNode:=new(Node)
  (*newNode).Value=x

  //se lista è vuota => puntatore non punta a niente
  if first==nil{
    newFirst=newNode
    return
  }

  //non mi fermo quando ho curs=nil ma quando curs punta a un nodo che ha come campo next nil
  for curs=first;(*curs).Next!=nil;curs=(*curs).Next {
  }
  //quindi mi serve curs che punta all'ultimo elemento
  (*curs).Next=newNode

  newFirst=first

  return

}
