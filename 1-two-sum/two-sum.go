// Running time O(n^2)

// func twoSum(nums []int, target int) []int {
//     for i := range nums {
//         for j := i + 1; j < len(nums); j++ {
//             if nums[i] + nums[j] == target {
//                 return []int{i, j}
//             }
//         }
//     }

//     return nil
// }


func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

    for i, num := range nums {
        complement := target - num
        if j, ok := seen[complement]; ok {
            return []int{i, j}
        }

        seen[num] = i
    }

    return nil
}
