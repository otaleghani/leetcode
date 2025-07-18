function canConstruct(ransomNote: string, magazine: string): boolean {
    const noteMap = new Map<string, number>;
    
    // Fills the noteMap with the frequencies of the chars in random note
    // At the end it will be Map<char, frequency>
    for (let i = 0; i < ransomNote.length; i++) {
        if (noteMap.get(ransomNote[i])) {
            noteMap.set(ransomNote[i], noteMap.get(ransomNote[i]) + 1);
        } else {
            noteMap.set(ransomNote[i], 1);
        }
    }

    // Everytime we find a char that we need we remove 1 from the frequency
    for (let i = 0; i < magazine.length; i++) {
        if (noteMap.get(magazine[i])) {
            noteMap.set(magazine[i], noteMap.get(magazine[i]) - 1);
        }
    }

    // Finally we check if the frequencies are all 0
    for (const [key, val] of noteMap) {
        if (val > 0) return false;
    }

    return true;
};
