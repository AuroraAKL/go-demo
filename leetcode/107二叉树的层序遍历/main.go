package main

import "fmt"

func threeSum(nums []int) [][]int {
  vis := make(map[int]map[int]map[int]bool)
 n := len(nums)
 m := make(map[int]int)

 for i, num := range nums {
   m[num] = i + 1
 }

 var ans [][]int
 for i := 0; i < n; i++ {
   for j := i + 1; j < n; j++ {
     t := nums[i] + nums[j]
     k := m[-t] - 1
     if k < 0 {
       continue
     }
     fmt.Println(i, j, k, isVis(nums[i], nums[j], nums[k], vis))
     if k > i && k > j && !isVis(nums[i], nums[j], nums[k], vis) {
       // 如果存在
       ans = append(ans, []int{nums[i], nums[j], nums[k]})
       visNum(nums[i], nums[j], nums[k], vis)
     }
   }
 }
  return ans
}

func isVis(a int, b int, c int, vis map[int]map[int]map[int]bool) bool {
  return vis[a] != nil && vis[a][b] != nil && vis[a][b][c]
}

func visNum(a int, b int, c int, vis map[int]map[int]map[int]bool) {
  visANum(a, b, c, vis)
  visANum(a, c, b, vis)
  visANum(b, a, c, vis)
  visANum(b, c, a, vis)
  visANum(c, a, b, vis)
  visANum(c, b, a, vis)
}

func visANum(a int, b int, c int, vis map[int]map[int]map[int]bool)  {
  if vis[a] == nil {
    vis[a] = make(map[int]map[int]bool)
  }
  if vis[a][b] == nil {
    vis[a][b] = make(map[int]bool)
  }
  vis[a][b][c] = true
}


func main() {

  threeSum([]int{-1, 0, 1})

}
