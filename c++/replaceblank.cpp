/*************************************************************************
	> File Name: voidreplaceblank.cpp
	> Author: 
	> Mail: 
	> Created Time: Wed 25 Feb 2015 03:11:25 PM CST
 ************************************************************************/

#include<iostream>
using namespace std;
#include<cstring>

void ReplaceBlank(char string[], int length){
    if(string == NULL || length <=0){
        return;
    }

    int originlength=strlen(string)+1;
    int blankNumber=0;

    int i=0;
    while(string[i] != '\0'){
        if(string[i]==' '){
            blankNumber++;
        }
        i++;
    }

    int newlength=originlength+2*blankNumber;

    if(newlength > length){
        return;
    }

    int indexorigin=originlength;
    int indexnew=newlength;

    while(indexorigin >=0 && indexorigin != indexnew){
        if(string[indexorigin] == ' '){
            string[indexnew--]='0';
            string[indexnew--]='2';
            string[indexnew--]='%';
            indexorigin--;
        }else{
            string[indexnew--]=string[indexorigin--];
        }
    }
}

int main(){
    char str[100]="we are happy";
    ReplaceBlank(str,100);
    cout<<str<<endl;
}
