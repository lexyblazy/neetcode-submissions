class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isPalindrome(s: string): boolean {
        if (s.length < 1) {
            return true;
        }
        const input = s
            .toLowerCase()
            .split(" ")
            .join("")
            .replace(/[^a-zA-Z0-9]/g, "");


        for (let i = 0; i < Math.ceil(input.length / 2); i++) {
            if (input[i] !== input[input.length - i - 1]) {
                return false;
            }
        }

        return true;
    }
}
