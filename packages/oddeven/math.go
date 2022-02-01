package math



func Evenodd(num int,flag bool) bool {
	
	if num%2 == 0 {
		flag = true
	} else {
		flag = false
	}
	
	return flag
}