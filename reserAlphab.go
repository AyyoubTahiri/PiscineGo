package piscine 

import "fmt"

func resrAlphab(str string) string {
	result :=" "
	for _, v := range str {
		if(v >= 'a' && v <= 'z'){
			index := int(v) - int('a')+1
			for i:= 0 ; i < index ; i++{
				result += string(v)
			}
		}else if(v >= 'A' && v <= 'Z'){
			index := int(v) - int('a')+1
			for i:= 0 ; i < index ; i++{
				result += string(v)
			}
		}else {
		   		 result += string(v)
		}

	}
}