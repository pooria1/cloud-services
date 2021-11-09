#include <stdio.h>

int main() {
    int password = 1234;

    int input;
    
    printf("Enter password: ");
    scanf("%d", &input);

/*
    we use single '=' for assigning variables.
    but we use double '==' for checking conditions
    for example:
        a == b
        a > b
        a < b
        a >= b
        a <= b
        a != b
*/

    if (input == password)
    {
        printf("Correct password! WELCOME!!!\n");
    } else {
        printf("Wrong Password!\n");
    }
    

}