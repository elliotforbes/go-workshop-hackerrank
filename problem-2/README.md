Problem 2
===========

Monica wants to buy a keyboard and a USB drive from her favorite electronics store. The store has several models of each. Monica wants to spend as much as possible for the `2` items, given her budget.

Given the price lists for the store's keyboards and USB drives, and Monica's budget, find and print the amount of money Monica will spend. If she doesn't have enough money to both a keyboard and a USB drive, print -1 instead. She will buy only the two required items.

For example, suppose she has `b = 60` to spend. Three types of keyboards cost `keyboards = [40, 50, 60]`. Two USB drives cost `drives = [5, 8, 12]`. She could purchase a `40 keyboard + 12 USB drive = 52`, or a `50 keyboard + 8 USB drive = 58`. She chooses the latter. She can't buy more than `2` items so she can't spend exactly `60`.

## Function Description

Complete the `getMoneySpent` function in the editor below. It should return the maximum total price for the two items within Monica's budget, or `-1` if she cannot afford both items.

`getMoneySpent` has the following parameter(s):

* keyboards: an array of integers representing keyboard prices
* drives: an array of integers representing drive prices
* b: the units of currency in Monica's budget

## Input Format

The first line contains three space-separated integers `b`, `n`, and `m`, her budget, the number of keyboard models and the number of USB drive models. 
The second line contains `n` space-separated integers `keyboard[i]`, the prices of each keyboard model. 
The third line contains `m` space-separated integers `drives`, the prices of the USB drives.

## Output Format

Print a single integer denoting the amount of money Monica will spend. If she doesn't have enough money to buy one keyboard and one USB drive, print `-1` instead.

### Sample Input 0

```
10 2 3
3 1
5 2 8
```

### Sample Output 0

```
9
```

### Explanation 0

She can buy the  keyboard and the  USB drive for a total cost of .

### Sample Input 1

```
5 1 1
4
5
```

### Sample Output 1

```
-1
```

### Explanation 1

There is no way to buy one keyboard and one USB drive because , so we print -1.