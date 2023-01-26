package main

import "fmt"

func main(){

//  result,x:= a()
//  fmt.Println(result)
//  fmt.Println(x)
 w := new(int)
 name := new(string)
 result,x :=c( w,name)
 fmt.Println(x)
 fmt.Println(result)
 fmt.Println(*name)

//  r,_ := b(1,2,3,4,5,6,6)
//  fmt.Print(r)
}

func  a()(int,string){
	return 122, "abc";
}
func  b(args ...int)(bool,int){
  for _,v := range args{
	fmt.Println(v)
  }
   return true,11;
}
func  c(w *int , name *string)(i int , j string){
   i = 10
   j = "rahul"
   *w = 100
   *name = "code"
   return  
}