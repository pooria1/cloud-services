#include <stdio.h>

int main()
{

    // one dimentional arrays

    int jafarzadeh[5];

    jafarzadeh[0] = 1;
    jafarzadeh[1] = 5;
    jafarzadeh[2] = 6;
    jafarzadeh[3] = 14;
    jafarzadeh[4] = 43;

    int kalantari[100];

    for (int i = 0; i < 100; i++)
    {
        kalantari[i] = i;
    }

    int soleymani[10];

    for (int i = 0; i < 10; i++)
    {
        scanf("%d", &soleymani[i]);
    }

    printf("everything received\n");
    printf("soley is: %d %d %d %d\n", soleymani[2], soleymani[4], soleymani[5], soleymani[9]);

}