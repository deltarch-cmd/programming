#include <stdio.h>

// C does not have the bool type, so usually it is done like this
#define BOOL char
#define FALSE 0
#define TRUE 1

int main() {
    int a = 3;
    float b = 4.5;
    double c = 5.25;
    float sum = a + b + c;

    printf("The sum of a, b, and c is %f.", sum);
    return 0;
}
