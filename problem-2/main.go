package main

import (
    "fmt"
    "strconv"
    "strings"
)

/*
 * Complete the GetMoneySpent function below.
 */
func GetMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
    /*
     * Write your code here.
	 */
	fmt.Println(keyboards)
	fmt.Println(drives)
	fmt.Println(b)
    return 0
}

func main() {

    bnm := strings.Split("10 2 3", " ")

    bTemp, err := strconv.ParseInt(bnm[0], 10, 64)
    if err != nil {
		fmt.Println(err)
	}
    b := int32(bTemp)

    nTemp, err := strconv.ParseInt(bnm[1], 10, 64)
    if err != nil {
		fmt.Println(err)
	}
    n := int32(nTemp)

    mTemp, err := strconv.ParseInt(bnm[2], 10, 64)
    if err != nil {
		fmt.Println(err)
	}
    m := int32(mTemp)

	// keyboardsTemp := strings.Split(readLine(reader), " ")
	
	keyboardsTemp := []string{"3", "1",}

    var keyboards []int32

    for keyboardsItr := 0; keyboardsItr < int(n); keyboardsItr++ {
        keyboardsItemTemp, err := strconv.ParseInt(keyboardsTemp[keyboardsItr], 10, 64)
        if err != nil {
			fmt.Println(err)
		}
        keyboardsItem := int32(keyboardsItemTemp)
        keyboards = append(keyboards, keyboardsItem)
    }

    drivesTemp := strings.Split("5 2 8", " ")

    var drives []int32

    for drivesItr := 0; drivesItr < int(m); drivesItr++ {
		drivesItemTemp, err := strconv.ParseInt(drivesTemp[drivesItr], 10, 64)
        if err != nil {
			fmt.Println(err)
		}
        drivesItem := int32(drivesItemTemp)
        drives = append(drives, drivesItem)
    }

    /*
     * The maximum amount of money she can spend on a keyboard and USB drive, or -1 if she can't purchase both items
     */

    moneySpent := GetMoneySpent(keyboards, drives, b)

    fmt.Printf("%d\n", moneySpent)

}