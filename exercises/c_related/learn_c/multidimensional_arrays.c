#include <stdio.h>

int main() {
    float average;
    int i;
    int j;

    int grades[2][5];

    grades[0][0] = 80;
    grades[0][1] = 70;
    grades[0][2] = 65;
    grades[0][3] = 89;
    grades[0][4] = 90;

    grades[1][0] = 85;
    grades[1][1] = 80;
    grades[1][2] = 80;
    grades[1][3] = 82;
    grades[1][4] = 87;

    int row = sizeof(grades) / sizeof(grades[0]);
    int column = sizeof(grades) / (sizeof(grades[0][0]) * row);

    for (i = 0; i < row; i++) {
        average = 0;
        for (j = 0; j < column; j++) {
            average += grades[i][j];
        }

        average /= 5;
        printf("The average marks obtained in subject %d is: %.2f\n", i, average);
    }

    return 0;
}
