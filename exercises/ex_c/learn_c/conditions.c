#include <stdio.h>

#define GUESS 555

void guessNumber(int guess) {
    if (guess < GUESS) {
        printf("Your guess is too low.\n");
        return;
    }

    if (guess > GUESS) {
        printf("Your guess is too high.\n");
        return;
    }
    
    printf("Correct. You guessed it!\n");
}

int main() {
    guessNumber(500);
    guessNumber(600);
    guessNumber(555);
}
