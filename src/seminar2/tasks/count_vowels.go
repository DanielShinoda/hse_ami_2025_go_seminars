package tasks

// CountVowels подсчитывает количество гласных в строке
func CountVowels(s string) int {
    count := 0
    for _, w := range s {
        for _, w1 := range "aeiouAEIOU" {
            if w == w1 {
                count += 1
                break
            }
        }
    }
    return count
}