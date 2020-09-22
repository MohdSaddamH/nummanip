package calc

/*** Add two number 
* return int
*/
func Add(numbers ...int) int{
	
	var sum int = 0
	for _, num := range numbers{
		sum += num
	}
	
	return sum
}	
