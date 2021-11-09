#include <stdio.h>

// recurcive method for factorial

int factorial(int n); // function protorype
int power(int n, int m);

int main() {

    int n;

    scanf("%d", &n);

    int factorial_result = factorial(n);
    int power_result = power(n, 3);

    printf("factorial is %d\n", factorial_result);
    printf("power is %d\n", power_result);

}

int factorial(int n) { // function decleration
    if (n == 1) {
        return 1;
    }
    // n! = n * (n-1)!
    return n * factorial(n-1);
}

int power(int n, int m) {
    int result = 1;

    for (int i = 1; i <= m; i++) {
        result = result * n;
    }

    return result;
}