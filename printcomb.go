package piscine

import "github.com/01-edu/z01"

func PrintComb(){
 var ones int=0
  var tens int=0
  var hands int=0
  
  
  for i:=0 ; i<=789 ; i++ {

    if ones<9{
      ones++
    }else if tens<9{
      tens++;ones=0
    }else{
      hands++;tens=0;ones=0
    }
    if ones>tens && tens >hands{
    z01.PrintRune(rune(hands) +'0')
    z01.PrintRune(rune(tens) +'0')
    z01.PrintRune(rune(ones) +'0')
    if i<789{z01.PrintRune(',')}
  }
  }
  
  }
