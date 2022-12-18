package liste

/* creare funzione concatenate
che date due liste le concatena
(elementi della seconda lista dopo gli elementi della prima lista, mantenendo l’ordine che hanno le liste).
Deve restituire il puntatore al nuovo primo elemento della lista risultante. */

func Concatenate(firstList, secondList *Node) (ok bool, first *Node){
  var curs *Node

  //nel caso in cui le liste non esistono
  if firstList==nil && secondList==nil{
    return
  }

 //nel caso in cui la lista prima lista non esiste
 if firstList==nil {
   first=secondList
   return
 }


 //nel caso in cui la seconda lista non esiste
 if secondList==nil {
   first=firstList
   return
 }

 ok=true

 for curs=firstList;(*curs).Next!=nil;curs=(*curs).Next {
 }

 (*curs).Next=secondList

 first=firstList

 return
}
