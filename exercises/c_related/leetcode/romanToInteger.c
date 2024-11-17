#include <stdio.h>

int checkStrLength(char* s){
    int lenght = 0;

    while(*s != '\0'){
        lenght++;
        *s++;
    }

    return lenght;
}

int romanToInt(char* s){
    int l = checkStrLength(s);
}

int main() {

}
