package main

const (
	B   = 1000000000
	M   = 1000000
	T   = 1000
	H   = 100
	TEN = 10
)

func numberToWords(num int) string {
	res := ""
	b := num / B
	if b > 0 {
		res += toHundredWords(b)
		res += " Billion "
	}
	m := (num - b*B) / M
	if m > 0 {
		res += toHundredWords(m)
		res += " Million "
	}
	t := (num - b*B - m*M) / T
	if t > 0 {
		res += toHundredWords(t)
		res += " Thousand "
	}
	h := num - b*B - m*M - t*T
	res += toHundredWords(h)
	if res == "" {
		return "Zero"
	}
	if h == 0 {
		res = res[:len(res)-1]
	}
	return res
}

func toHundredWords(num int) string {
	res := ""
	h := num / H
	if h > 0 {
		res += getOne(h)
		res += " Hundred "
	}
	ten := (num - h*H) / TEN
	if ten > 1 {
		res += getTen(ten)
	}
	one := num - h*H - ten*TEN
	if ten == 1 {
		res += getOneTen(one)
	} else {
		res += getOne(one)
	}
	if len(res) > 0 && res[len(res)-1] == ' ' {
		res = res[:len(res)-1]
	}
	return res
}

func getTen(num int) string {
	switch num {
	case 2:
		return "Twenty "
	case 3:
		return "Thirty "
	case 4:
		return "Forty "
	case 5:
		return "Fifty "
	case 6:
		return "Sixty "
	case 7:
		return "Seventy "
	case 8:
		return "Eighty "
	case 9:
		return "Ninety "
	default:
		return ""
	}
}

func getOneTen(num int) string {
	switch num {
	case 0:
		return "Ten"
	case 1:
		return "Eleven"
	case 2:
		return "Twelve"
	case 3:
		return "Thirteen"
	case 4:
		return "Fourteen"
	case 5:
		return "Fifteen"
	case 6:
		return "Sixteen"
	case 7:
		return "Seventeen"
	case 8:
		return "Eighteen"
	case 9:
		return "Nineteen"
	default:
		return ""
	}
}

func getOne(num int) string {
	switch num {
	case 1:
		return "One"
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	default:
		return ""
	}
}
