#include <stdio.h>


int main() {

    int a[2][3] = {
        {1,2,4},
        {6,7,8}
    };

    int b[3][2] = {
        {1,2},
        {4,3},
        {6,4}
    };

    int c[2][3][2] = {
        {{1,2},{3,4},{5,6}},
        {{7,8},{9,10},{11,12}}
    };

/*
    DO NOT forget to use correct indexing for arrays 
    and showing their position!!!
*/

    printf("a[1][2] is %d\n", a[0][1]);
    printf("c[2][2][2] is %d\n", c[1][1][1]);

}