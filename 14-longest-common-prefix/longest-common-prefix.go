/* 
Method 1: Repeatedly trim the prefix

- the prefix is initially set to the first string in the array
- compare the prefix with the next string
- if there isn't a common prefix, we can return "" early
- if there is a common prefix, we set the prefix to the common prefix
- repeat until the array is exhausted

Let's say:

m = number of strings in the array
n = the length of the shortest string

In the worst-case, for each of m strings we might compare up to n characters when shrinking the prefix. So complexity is O(m * n)
*/

func longestCommonPrefix(strs []string) string {
    prefix := strs[0]

    for i := 1; i < len(strs); i++ {
        j := 0

        for j < len(strs[i]) && j < len(prefix) && prefix[j] == strs[i][j] {
            j++
        }
        
        prefix = prefix[:j]
        if prefix == "" {
            return ""
        }
    }

    return prefix
}
