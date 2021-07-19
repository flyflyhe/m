package atoi

type Automation struct {
	Sign int
	Ans int64
	State string
	Table map[string][]string
}

var table map[string][]string

func init()  {
	table = make(map[string][]string)
	table["start"] = []string{"start", "signed", "in_number", "end"}
	table["signed"] = []string{"end", "end", "in_number", "end"}
	table["in_number"] = []string{"end", "end", "in_number", "end"}
	table["end"] = []string{"end", "end", "end", "end"}
}

func GetAutomation() *Automation {
	return &Automation{Sign: 1,State: "start",Table:table}
}

func (this *Automation) get(c byte) {
	this.State = table[this.State][this.getCol(c)]
	if "in_number" == this.State {
		this.Ans = this.Ans * 10 + int64(c - '0')
	} else if "signed" == this.State {
		if c == '+' {
			this.Sign = 1
		} else {
			this.Sign = -1
		}
	}
}

//获取状态
func (this *Automation) getCol(c byte) int {
	if c == ' ' {
		return 0
	}
	if c == '+' || c == '-' {
		return 1
	}
	if  c - '0' >= 0 && c - '0' <= 9 {
		return 2
	}
	return 3
}

