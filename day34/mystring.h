#ifndef new
#define new {.equals=equals, .length=length, .equalsIgnoreCase=equalsIgnoreCase, .substring=substring,.toUpperCase=toUpperCase, .toLowerCase=toLowerCase, .replace=replace, .lastIndexOf=lastIndexOf, .firstIndexOf=firstIndexOf};
#endif

#include "mystring.c"

int equals(char* str1, char* str2);
int length(char *str);
int firstIndexOf(char *str,char i);
int lastIndexOf(char *str,char i);
void toUpperCase(char *str);
void toLowerCase(char *str);
void replace(char *str, char o, char n);
void substring(char *str, char *sub, int ini, int end);
int equalsIgnoreCase(char *string, char *string_aux);

typedef struct {
    int   (*equals)(char*, char*);
    int   (*length)(char*);
    int   (*equalsIgnoreCase)(char*, char*);
    void  (*substring)(char*, char*, int, int);
    void  (*toUpperCase)(char*);
    void  (*toLowerCase)(char*);
    void  (*replace)(char*, char, char);
    int   (*lastIndexOf)(char*, char);
    int   (*firstIndexOf)(char*, char);

}String;
