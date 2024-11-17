#include <stdio.h>

int main() {
    int array[] = {1, 7, 4, 5, 9, 3, 5, 11, 6, 3, 4};
    int i = -1;
    int a = -1;

    while (i < 10) {
        i++;
        /* your code goes here */
        a = array[i];

        if (a < 5) continue;
        if (a > 10) break;

        printf("%d\n", a);
    }

    return 0;
}
