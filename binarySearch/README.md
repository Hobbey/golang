```
TEST CASE 1, Search 14 in array [11 14 19 20 31 45 63 70 78]
--------------------
0 4 8
14 <= arr[4], Target on Left: [0,4]
--------------------
0 2 4
14 <= arr[2], Target on Left: [0,2]
--------------------
0 1 2
14 <= arr[1], Target on Left: [0,1]
--------------------
0 0 1
arr[0] < 14, Target on Right: (0,1] ==> [1,1]
for loop Exit
**********************
Value found

TEST CASE 2, Search 71 in array [11 14 19 20 31 45 63 70 78]
--------------------
0 4 8
arr[4] < 71, Target on Right: (4,8] ==> [5,8]
--------------------
5 6 8
arr[6] < 71, Target on Right: (6,8] ==> [7,8]
--------------------
7 7 8
arr[7] < 71, Target on Right: (7,8] ==> [8,8]
for loop Exit
**********************
Value in range but NOT found

TEST CASE 3, Search 888 in array [11 14 19 20 31 45 63 70 78]
Array not valid or given value out of range

TEST CASE 4, Search 70 in array []
Array not valid or given value out of range
```

