package piscine

func zipString(s string) string {
	if s==" "{
		return " "
	}
	var result [] byte
	n:=len (s)
	i:=0
	for i< n{
		count :=1 
		for i+1 < n && s[i]==s[i+1]{
			i++
			count ++
		}
		result =append(result,byte(count+'0'))
		result =append(result,s[i])
		i++
	}
	return string ()

}