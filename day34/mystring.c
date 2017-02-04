#include <stdlib.h>
#include <stdio.h>

#define TRUE 1
#define FALSE 0
#define NOTFOUND -1

int length(char *str){
    int i=0;
    while(*str++ != '\0'){
        i++;
    }
    return i;
}

int firstIndexOf(char *str, char c){
    int indice = 0;
    while(*str != '\0'){
        if(*str == c){
            return indice;
        }
        indice++;
        str++;
    }
    return NOTFOUND;
}

int lastIndexOf(char *str, char c){
    int tamanho = length(str);
    int i;
    for (i = tamanho; i > 0; i--) {
        if(*(str + i -1) == c){
            return i -1;
        }
    }
    return NOTFOUND;
}

int equals(char *str1, char *str2){
    int tamanho_str1 = length(str1);
    int tamanho_str2 = length(str2);
    if (tamanho_str1 != tamanho_str2) {
        return FALSE;
    }
    while(*str1++!='\0' && *str2++!='\0'){
        if (*str1 != *str2) {
            return FALSE;
        }
    }
    return TRUE;
}

char toLower(char c){
    if(c >= 65 && c <= 90){
        return c + 32;
    }
    return c;
}

void toUpperCase(char *str){
    while(*str != '\0'){
        if(*str >= 97 && *str <= 122){
            *str = *str - 32;
        }
        str++;
    }
}

void toLowerCase(char *str){
    while(*str != '\0'){
        *str = toLower(*str);
        str++;
    }
}

int equalsIgnoreCase(char *str1, char *str2){
    if (length(str1) != length(str2)){
        return FALSE;
    }
    while(*str1 != '\0'){
        if (toLower(*str1) != toLower(*str2)){
            return FALSE;
        }
        str1++;
        str2++;
    }
    return TRUE;
}

void substring(char *str, char *sub, int ini, int end){
    int i, j;
    for (i = ini,j=0; i < end; ++i, j++) {
        sub[j] = str[i];
    }
    sub[end - ini] = '\0';
}

void replace(char *str, char o, char n){
    while(*str != '\0'){
        if (*str == o){
            *str = n;
        }
        str++;
    }
}
