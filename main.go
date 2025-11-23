package main

import "fmt"

// Setup Array size 
const ArraySize = 7

// ----- struct -----
// HashTable struct
type HashTable struct{
	array [ArraySize] *bucket
}
// Bucket struct
type bucket struct{
	head *bucketNode
}
// Bucket Node struct
type bucketNode struct{
	key string
	next *bucketNode
}

// ----- functions ----- HASH TABLE
func(h *HashTable) Inser(key string){
	index := hash(key)
	h.array[index].insert(key)
}

func(h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

func(h *HashTable) Delete(key string) bool{
	index := hash(key)
	return h.array[index].search(key)
}

//HASH CODE
func hash(key string) int {
	sum := 0
	for _, v := range key{
		sum += int(v)
	}
	return sum % ArraySize
}

// FUNC FOR BUCKET
func (b *bucket) insert(k string){
	if b.search(k){
		fmt.Print("Already exist")
		return
	}
	newNode := &bucketNode{key:k}
	newNode.next = b.head
	b.head = newNode
	
}
// search
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil{
		if currentNode.key == k{
			return true
		}
		currentNode = currentNode.next
	}
	return false
}
// delete
func (b *bucket) delete(k string){
	if b.head.key == k{
		b.head = b.head.next
		return
	}
	prevNode := b.head
	for prevNode.next != nil{
		if prevNode.next.key == k{
			prevNode.next = prevNode.next.next
			return
		}
		prevNode = prevNode.next
	}
}
// init
func Init() *HashTable{
	result := &HashTable{}
	for i := range result.array{
		result.array[i] = &bucket{}
	}
	return result
}


func main(){
	haseTable := Init()
	list:= []string{
		"MAX",
		"ANNA",
		"JACK",
		"TOM",
		"BREADI",
		"LEBRON",
		"JAMES"	,
	}
	
	for _, v := range list{
		haseTable.Inser(v)

	}
	
	anna := haseTable.Search("ANNA")
	fmt.Println(anna)
	
}