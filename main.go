package main

func main() {
	var it *Three
	it = it.Insert(OrderableInt(5))
}

type Orderable interface{
	Order (any) int
}

type OrderableInt int

func (oi OrderableInt) Order(val any) int{
	return int(oi - val.(OrderableInt))
}

type Three struct{
	val Orderable
	left, right *Three
}

func (t *Three) Insert(val Orderable) *Three{
	if t == nil{
		return &Three{val: val}
	}
	
	switch comp := val.Order(t.val); {
		case comp < 0:
		t.left = t.left.Insert(val)
		case comp > 0:
		t.right = t.left.right.Insert(val)
	}
	return t
}