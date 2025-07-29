# Add Single Number

## 問題
給一個 single linked list 的 head 及一個非負整數 digit，linked list 以反向順序（即每個節點的值代表一位數，最低位在最前面）表示一個非負整數。

請撰寫一個函式，將整數加到 linked list 所表示的數字上，並回傳新的 linked list head node

保證輸入的 linked list 不為空，且 digit 在 0 到 9 之間。

### Example
```
Input: head = [2, 4, 3], digit = 9  
Output: [1, 5, 3]  
說明：linked list 表示的數字為 342，342 + 9 = 351，回傳的 linked list 為 [1, 5, 3]。
```

