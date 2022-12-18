// Provides a definition of the basic datatype required for lists

package main

//definizione di tipo
type Node struct {     //quindi questo tipo è una struttura che contiene puntatori allo stesso tipo => è un tipo di dato ricorsivo 
  Value string
  Next *Node
}
