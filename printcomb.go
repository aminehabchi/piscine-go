package piscine

import "github.com/01-edu/z01"

func PrintComb(){
 ones :='0'
  tens :='0'
  hands :='0'
  
  
  for i:=0 ; i<=789 ; i++ {

    if ones<'9'{
      ones++
    }else if tens<'9'{
      tens++;ones='0'
    }else{
      hands++;tens='0';ones='0'
    }
    if ones>tens && tens >hands{
     
    z01.PrintRune(hands)
    z01.PrintRune(tens)
    z01.PrintRune(ones )
      
    if hands!='7'  {
      z01.PrintRune(',')
      z01.PrintRune(' ')}
    }
  }
        z01.PrintRune('\n')}
  }
