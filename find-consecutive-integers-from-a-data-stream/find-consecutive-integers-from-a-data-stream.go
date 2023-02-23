package find_consecutive_integers_from_a_data_stream

type DataStream struct {
	value  int
	values []int
	count  int
}

func Constructor(value int, k int) DataStream {
	return DataStream{
		value:  value,
		count:  k,
		values: nil,
	}
}

func (this *DataStream) Consec(num int) bool {
	if len(this.values) == this.count {
		this.values = this.values[1:]
	}
	if len(this.values) < this.count {
		this.values = append(this.values, num)
	}
	if len(this.values) == this.count {
		for _, value := range this.values {
			if value != this.value {
				return false
			}
		}
		return true
	}
	return false
}
