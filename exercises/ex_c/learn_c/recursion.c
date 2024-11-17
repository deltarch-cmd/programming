#include <stdio.h>

int factorial(int num){
    if (num == 0) return 1;
    if (num == 1) return num;
    else return num * factorial(num-1);
}

int main() {
    /* testing code */
    printf("0! = %i\n", factorial(0));
    printf("1! = %i\n", factorial(1));
    printf("3! = %i\n", factorial(3));
    printf("5! = %i\n", factorial(5));
}
