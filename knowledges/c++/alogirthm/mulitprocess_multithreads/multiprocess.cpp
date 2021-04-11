#include <unistd.h>  
#include <stdio.h>   
#include<stdlib.h>

using namespace std;
int main ()   
{   
    pid_t fpid; //fpid表示fork函数返回的值  
    int count=0;  
    fpid=fork(); //拷贝一个新的进程  
    if (fpid < 0)   
        printf("error in fork!");   
    else if (fpid == 0) {  
        printf("i am the child process（子进程）, my process id is %d\n",getpid());   
        printf("我是爹的儿子\n");//对某些人来说中文看着更直白。  
        for(int i = 0;i < 5;i++){
           printf("子进程 i=: %d\n",i);  
            sleep(1);
        }
        count++;  
         exit(0);
    }  
    else {      
        printf("i am the parent process(父进程), my process id is %d\n",getpid());   
        printf("我是孩子他爹\n");  
        for(int i = 0;i < 5;i++){
            printf("父进程 i=: %d\n",i);  
            sleep(1);
        }
        count++;  
        exit(0);
    }  
    printf("统计结果是: %d/n",count);  
    return 0;  
}  