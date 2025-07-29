# Intersection point

[Leetcode](https://leetcode.com/problems/intersection-of-two-linked-lists)

## 問題
找出兩個 Linked list 交集的節點。

### 範例:
```text
Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
Output: Intersected at '8'
```

## 思路
![intersection](./intersection.png)

如果有交集，兩個 Linked list 到結尾的步數是相同的。

1. 先算出這兩個 linked list 的長度
2. 比較兩者長度
3. 讓兩者起點是一樣的
4. 開始遍歷，如果遇到節點值一樣，則表示該 Node 是交點，