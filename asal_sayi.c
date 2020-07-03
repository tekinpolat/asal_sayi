#include <stdio.h>
#include <stdbool.h>
bool asal_sayi_mi(int sayi){
    int max_kontrol_sayi = (int)sayi/2;
    for(int index = 2; index < max_kontrol_sayi; index++){
        if (sayi%index == 0){
            return false;
        }
    }

    return true;
}

int main(){
    for (int index = 2; index < 1000000; index++){
        bool sonuc = asal_sayi_mi(index);
        if(sonuc){
            printf("%d\n", index);
        }
    }
    return 0;
}