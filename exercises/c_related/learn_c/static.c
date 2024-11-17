#include <stdio.h>

int sum (int num) {
    static int added = 0;
    added += num;
    return added;
}

int main() {
    printf("%d ",sum(55));
    printf("%d ",sum(45));
    printf("%d ",sum(50));
    return 0;
}
