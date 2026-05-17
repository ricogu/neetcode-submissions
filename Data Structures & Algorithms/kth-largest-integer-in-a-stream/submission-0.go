
import "slices"

type KthLargest struct {
    elements []int
	k int
}


func Constructor(k int, nums []int) KthLargest {
	
	sortNums := nums
	slices.Sort(sortNums)

	return KthLargest{
		elements: sortNums,
		k : k,
	}
    
}


func (this *KthLargest) Add(val int) int {
   indexToInsert := len(this.elements)

   for idx , v := range this.elements {
	if val <= v{
		indexToInsert = idx
		break
	}
   }

   this.elements = slices.Insert(this.elements, indexToInsert, val)

   return this.elements[len(this.elements)-this.k]

    
}
