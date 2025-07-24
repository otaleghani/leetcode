package main

func IsAnagram(s string, t string) bool {
    var m = make(map[byte]int)
    
    if (len(s) != len(t)) {
        return false
    }

    for i := range len(s) {
        m[s[i]] += 1
        m[t[i]] -= 1
    }

    for _, value := range m {
        if value != 0 {
            return false;
        }
    }

    return true;
}
