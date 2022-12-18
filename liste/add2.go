package main


//se la lista di stringhe è già in ordine alfabetico e viene chiamata questa funzione, verrà aggiunta stringa x mantenendo la lista ordinata in ordine alfabetico
//aggiunge in coda
func AddInOrder(first *Node, x string) (newFirst *Node){
  var curs,prev *Node

  newNode:=new(Node)
  (*newNode).Value=x

  if first==nil {
   newFirst=newNode
   return
  }

  //exit condition: curs==nil || (*curs).Value > x
  for curs=first; curs!=nil && (*curs).Value<=x ;curs=(*curs).Next {
    prev=curs
  }

 //questo if è stato aggiunto dopo, quando ci siamo accorti che in fase di compilazione dava errore sull'attuale riga 28: dice che non era possibile
 //deferenziare prev (con *prev) in quanto prev era nil. Questo significa che stavamo aggiungendo un nodo in cima alla lista che quindi per le condizioni del for
 //non era entrato nel ciclo e prev aveva valore nil => bisogna gestire il caso a parrte come scritto sotto:
  if prev==nil{              //prev è uguale a nil se si mette nodo in cima
    (*newNode).Next=curs
    newFirst=newNode
    return
  }

  //mette il nuovo nodo IN MEZZO
  (*prev).Next=newNode
  (*newNode).Next=curs

  newFirst=first

  return

}
