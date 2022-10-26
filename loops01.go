/* Alta3 Research | RZFeeser
   Looping - basic for-loop    */

package main

import "fmt"

func main() {

    total := 0  // init total
    
    // init;      condition; post statement
    for {
        fmt.Println("The value of alta -> ", alta)
        
        total += alta   // use total to keep track of alta
        break
    }

fmt.Println("\nThe value of total is", total)
    
}

