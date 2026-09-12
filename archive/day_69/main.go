package main

func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	res := []byte{}

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}
		carry = sum / 10
		res = append(res, byte(sum%10)+'0')
	}
	for x, y := 0, len(res)-1; x < y; x, y = x+1, y-1 {
		res[x], res[y] = res[y], res[x]
	}
	return string(res)
}
