import "strings"

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	res := ""
	billion := num / 1000000000
	if billion > 0 {
		res += toHundred(billion) + " Billion "
	}
	million := (num - billion*1000000000) / 1000000
	if million > 0 {
		res += toHundred(million) + " Million "
	}
	thousand := (num - billion*1000000000 - million*1000000) / 1000
	if thousand > 0 {
		res += toHundred(thousand) + " Thousand "
	}
	res += toHundred(num - billion*1000000000 - million*1000000 - thousand*1000)
	return strings.TrimSpace(res)
}

func toHundred(num int) string {
	res := ""
	hundred := num / 100
	if hundred > 0 {
		res += getNum(hundred) + " Hundred "
	}
	ten := (num - hundred*100) / 10
	n := num - hundred*100 - ten*10
	if ten > 1 {
		res += getTen(ten) + " "
	} else if ten == 1 {
		res += getTen1(n) + " "
	}
	if n > 0 && ten != 1 {
		res += getNum(n) + " "
	}
	return strings.TrimSpace(res)
}

func getNum(num int) string {
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

func getTen(num int) string {
	switch num {
	case 2:
		return "Twenty"
	case 3:
		return "Thirty"
	case 4:
		return "Forty"
	case 5:
		return "Fifty"
	case 6:
		return "Sixty"
	case 7:
		return "Seventy"
	case 8:
		return "Eighty"
	case 9:
		return "Ninety"
	default:
		return ""
	}
}

func getTen1(num int) string {
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
