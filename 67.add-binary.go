func addBinary(a string, b string) string {
	res := []byte{}
	carry := byte(0)
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		sum := carry
		if i >= 0 {
			sum += a[i] - '0'
		}
		if j >= 0 {
			sum += b[j] - '0'
		}
		if sum > 1 {
			carry = 1
		} else {
			carry = 0
		}
		res = append([]byte{sum%2 + '0'}, res...)
	}
	if carry > 0 {
		res = append([]byte{'1'}, res...)
	}
	return string(res)
}
