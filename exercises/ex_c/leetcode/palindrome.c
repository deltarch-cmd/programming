#include <stdio.h>

// Binary exponentiation (pretty cool!)
unsigned long int power(int base, int exp) {
    if (exp == 0)
        return 1;
    else if (exp % 2)
        return base * power(base, exp - 1);
    else {
        unsigned long int temp = power(base, exp / 2);
        return temp * temp;
    }
}

int isPalindrome(int x) {
    if (x < 0) return 0;

    int digits = 1;
    unsigned long int aux = 10;

    // First I calculate the number of digits of the number
    while (aux <= x) {
        digits++;
        aux *= 10;

        printf("DIGITS: %d; AUX: %ld; NUMBER: %d\n", digits, aux, x);
    }

    // Trackers to make the division
    int start = 0;
    int end = digits - 1;
    
    // 'is_not' tell us if the number is palindrome
    int is_not = 0;

    while (is_not != 1) {
        // Mod is used to take the last digit
        int front = x / power(10, end) % 10;
        int back = x / power(10, start) % 10;

        printf("FRONT: %d; BACK: %d\n", front, back);

        if (front != back) is_not = 1;
        else {
            start++; end--;
            // If the trackers meet means there are no more digits
            if (start > end) break;
        }
    }
    return !is_not;
}

int main(){
    int x = 1410110141;
    int aux = isPalindrome(x);

    if (aux == 0) printf("El número %d no es palíndromo\n", x);
    else printf("El número %d ES palíndromo!\n", x);
}
