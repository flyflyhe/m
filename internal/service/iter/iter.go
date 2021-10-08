package iter

/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type Iterator struct {}

func (this *Iterator) hasNext() bool {
	return true
}

func (this *Iterator) next() int {
	return 1
}

type PeekingIterator struct {
	Result *Result
	Iter *Iterator
}

type Result struct {
	Val int
}

func Constructor(iter *Iterator) *PeekingIterator {
	var result *Result
	if iter.hasNext() {
		result = &Result{
			Val: iter.next(),
		}
	}
	return &PeekingIterator{
		Result: result,
		Iter: iter,
	}
}

func (this *PeekingIterator) hasNext() bool {
	return this.Result != nil
}

func (this *PeekingIterator) next() int {
	ret := this.Result.Val
	if this.Iter.hasNext() {
		this.Result.Val = this.Iter.next()
	} else {
		this.Result = nil
	}
	return ret
}

func (this *PeekingIterator) peek() int {
	return this.Result.Val
}
