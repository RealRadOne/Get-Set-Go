package main
	
import (
	"fmt"
	"time"
	"bytes"
	"runtime"
	"strconv"
 )
func getGID() uint64 {
    b := make([]byte, 64)
    b = b[:runtime.Stack(b, false)]
    b = bytes.TrimPrefix(b, []byte("goroutine "))
    b = b[:bytes.IndexByte(b, ' ')]
    n, _ := strconv.ParseUint(string(b), 10, 64)
    return n
}

func assigner1(n int,m int,c chan int)	{
	//Goroutine that runs parallel to main
	fmt.Println("This was the work of thread ",getGID())
	ele :=n*m
    c <- ele
}
func assigner2(n int,m int,c chan int) {
	fmt.Println("This was the work of ",getGID())
	ele :=n+m
	c <- ele
	//The goroutines can communicate to each other by the use of channels
}
func main(){
	fmt.Println("Enter the number of elements")
    var matrix1[100][100] int
    var matrix2[100][100] int

    //var n int

    //fmt.Scanln(&n)-In case user input is needed

    for n :=2 ; n <50; n+=5{
    //Checking the time needed for different load

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {

          //Distributing the work of assigning elements by calling a goroutine

           c := make(chan int)
           go assigner1(i,j,c)
           x := <-c
           //Value recieved by the channel
           matrix1[i][j]=x
        }
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
           p := make(chan int)
           go assigner2(i,j,p)
           y := <-p
           matrix2[i][j]=y
        }
    }
    //Printing the matrix for the user convenience
    fmt.Println()
    fmt.Println("========== Matrix1 =============")
    fmt.Println()
     for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
           fmt.Printf("%d ",matrix1[i][j])
        }
        fmt.Println()
    }
    fmt.Println()
    fmt.Println("========== Matrix2 =============")
    fmt.Println()
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
           fmt.Printf("%d ",matrix2[i][j])
        }
        fmt.Println()
    }
    //Checking the performance as per load

    //Getting the time taken to perform the multiplica


    start:=time.Now()
			
	var mult int
	var subs int
	for i :=0; i < n;i++ {
		mult=0
		for j :=0; j <n;j++{
			subs=0
			for k :=0;k <n;k++{
			//Calling goroutine to parallelize multiplication
			 c := make(chan int)
             go assigner1(matrix1[j][k],matrix2[i][k],c)
             time.Sleep(30 * time.Millisecond)
             x := <-c
             subs=x
             p := make(chan int)
             go assigner2(subs,mult,p)
             time.Sleep(30 * time.Millisecond)
             y := <-p
             mult=y
			}
		}
    fmt.Println(mult)
	}

	t :=time.Now()
	elapsed :=t.Sub(start)
	fmt.Println("The time taken for operation is ",elapsed)
}

}