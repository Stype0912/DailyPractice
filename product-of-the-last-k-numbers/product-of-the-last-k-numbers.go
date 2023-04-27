package product_of_the_last_k_numbers

type ProductOfNumbers struct {
	Nums    []int
	Product map[int]int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		Nums:    []int{},
		Product: map[int]int{},
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if len(this.Nums) != 0 {
		for i := len(this.Nums); i >= 1; i-- {
			this.Product[i+1] = this.Product[i] * num
		}
	}
	this.Product[1] = num
	this.Nums = append(this.Nums, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	return this.Product[k]
}
